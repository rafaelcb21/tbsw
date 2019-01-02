package main

import (
    "log"
    "net/http"
    "golang.org/x/net/html"
	"strings"
	"sync"
	//"fmt"
	"strconv"
)

type Naves struct {
	player string
	nomeGuilda string
	nome string
	allycode string
	codechar string
	estrelas string
	nivel string
}



func getShips(urls []string, nomeGuilda string)  chan Naves {
	ch := make(chan Naves)

	wg := &sync.WaitGroup{}
	wg.Add(len(urls))

    for _, url := range urls {
		go listarTodasNavesDaGuilda(url, nomeGuilda, ch, wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func listarTodasNavesDaGuilda(url, nomeGuilda string, ch chan Naves, wg *sync.WaitGroup) {
	var lista []string
	var lista2 []string
	var indexs []int
	var character Naves
	res, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	defer res.Body.Close()
	page := html.NewTokenizer(res.Body)
	getNext := false

	for {
		_ = page.Next()
		token := page.Token()
		
		if token.Type == html.ErrorToken {
			break
		}

		//fmt.Println("@#$", token.Attr, " | ",token.Data, " | ",token.DataAtom, " | ",token.Type, " | ",html.TextToken)

		if len(token.Attr) == 3 && token.Data == "a" && token.Attr[0].Key == "href" && token.Attr[1].Val == "ship-portrait-full-link" {
			lista = append(lista, token.Attr[0].Val) //link
			getNext = false
		}

		if len(token.Attr) == 1 && token.Data == "div" && token.Attr[0].Key == "class" && string(token.Attr[0].Val)[0:2] == "sh" {
			s := string(token.Attr[0].Val)
			sizeS := len(s)
			classe := s[sizeS -1 : sizeS]
			
			if classe != "s" && s[19:20] == "s" {
				lista = append(lista, token.Attr[0].Val) //estrelas
			}
			
			getNext = false
		} 

		if getNext == true {
			lista = append(lista, token.Data)
			getNext = false
		}

		if len(token.Attr) == 3 && token.Data == "a" && token.Attr[0].Key == "class" && token.Attr[0].Val == "collection-ship-name-link" {
			getNext = true //nome da nave
		}

		if len(token.Attr) == 1 && token.Data == "div" && token.Attr[0].Key == "class" && string(token.Attr[0].Val) == "ship-portrait-full-frame-level" {
			getNext = true //level
		}
		
		if len(token.Attr) == 1 && token.Data == "h5" && token.Attr[0].Key == "class" && string(token.Attr[0].Val) == "panel-title text-center m-b-sm" {
			getNext = true //player
		}
	}

	
	
	//lista2 = lista[1: len(lista) - 1]
	lista2 = lista[1:]
	//fmt.Println(lista[1:], "=======", lista2)
	player := strings.Replace(lista[0], "\n", "", -1)
	for i, val := range lista2 {
		if string(val[0]) == "/" {
			indexs = append(indexs, i)
		}
	}

	
	for i := 0; i < len(indexs); i++ {
		if i != len(indexs) - 1 {
			l := lista2[indexs[i]:indexs[i+1]]

			link := strings.Split(l[0], "/")
			allycode := link[2]
			codechar := link[4]

			character = Naves {
				player: player,
				nomeGuilda: nomeGuilda,
				nome: l[9],
				allycode: allycode,
				codechar: codechar,
				estrelas: estrelasNave(l[1:8]),
				nivel: l[len(l)-2],
			}
			
			ch <- character

		} else {
			l := lista2[indexs[i]:len(lista2)]
			link := strings.Split(l[0], "/")
			allycode := link[2]
			codechar := link[4]
			//fmt.Println(lista2[indexs[i]:], len(lista2))
			character = Naves {
				player: player,
				nomeGuilda: nomeGuilda,
				nome: l[9],
				allycode: allycode,
				codechar: codechar,
				estrelas: estrelasNave(l[1:8]),
				nivel: l[len(l)-2],
			}
			
			ch <- character

		}
	}
	wg.Done()
}

func estrelasNave(stars []string) string {
	count := 0
	for _, valor := range stars {
		if len(valor) == 24 {
			count++
		}
	}
	return strconv.Itoa(count)
}