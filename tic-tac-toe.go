package main

import "fmt"
import "strconv"
func currentplayer() {
	var currentPlayer string = "X"
if currentPlayer == "X" {
	currentPlayer = "O"
} else {
	currentPlayer = "X"
}

}
func main() {
	var case1, case2, case3, case4, case5, case6, case7, case8, case9 int
	var win int
	var bestplayer string

	if win != 1 {
		currentplayer()
		var currentcase string
		fmt.Println("Quelle case voulez-vous choisir ?")
		fmt.Scanln(&currentcase)
		if currentcase == strconv.Itoa(case1) {
			case1 = 1
		}
		if currentcase == strconv.Itoa(case2) {
			case2 = 1
		}
		if currentcase == strconv.Itoa(case3) {
			case3 = 1
		}
		if currentcase == strconv.Itoa(case4) {
			case4 = 1
		}
		if currentcase == strconv.Itoa(case5) {
			case5 = 1
		}
		if currentcase == strconv.Itoa(case6) {
			case6 = 1
		}
		if currentcase == strconv.Itoa(case7) {
			case7 = 1
		}
		if currentcase == strconv.Itoa(case8) {
			case8 = 1
		}
		if currentcase == strconv.Itoa(case9) {
			case9 = 1
		}
		if case1 == 1 {
			case1 = strconv.Itoa(currentplayer)
		}
		if case2 == 1 {
			case2 = currentplayer()
		}
		if case3 == 1 {
			case3 = currentplayer
		}
		if case4 == 1 {
			case4 = currentplayer
		}
		if case5 == 1 {
			case5 = currentplayer
		}
		if case6 == 1 {
			case6 = currentplayer
		}
		if case7 == 1 {
			case7 = currentplayer
		}
		if case8 == 1 {
			case8 = currentplayer
		}
		if case9 == 1 {
			case9 = strconv.Atoi(currentplayer)
		}
		fmt.Println("|" + strconv.Itoa(case1) + "|" + strconv.Itoa(case2) + "|" + strconv.Itoa(case3) + "|+\n+|" + strconv.Itoa(case4) + "|" + strconv.Itoa(case5) + "|" + strconv.Itoa(case6) + "|+\n+|" + strconv.Itoa(case7) + "|" + strconv.Itoa(case8) + "|" + strconv.Itoa(case9) + "|")
	} else {
		fmt.Println("bravo" + bestplayer + "vous avez gagné")
	}

}





