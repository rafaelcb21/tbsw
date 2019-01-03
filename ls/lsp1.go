package main

import (
	"fmt"
)

func main() {
	sumPgPersonagem := 34562239
	//sumPgNave := 24248873
	//sumPgPersonagemPgNave := 58811112



	x := stars(sumPgPersonagem, []int{1900000, 3800000, 19200000, 39000000, 82800000, 137800000})
	fmt.Println(x)


	//{885000, 7465000, 53065000}
	//{1900000, 3800000, 19200000, 39000000, 82800000, 137800000}
	//{3510000, 7020000, 29420000, 57020000, 109220000, 174020000}
	//{5220000, 10440000, 38740000, 73440000, 136040000, 214140000}
	//{11100000, 25200000, 66200000, 115500000, 187100000, 276900000}
	//{26400000, 57400000, 116700000, 188700000, 270200000, 370200000}
}

func stars(pg int, pgs[]int) int {
	for i, value := range pgs {
		fmt.Println(pg, value, i)
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



  