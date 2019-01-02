package main

import (
	"fmt"
    "log"
	"net/http"
	"regexp"
	"io"
    "io/ioutil"
)

var url string = "https://swgoh.gg"

func getGuilda() (string, string) {
    res, err := http.Get("https://swgoh.gg/p/948239354/")
    if err != nil {
        log.Fatal(err)
    }

	regex := "<p>Guild <strong class=\"pull-right\"><a href=\"([a-zA-Z0-9/-]+)\">(.*)</a"
	lista := parseHTML(res.Body, regex)

	urlGuilda := lista[0][1]
	nomeGuilda := lista[0][2]

	defer res.Body.Close()
	
	return urlGuilda, nomeGuilda

}

 
func getUrlsMembersGuild(characters, guilda, ships string) ([]string, []string, []string) {

	var urlsCharacters []string
	var urlsPG []string
	var urlsShips []string

	res, err := http.Get(guilda)
	
    if err != nil {
        log.Fatal(err)
    }

    regex := "href=\"(/p/[0-9]+/)"
    lista := parseHTML(res.Body, regex)

    defer res.Body.Close()

    for _, v := range lista {
		urlsCharacters = append(urlsCharacters, url + v[1] + characters)
		urlsPG = append(urlsPG, url + v[1])
		urlsShips = append(urlsShips, url + v[1] + ships)
    }

    return urlsCharacters, urlsPG, urlsShips

}

func parseHTML(res io.Reader, regex string) [][]string {
    robots, err := ioutil.ReadAll(res)

    if err != nil {
        log.Fatal(err)
    }

    r, _ := regexp.Compile(regex)
    guild := r.FindAllStringSubmatch(string(robots), -1)

    return guild

}


func main() {
    guilda, nomeGuilda := getGuilda()
	urlsMembersCharacters, urlMembersPG, urlsMembersShips := getUrlsMembersGuild("characters/", url + guilda, "ships/")
	
	chPersonagens := getCharacters(urlsMembersCharacters, nomeGuilda)
	chNaves := getShips(urlsMembersShips, nomeGuilda)
	chPG := getPG(urlMembersPG, nomeGuilda)

	var todosPersonagensDaGuilda []Personagem
	var todasNavesDaGuilda []Naves
	var todosPGsDaGuilda []PG
	var pgPersonagem []int
	var pgNave []int
	var sumPgPersonagem int
	var sumPgNave int

	for i := range chPersonagens {
		todosPersonagensDaGuilda = append(todosPersonagensDaGuilda, i)
	}

	for i := range chNaves {
		todasNavesDaGuilda = append(todasNavesDaGuilda, i)
	}

	for i := range chPG {
		pgPersonagem = append(pgPersonagem, i.pgPersonagens)
		pgNave = append(pgNave, i.pgNaves)
		todosPGsDaGuilda = append(todosPGsDaGuilda, i)
	}

	for _, i := range pgPersonagem {
		sumPgPersonagem += i
	}

	for _, i := range pgNave {
		sumPgNave += i
	}

	fmt.Println(todosPersonagensDaGuilda)
	fmt.Println("=========================")
	fmt.Println(todasNavesDaGuilda)
	fmt.Println("=========================")
	fmt.Println(todosPGsDaGuilda)
	fmt.Println("=========================")
	fmt.Println(sumPgPersonagem, sumPgNave, sumPgPersonagem+sumPgNave)


}
