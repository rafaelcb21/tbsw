package main

import (
	"fmt"
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

func main() {
	sumPgPersonagem := 34562239
	sumPgNave := 24248873
	//sumPgPersonagemPgNave := 58811112

	pgsPhasesCharLS := [][]int{
		[]int{885000, 7465000, 53065000},
		[]int{1900000, 3800000, 19200000, 39000000, 82800000, 137800000},
		[]int{3510000, 7020000, 29420000, 57020000, 109220000, 174020000},
		[]int{5220000, 10440000, 38740000, 73440000, 136040000, 214140000},
		[]int{11100000, 25200000, 66200000, 115500000, 187100000, 276900000},
		[]int{26400000, 57400000, 116700000, 188700000, 270200000, 370200000},
	}

	pgsPhasesShipsLS := [][]int{
		[]int{1920000, 18420000, 44720000},
		[]int{2176000, 20876000, 50676000},
		[]int{18000000, 52000000, 102000000},
		[]int{21600000, 62400000, 122400000},
	}

	for _, i := range pgsPhasesCharLS {
		x := stars(sumPgPersonagem, i)
		fmt.Println(x)
	}
	fmt.Println("======================")
	for _, i := range pgsPhasesShipsLS {
		x := stars(sumPgNave, i)
		fmt.Println(x)
	}	
}

func stars(pg int, pgs[]int) int {
	for i, value := range pgs {
		if pg < value {
			
			switch i {
			case 0:
				return 0
			case 1:
				return 1
			case 2:
				return 2
			case 3:
				return 3
			case 4:
				return 4
			case 5:
				return 5
			}
		} else if pg > value && i == 5 {
			return 6
		}
	}
	return 0
}

func combatPhase1LS(nivel, players int, percent float32) float32 {
	combat := [...]int{24000, 51000, 91000, 144000, 211000, 291000}
	pgCombat := float32(combat[nivel - 1] * 2 * players) * percent

	return pgCombat
}

func combatPhase2LS(nivel, players int, percent float32) float32 {
	combat := [...]int{43000, 72000, 115000, 172000, 243000, 329000}
	pgCombat := float32(combat[nivel - 1] * 2 * players) * percent

	return pgCombat
}

func combatPhase3LS(nivel, players int, percent float32) float32 {
	combat := [...]int{65000, 96000, 142000, 203000, 280000, 372000}
	pgCombat := float32(combat[nivel - 1] * 2 * players) * percent

	return pgCombat
}

func combatPhase4LS(nivel, players int, percent float32) float32 {
	combat := [...]int{76000, 111000, 163000, 232000, 319000, 423000}
	pgCombat := float32(combat[nivel - 1] * 2 * players) * percent

	return pgCombat
}

func combatPhase5LS(nivel, players int, percent float32) float32 {
	combat := [...]int{90000, 128000, 185000, 261000, 356000, 470000}
	pgCombat := float32(combat[nivel - 1] * players) * percent

	return pgCombat
}

func combatPhase6LS(nivel, players int, percent float32) float32 {
	combat := [...]int{152000, 191000, 249000, 327000, 424000, 541000}
	pgCombat := float32(combat[nivel - 1] * players) * percent

	return pgCombat
}

func combatPhasesShipLS(phase, players int, percent float32) float32 {
	combat := [...]int{0, 0, 371000, 478000, 536000, 614000} 
	pgCombatShip := float32(combat[phase - 1] * players) * percent 

	return pgCombatShip
}

func find(lista []Personagem, estrela int, parametro []string, estrelas string) {
	for _, dict := range lista {
		if contains(parametro, dict.codechar) && dict.estrelas == estrelas {

		}
	}
}

func contains(stringSlice []string, searchString string) bool {
    for _, value := range stringSlice {
        if value == searchString {
            return true
        }
    }
    return false
}



  