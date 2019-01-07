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
var totalEstrelas []int
var totalEstrela int
var estrelasPorPhase []int
var pgCombatByPhase []float32

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

	//fmt.Println(todosPersonagensDaGuilda)
	//fmt.Println("=========================")
	//fmt.Println(todasNavesDaGuilda)
	//fmt.Println("=========================")
	//fmt.Println(todosPGsDaGuilda)
	//fmt.Println("=========================")
	//fmt.Println(sumPgPersonagem, sumPgNave, sumPgPersonagem+sumPgNave)


	pgsPhasesCharLS := [][]int{
		[]int{885000, 7465000, 53065000},
		[]int{1900000, 3800000, 19200000, 39000000, 82800000, 137800000},
		[]int{3510000, 7020000, 29420000, 57020000, 109220000, 174020000},
		[]int{5220000, 10440000, 38740000, 73440000, 136040000, 214140000},
		//[]int{11100000, 25200000, 66200000, 115500000, 187100000, 276900000},
		//[]int{26400000, 57400000, 116700000, 188700000, 270200000, 370200000},
	}

	//pgsPhasesShipsLS := [][]int{
	//	[]int{1920000, 18420000, 44720000},
	//	[]int{2176000, 20876000, 50676000},
	//	[]int{18000000, 52000000, 102000000},
	//	[]int{21600000, 62400000, 122400000},
	//}

	//for phase, i := range pgsPhasesCharLS {
		pgCombatPhase1, gePhase1 := combatPhase1LS(6, len(pgPersonagem), 1.0, todosPersonagensDaGuilda)
		pgCombatPhase2, gePhase2 := combatPhase2LS(6, len(pgPersonagem), 1.0, todosPersonagensDaGuilda)
		pgCombatPhase3, roloPhase3 := combatPhase3LS(6, len(pgPersonagem), 1.0, todosPersonagensDaGuilda)
		pgCombatPhase4, gePhase4 := combatPhase4LS(6, len(pgPersonagem), 1.0, todosPersonagensDaGuilda)

		pgCombatByPhase = []float32{pgCombatPhase1, pgCombatPhase2, pgCombatPhase3, pgCombatPhase4}
		//totalPG := pgCombatPhase1 + pgCombatPhase2  + pgCombatPhase3 + pgCombatPhase4 + float32(sumPgPersonagem)
		//x := stars(pgCombatPhase1, pgCombatPhase2, pgCombatPhase3, pgCombatPhase4, i)
		//totalEstrelas = append(totalEstrelas, x)
		//estrelasPorPhase = append(estrelasPorPhase, x)
		fmt.Println("=>", gePhase1, gePhase2, roloPhase3, gePhase4)
	//}

	x := stars(pgCombatByPhase, pgsPhasesCharLS)
	fmt.Println("==>", x)

	for _, i := range totalEstrelas {
		totalEstrela += i
	}
	fmt.Println(estrelasPorPhase)
	fmt.Println(totalEstrela)


	//fmt.Println("======================")
	//for phase, i := range pgsPhasesShipsLS {
	//	x := stars(float32(sumPgNave), i)
	//	fmt.Println("=>", sumPgNave, phase + 1, x)
	//}

}

func stars(pg []float32, pgs [][]int) []int {
	var estrelas []int
	for i, value := range pgs {
		for j, val := range value {
			fmt.Println(int(pg[i]), val)
			if pg[i] < float32(val) {
				estrelas = append(estrelas, i)
				
			} else if pg[i] > float32(val) && j == len(value) - 1 {
				fmt.Println(int(pg[i]), val)
				estrelas = append(estrelas, len(value))
			}
		}	
	}
	return estrelas
}
