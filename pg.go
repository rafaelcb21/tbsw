package main

import (
    "log"
    "net/http"
    //"golang.org/x/net/html"
	"strings"
	"sync"
	"strconv"
	//"fmt"
	"io/ioutil"
	"regexp"
)

type PG struct {
	player string
	nomeGuilda string
	allycode string
	pgPersonagens int
	pgNaves int
	pgTotal int
}

func getPG(urls []string, nomeGuilda string)  chan PG {
	ch := make(chan PG)

	wg := &sync.WaitGroup{}
	wg.Add(len(urls))

    for _, url := range urls {
		go listarTodosPGsDaGuilda(url, nomeGuilda, ch, wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func listarTodosPGsDaGuilda(url, nomeGuilda string, ch chan PG, wg *sync.WaitGroup) {
	res, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	defer res.Body.Close()


	bodyBytes, _ := ioutil.ReadAll(res.Body)
	bodyString := string(bodyBytes)

	regexPG := "<p>Galactic Power <strong class=\"pull-right\">([0-9,]+)</strong></p>"
	regexPersonagemPG := "<p>Galactic Power \\(Characters\\) <strong class=\"pull-right\">([0-9,]+)</strong></p>"
	regexNavesPG := "<p>Galactic Power \\(Ships\\) <strong class=\"pull-right\">([0-9,]+)</strong></p>"
	regexPlayer := "<h5 class=\"panel-title text-center\">(.*)</h5>"

	listaPG := parseStringHTML(bodyString, regexPG)
	listaPersonagemPG := parseStringHTML(bodyString, regexPersonagemPG)
	listaNavesPG := parseStringHTML(bodyString, regexNavesPG)
	player := parseStringHTML(bodyString, regexPlayer)[0][1]

	pgPlayer := strings.Replace(listaPG[0][1], ",", "", -1)
	pgPlayerInt, _ := strconv.Atoi(pgPlayer)

	pgPersonagemPlayer := strings.Replace(listaPersonagemPG[0][1], ",", "", -1)
	pgPersonagemPlayerInt, _ := strconv.Atoi(pgPersonagemPlayer)

	pgNavesPlayer := strings.Replace(listaNavesPG[0][1], ",", "", -1)
	pgNavesPlayerInt, _ := strconv.Atoi(pgNavesPlayer)

	link := strings.Split(url, "/")
	allycode := link[4]

	pgPlayerRegister := PG {
		player: player,
		nomeGuilda: nomeGuilda,
		allycode: allycode,
		pgPersonagens: pgPersonagemPlayerInt,
		pgNaves: pgNavesPlayerInt,
		pgTotal: pgPlayerInt,
	}

	ch <- pgPlayerRegister
	
	wg.Done()
}

func parseStringHTML(res, regex string) [][]string {
    r, _ := regexp.Compile(regex)
    guild := r.FindAllStringSubmatch(res, -1)

    return guild
}



