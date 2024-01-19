package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	spaces := strings.Repeat(" ", 50)
	spaces2 := strings.Repeat(" ", 35)
	bleu := "\033[34m" // Define the different colors
	vert := "\033[32m"
	reset := "\033[0m"
	position := " , , , , , , , , " // Define the game board
	var etat string
	reader := bufio.NewReader(os.Stdin)
	for {
		//We choose if we are the server or the client
		fmt.Print("\033[H\033[2J")
		fmt.Println(spaces, vert+"                    ||                 ", spaces)
		fmt.Println(spaces, "                    ||                 ", spaces)
		fmt.Println(spaces, "-------------------Menu--------------------", spaces)
		fmt.Println(spaces, "|                                         |", spaces)
		fmt.Println(spaces, "|   Voulez vous etre client ou serveur?   |", spaces)
		fmt.Println(spaces, "|                                         |", spaces)
		fmt.Println(spaces, "|   [1] Serveur                           |", spaces)
		fmt.Println(spaces, "|   [2] Client                            |", spaces)
		fmt.Println(spaces, "|                                         |", spaces)
		fmt.Println(spaces, "-------------------------------------------", spaces+reset)
		fmt.Println()
		fmt.Println()
		fmt.Println(spaces2, bleu+"████████╗██╗ ██████╗    ████████╗ █████╗  ██████╗    ████████╗ ██████╗ ███████╗")
		fmt.Println(spaces2, "╚══██╔══╝██║██╔════╝    ╚══██╔══╝██╔══██╗██╔════╝    ╚══██╔══╝██╔═══██╗██╔════╝")
		fmt.Println(spaces2, "   ██║   ██║██║            ██║   ███████║██║            ██║   ██║   ██║█████╗  ")
		fmt.Println(spaces2, "   ██║   ██║██║            ██║   ██╔══██║██║            ██║   ██║   ██║██╔══╝  ")
		fmt.Println(spaces2, "   ██║   ██║╚██████╗       ██║   ██║  ██║╚██████╗       ██║   ╚██████╔╝███████╗")
		fmt.Println(spaces2, "   ╚═╝   ╚═╝ ╚═════╝       ╚═╝   ╚═╝  ╚═╝ ╚═════╝       ╚═╝    ╚═════╝ ╚══════╝"+reset)

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)
		switch option {
		case "1":
			//If we are the server , we run the server function
			fmt.Print("\033[H\033[2J")
			Server(position, etat)
		case "2":
			//If we are the client , we run the client function
			fmt.Print("\033[H\033[2J")
			fmt.Print("Entrez l'adresse ip du serveur: ")
			ip, _ := reader.ReadString('\n')
			ip = strings.TrimSpace(ip)
			Client(position, etat, ip)
		}
	}
}

func Server(position string, etat string) { // Server function
	ln, err := net.Listen("tcp", ":8000") // Listen on the port 8000
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Le client se connecte⌛")
	conn, err := ln.Accept() // Accept the connection
	if err != nil {
		log.Fatal(err)
	}
	positionList := strings.Split(position, ",")
	fmt.Println("Les positions correspndent au placement des numéros du clavier numérique") // We display the game board
	fmt.Println("|" + positionList[6] + "|" + positionList[7] + "|" + positionList[8] + "|")
	fmt.Println("|" + positionList[3] + "|" + positionList[4] + "|" + positionList[5] + "|")
	fmt.Println("|" + positionList[0] + "|" + positionList[1] + "|" + positionList[2] + "|")
	for {
		fmt.Println("l'adversaire joue")
		message, _ := bufio.NewReader(conn).ReadString('\n') // Read the message
		place, _ := strconv.Atoi(strings.TrimSpace(message))
		position, etat = TicTacToe(place, position, conn)
		move := moves(position)
		if etat != "none" { // If the game is over , we exit the program
			os.Exit(0)
			break
		}
		reader := bufio.NewReader(os.Stdin) // We ask the server to play
		fmt.Print("A vous de jouer:")
		newmessage, _ := reader.ReadString('\n')
		place, _ = strconv.Atoi(strings.TrimSpace(newmessage))
		move = moves(position)
		for !Play(move, place) { // We check if the move is possible, if the move is not possible we ask the server to play again
			reader := bufio.NewReader(os.Stdin)
			fmt.Println("Cette position est deja prise!")
			fmt.Print("Choisissez une autre position: ")
			newmessage, _ = reader.ReadString('\n')
			place, _ = strconv.Atoi(strings.TrimSpace(newmessage))
		}
		position, etat = TicTacToe(place, position, conn)
		conn.Write([]byte(newmessage + "\n")) // We send the move to the client
	}
}

func Client(position string, etat string, ip string) { // Client function
	adresseip := ip + ":8000" // We get the ip address of the server
	conn, err := net.Dial("tcp", adresseip)
	positionList := strings.Split(position, ",")
	fmt.Println("Les positions correspndent au placement des numéros du clavier numérique") // We display the game board
	fmt.Println("|" + positionList[6] + "|" + positionList[7] + "|" + positionList[8] + "|")
	fmt.Println("|" + positionList[3] + "|" + positionList[4] + "|" + positionList[5] + "|")
	fmt.Println("|" + positionList[0] + "|" + positionList[1] + "|" + positionList[2] + "|")
	if err != nil {
		log.Fatal(err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("A vous de jouer:") // We ask the client to play
		text, _ := reader.ReadString('\n')
		place, _ := strconv.Atoi(strings.TrimSpace(text))
		move := moves(position)
		for !Play(move, place) { // We check if the move is possible, if the move is not possible we ask the client to play again
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("\033[H\033[2J")
			fmt.Println("Cette position est deja prise!")
			fmt.Print("Choisissez une autre position: ")
			text, _ = reader.ReadString('\n')
			place, _ = strconv.Atoi(strings.TrimSpace(text))
		}
		position, etat = TicTacToe(place, position, conn) // Modify the game board with the move of the client with the TicTacToe function
		fmt.Fprintf(conn, text+"\n")

		fmt.Println("l'adversaire joue")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		place, _ = strconv.Atoi(strings.TrimSpace(message))
		position, etat = TicTacToe(place, position, conn)
		move = moves(position)
		if etat != "none" { // If the game is over , we exit the program
			os.Exit(0)
			break
		}
	}
}

func Play(move []string, place int) bool {
	for i := range move { // We check if the move is possible
		if move[i] == strconv.Itoa(place) {
			return true
		}
	}
	return false
}

func moves(position string) []string { // Wecheck the possible moves
	var moveList []string
	positionList := strings.Split(position, ",")
	for i := range positionList {
		if positionList[i] == " " {
			moveList = append(moveList, strconv.Itoa(i+1))
		}
	}
	return moveList
}

func TicTacToe(place int, position string, conn net.Conn) (string, string) {
	player := 0
	positionList := strings.Split(position, ",")
	fmt.Println(place)
	for i := range positionList { // We search wich player has to play
		if positionList[i] == " " {
			player++
		}
	}
	if player%2 == 0 && place != 0 { //If it's the server's turn , we put an X
		positionList[place-1] = "X"
	} else if place != 0 { //If it's the client's turn , we put an O
		positionList[place-1] = "O"
	}
	fmt.Print("\033[H\033[2J")
	fmt.Print("\033[H\033[2J")
	fmt.Println("Les positions correspndent au placement des numéros du clavier numérique") // We display the game board
	fmt.Println("|" + positionList[6] + "|" + positionList[7] + "|" + positionList[8] + "|")
	fmt.Println("|" + positionList[3] + "|" + positionList[4] + "|" + positionList[5] + "|")
	fmt.Println("|" + positionList[0] + "|" + positionList[1] + "|" + positionList[2] + "|")
	position = strings.Join(positionList, ",")
	etat := win(position)  // We check if the game is over
	if etat == "egalite" { //If there is an equality , we display it
		fmt.Println("egalite")
	} else if etat == "server" {
		fmt.Println("server win") //If the server win , we display it
	} else if etat == "client" {
		fmt.Println("client win") //If the client win , we display it
	}
	return position, etat
}

func win(position string) string {
	positionList := strings.Split(position, ",")
	// We check if there is a winner or an equality
	if positionList[0] != " " && positionList[1] != " " && positionList[2] != " " && positionList[3] != " " && positionList[4] != " " && positionList[5] != " " && positionList[6] != " " && positionList[7] != " " && positionList[8] != " " {
		return "egalite"
	}
	if positionList[0] == positionList[1] && positionList[1] == positionList[2] && positionList[0] != " " {
		if positionList[0] == "X" {
			return "server"
		} else {
			return "client"
		}
	}
	if positionList[3] == positionList[4] && positionList[4] == positionList[5] && positionList[3] != " " {
		if positionList[3] == "X" {
			return "server"
		} else {
			return "client"
		}
	}
	if positionList[6] == positionList[7] && positionList[7] == positionList[8] && positionList[6] != " " {
		if positionList[6] == "X" {
			return "server"
		} else {
			return "client"
		}
	}
	if positionList[0] == positionList[3] && positionList[3] == positionList[6] && positionList[0] != " " {
		if positionList[0] == "X" {
			return "server"
		} else {
			return "client"
		}
	}
	if positionList[1] == positionList[4] && positionList[4] == positionList[7] && positionList[1] != " " {
		if positionList[1] == "X" {
			return "server"
		} else {
			return "client"
		}
	}
	if positionList[2] == positionList[5] && positionList[5] == positionList[8] && positionList[2] != " " {
		if positionList[2] == "X" {
			return "server"
		} else {
			return "client"
		}
	}
	if positionList[0] == positionList[4] && positionList[4] == positionList[8] && positionList[0] != " " {
		if positionList[0] == "X" {
			return "server"
		} else {
			return "client"
		}
	}
	if positionList[2] == positionList[4] && positionList[4] == positionList[6] && positionList[2] != " " {
		if positionList[2] == "X" {
			return "server"
		} else {
			return "client"
		}
	}
	return "none"

}
