package main

import (
    "log"
    "net/http"
    "golang.org/x/net/html"
	"strings"
	"sync"
	"strconv"
)

type Personagem struct {
	player string
	nomeGuilda string
	nome string
	allycode string
	codechar string
	estrelas string
	zeta string
	nivel string
	gear string
}



func getCharacters(urls []string, nomeGuilda string)  chan Personagem {
	ch := make(chan Personagem)

	wg := &sync.WaitGroup{}
	wg.Add(len(urls))

    for _, url := range urls {
		go listarTodosPersonagensDaGuilda(url, nomeGuilda, ch, wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func listarTodosPersonagensDaGuilda(url, nomeGuilda string, ch chan Personagem, wg *sync.WaitGroup) {
	var lista []string
	var lista2 []string
	var indexs []int
	var character Personagem
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

		if len(token.Attr) == 3 && token.Data == "a" && token.Attr[0].Key == "href" {
			lista = append(lista, token.Attr[0].Val)
			getNext = false
		}

		if len(token.Attr) == 5 && token.Data == "img" && token.Attr[2].Key == "alt" {
			lista = append(lista, token.Attr[2].Val)
			getNext = false
		}

		if len(token.Attr) == 1 && token.Data == "div" && token.Attr[0].Key == "class" && string(token.Attr[0].Val)[0:3] == "sta" {
			lista = append(lista, token.Attr[0].Val)
			getNext = false
		} 

		if getNext == true {
			lista = append(lista, token.Data)
			getNext = false
		}

		if len(token.Attr) == 1 && token.Data == "div" && token.Attr[0].Key == "class" && string(token.Attr[0].Val) == "char-portrait-full-zeta" {
			getNext = true
		}

		if len(token.Attr) == 1 && token.Data == "div" && token.Attr[0].Key == "class" && string(token.Attr[0].Val) == "char-portrait-full-level" {
			getNext = true
		}

		if len(token.Attr) == 1 && token.Data == "div" && token.Attr[0].Key == "class" && string(token.Attr[0].Val) == "char-portrait-full-gear-level" {
			getNext = true
		}
		
		if len(token.Attr) == 1 && token.Data == "h5" && token.Attr[0].Key == "class" && string(token.Attr[0].Val) == "panel-title text-center m-b-sm" {
			getNext = true
		}
	}

	lista2 = lista[1: len(lista) - 1]
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

			var zeta string

			if len(l) == 12 {
				zeta = l[9]
			} else {
				zeta = "0"
			}

			character = Personagem {
				player: player,
				nomeGuilda: nomeGuilda,
				nome: l[1],
				allycode: allycode,
				codechar: codechar,
				estrelas: estrelas(l[2:9]),
				zeta: zeta,
				nivel: l[len(l)-2],
				gear: l[len(l)-1],
			}
			
			ch <- character

		} else {
			l := lista2[indexs[i]:len(lista2)]

			link := strings.Split(l[0], "/")
			allycode := link[2]
			codechar := link[4]

			var zeta string

			if len(l) == 12 {
				zeta = l[9]
			} else {
				zeta = "0"
			}

			character = Personagem {
				player: player,
				nomeGuilda: nomeGuilda,
				nome: l[1],
				allycode: allycode,
				codechar: codechar,
				estrelas: estrelas(l[2:9]),
				zeta: zeta,
				nivel: l[len(l)-2],
				gear: l[len(l)-1],
			}
			
			ch <- character

		}
	}
	wg.Done()
}

func estrelas(stars []string) string {
	count := 0
	for _, valor := range stars {
		if len(valor) == 10 {
			count++
		}
	}
	return strconv.Itoa(count)
}



