package main

import (
	"fmt"
	"strconv"
	//"sync"
	//"reflect"
)

var squadPhoenix = []string{"chopper", "ezra-bridger", "garazeb-zeb-orrelios", "hera-syndulla", "kanan-jarrus", "sabine-wren"}

var squadRogueOne = []string{"baze-malbus", "bistan", "bodhi-rook", "cassian-andor", "chirrut-imwe", "jyn-erso", "k-2so", "pao", "scarif-rebel-pathfinder"}
	
var squadRebels = []string{"hoth-rebel-soldier", "hoth-rebel-scout", "chopper", "ezra-bridger", "garazeb-zeb-orrelios", "hera-syndulla", "kanan-jarrus", "sabine-wren",
"baze-malbus", "bistan", "bodhi-rook", "cassian-andor", "chirrut-imwe", "jyn-erso", "k-2so", "pao", "scarif-rebel-pathfinder",
"admiral-ackbar", "ahsoka-tano-fulcrum", "biggs-darklighter", "captain-han-solo", "commander-luke-skywalker", "han-solo",
"lando-calrissian", "lobot", "luke-skywalker-farmboy", "obi-wan-kenobi-old-ben", "princess-leia",
"r2-d2", "rebel-officer-leia-organa", "tormtrooper-han", "wedge-antilles", "chewbacca", "c-3po"}




func combatPhase1LS(nivel, integrantes int, percent float32, lista []Personagem) (float32, float32) {
	combat := [...]int{24000, 51000, 91000, 144000, 211000, 291000}
	pgCombat := float32(combat[nivel - 1] * 2 * integrantes) * percent
	totalJogadores := len(jogadores(lista))
	guildaPhoenix := find(jogadores(lista), lista, 2, squadPhoenix, "2" )
	var integrantesSpecial []string
	var melhorarIntegrantesSpecial []string

	for _, dict := range guildaPhoenix {
		for k, v:= range dict {
			if len(v) >= 5 {
				integrantesSpecial = append(integrantesSpecial, k)
			} else {
				melhorarIntegrantesSpecial = append(melhorarIntegrantesSpecial, k)
			}
		}
	}

	ge := float32(len(integrantesSpecial) * 7) //atendem ao requisito
	gePossivel := float32(totalJogadores * 7)

	fmt.Println("*******PHASE 1*******")
	fmt.Println(len(integrantesSpecial))
	fmt.Println("max GE :", ge, gePossivel, (1.0-(ge/gePossivel))*100)
	fmt.Println(melhorarIntegrantesSpecial)

	return pgCombat, ge
}

func combatPhase2LS(nivel, integrantes int, percent float32, lista []Personagem) (float32, float32) {
	combat := [...]int{43000, 72000, 115000, 172000, 243000, 329000}
	pgCombat := float32(combat[nivel - 1] * 2 * integrantes) * percent
	totalJogadores := len(jogadores(lista))
	guildaRebelsWithHRSoldierCombat := find(jogadores(lista), lista, 3, squadRebels, "3")
	guildaRogueOne := find(jogadores(lista), lista, 3, squadRogueOne, "3")

	var integrantesCombat []string
	var integrantesSpecial []string
	var melhorarIntegrantesCombat []string
	var melhorarIntegrantesSpecial []string
	
	for _, dict := range guildaRebelsWithHRSoldierCombat {
		for k, v:= range dict {
			if len(v) >= 5 && contains(v, "hoth-rebel-soldier") {
				integrantesCombat = append(integrantesCombat, k)
			} else {
				melhorarIntegrantesCombat = append(melhorarIntegrantesCombat, k)
			}
		}
	}

	pgCombat = pgCombat + float32(combat[nivel - 1] * len(integrantesCombat)) * percent

	for _, dict := range guildaRogueOne {
		for k, v:= range dict {
			if len(v) >= 5 {
				integrantesSpecial = append(integrantesSpecial, k)
			} else {
				melhorarIntegrantesSpecial = append(melhorarIntegrantesSpecial, k)
			}
		}
	}

	ge := float32(len(integrantesSpecial) * 8) //atendem ao requisito
	gePossivel := float32(totalJogadores * 8)

	fmt.Println("*******PHASE 2*******")
	fmt.Println(len(integrantesCombat))
	fmt.Println(len(integrantesSpecial))
	fmt.Println("max GE :", ge, gePossivel, (1.0-(ge/gePossivel))*100)
	fmt.Println(melhorarIntegrantesCombat)
	fmt.Println(melhorarIntegrantesSpecial)

	return pgCombat, ge
}

func combatPhase3LS(nivel, integrantes int, percent float32, lista []Personagem) (float32, float32) {
	combat := [...]int{65000, 96000, 142000, 203000, 280000, 372000}
	pgCombat := float32(combat[nivel - 1] * 2 * integrantes) * percent

	totalJogadores := len(jogadores(lista))
	guildaRebelsWithHRScoutCombat := find(jogadores(lista), lista, 4, squadRebels, "4")
	guildaHRSoldierSpecial := find(jogadores(lista), lista, 5, []string{"hoth-rebel-soldier"}, "5")

	var integrantesCombat []string
	var integrantesSpecial []string
	var melhorarIntegrantesCombat []string
	var melhorarIntegrantesSpecial []string
	
	for _, dict := range guildaRebelsWithHRScoutCombat {
		for k, v:= range dict {
			if len(v) >= 5 && contains(v, "hoth-rebel-scout") {
				integrantesCombat = append(integrantesCombat, k)
			} else {
				melhorarIntegrantesCombat = append(melhorarIntegrantesCombat, k)
			}
		}
	}

	pgCombat = pgCombat + float32(combat[nivel - 1] * len(integrantesCombat)) * percent

	for _, dict := range guildaHRSoldierSpecial {
		for k, v:= range dict {
			if len(v) == 1 {
				integrantesSpecial = append(integrantesSpecial, k)
			} else {
				melhorarIntegrantesSpecial = append(melhorarIntegrantesSpecial, k)
			}
		}
	}

	rolo := float32(len(integrantesSpecial) * 2) //atendem ao requisito
	roloPossivel := float32(totalJogadores * 2)

	fmt.Println("*******PHASE 3*******")
	fmt.Println(len(integrantesCombat))
	fmt.Println(len(integrantesSpecial))
	fmt.Println("max ROLO :", rolo, roloPossivel, (1.0-(rolo/roloPossivel))*100)
	fmt.Println(melhorarIntegrantesCombat)
	fmt.Println(melhorarIntegrantesSpecial)

	return pgCombat, rolo
}

func combatPhase4LS(nivel, integrantes int, percent float32, lista []Personagem) (float32, float32) {
	combat := [...]int{76000, 111000, 163000, 232000, 319000, 423000}
	pgCombat := float32(combat[nivel - 1] * 2 * integrantes) * percent

	totalJogadores := len(jogadores(lista))
	guildaRebelsWithHRSoldierCombat := find(jogadores(lista), lista, 5, squadRebels, "5")
	guildaRoloSpecial := find(jogadores(lista), lista, 5, []string{"rebel-officer-leia-organa"}, "5")

	var integrantesCombat []string
	var integrantesSpecial []string
	var melhorarIntegrantesCombat []string
	var melhorarIntegrantesSpecial []string
	
	for _, dict := range guildaRebelsWithHRSoldierCombat {
		for k, v:= range dict {
			if len(v) >= 5 && contains(v, "hoth-rebel-soldier") {
				integrantesCombat = append(integrantesCombat, k)
			} else {
				melhorarIntegrantesCombat = append(melhorarIntegrantesCombat, k)
			}
		}
	}

	pgCombat = pgCombat + float32(combat[nivel - 1] * len(integrantesCombat)) * percent

	for _, dict := range guildaRoloSpecial {
		for k, v:= range dict {
			if len(v) == 1 {
				integrantesSpecial = append(integrantesSpecial, k)
			} else {
				melhorarIntegrantesSpecial = append(melhorarIntegrantesSpecial, k)
			}
		}
	}

	ge := float32(len(integrantesSpecial) * 20) //atendem ao requisito
	gePossivel := float32(totalJogadores * 20)

	fmt.Println("*******PHASE 4*******")
	fmt.Println(len(integrantesCombat))
	fmt.Println(len(integrantesSpecial))
	fmt.Println("max GE :", ge, gePossivel, (1.0-(ge/gePossivel))*100)
	fmt.Println(melhorarIntegrantesCombat)
	fmt.Println(melhorarIntegrantesSpecial)

	return pgCombat, ge
}

func combatPhase5LS(nivel, integrantes int, percent float32, lista []Personagem) float32 {
	combat := [...]int{90000, 128000, 185000, 261000, 356000, 470000}
	pgCombat := float32(combat[nivel - 1] * integrantes) * percent

	//squadPhoenix
	//squadRebels => "hoth-rebel-scout"
	//special only "commander-luke-skywalker"


	return pgCombat
}

func combatPhase6LS(nivel, integrantes int, percent float32, lista []Personagem) float32 {
	combat := [...]int{152000, 191000, 249000, 327000, 424000, 541000}
	pgCombat := float32(combat[nivel - 1] * integrantes) * percent

	//squadRogueOne
	//squadRebels
	//special only "rebel-officer-leia-organa"

	return pgCombat
}

func combatPhasesShipLS(phase, players int, percent float32) float32 {
	combat := [...]int{0, 0, 371000, 478000, 536000, 614000} 
	pgCombatShip := float32(combat[phase - 1] * players) * percent 

	return pgCombatShip
}

func find(players []string, lista []Personagem, estrela int, parametro []string, estrelas string) []map[string][]string {
	var todosPlayers []map[string][]string
	for _, player := range players {
		todosPlayers = append(todosPlayers, buscaEntreTodosPlayers(lista, estrela, parametro, estrelas, player))
	}
	
	return todosPlayers
}

func buscaEntreTodosPlayers(lista []Personagem, estrela int, parametro []string, estrelas string, player string) map[string][]string {
	var list []string
	
	dictPlayerCharacters := make(map[string][]string)
	for _, dict := range lista {
		estrelasPersonagem, _ := strconv.Atoi(dict.estrelas)
		estrelasParametro, _ := strconv.Atoi(estrelas)
		if contains(parametro, dict.codechar) &&  estrelasPersonagem >= estrelasParametro && dict.player == player {
			list = append(list, dict.codechar)
		}
	}
	
	dictPlayerCharacters[player] = list

	return dictPlayerCharacters

}

func contains(stringSlice []string, searchString string) bool {
    for _, value := range stringSlice {
        if value == searchString {
            return true
        }
    }
    return false
}

func jogadores(elements []Personagem) []string {
	var playersList []string
	encountered := map[string]bool{}

	for _, value := range elements {
        if encountered[value.player] == true {
            // Do not add duplicate.
        } else {
            encountered[value.player] = true
            playersList = append(playersList, value.player)
        }
	}

    return playersList
}


  