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
	position := " , , , , , , , , "
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Voulez vous etre client ou serveur?")
		fmt.Println("1 serveur")
		fmt.Println("2 client")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)
		switch option {
		case "1":
			Server(position)
		case "2":
			Client(position)
		}
	}
}

func Server(position string) {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 8000")
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		message = strings.TrimSpace(message)
		fmt.Println(message)
		place, _ := strconv.Atoi(message)
		position = TicTacToe(place, position)
		fmt.Print("Text to send: ")
		newmessage, _ := reader.ReadString('\n')
		newmessage = strings.TrimSpace(newmessage)
		place, _ = strconv.Atoi(newmessage)
		position = TicTacToe(place, position)
		conn.Write([]byte(newmessage + "\n"))
	}
}

func Client(position string) {
	adresseip := os.Args[1] + ":8000"
	conn, err := net.Dial("tcp", adresseip)
	if err != nil {
		log.Fatal(err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		place, _ := strconv.Atoi(text)
		position = TicTacToe(place, position)
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		message = strings.TrimSpace(message)
		fmt.Println(message)
		place, _ = strconv.Atoi(message)
		position = TicTacToe(place, position)
	}
}

func TicTacToe(place int, position string) string {
	player := 0
	positionList := strings.Split(position, ",")
	fmt.Println(place)
	for i := range positionList {
		if positionList[i] == " " {
			player++
		}
	}
	if player%2 == 0 {
		positionList[place-1] = "X"
	} else {
		positionList[place-1] = "O"
	}
	fmt.Println("|" + positionList[6] + "|" + positionList[7] + "|" + positionList[8] + "|")
	fmt.Println("|" + positionList[3] + "|" + positionList[4] + "|" + positionList[5] + "|")
	fmt.Println("|" + positionList[0] + "|" + positionList[1] + "|" + positionList[2] + "|")
	position = strings.Join(positionList, ",")
	fmt.Println(position)
	etat := win(position)
	if etat == "egalite" {
		fmt.Println("egalite")
		os.Exit(0)
	} else if etat == "server" {
		fmt.Println("server win")
		os.Exit(0)
	} else if etat == "client" {
		fmt.Println("client win")
		os.Exit(0)
	}
	return position
}


func win(position string) string {
	positionList := strings.Split(position, ",")
	if positionList[0]!= " " && positionList[1]!= " " && positionList[2]!= " " && positionList[3]!= " " && positionList[4]!= " " && positionList[5]!= " " && positionList[6]!= " " && positionList[7]!= " " && positionList[8]!= " "{
		return "egalite"
	}
	if positionList[0] == positionList[1] && positionList[1] == positionList[2] && positionList[0]!= " "{
		if positionList[0] == "X"{
			return "server"
		} else {
			return "client"
		}
	}
	if positionList[3] == positionList[4] && positionList[4] == positionList[5] && positionList[3]!= " "{
		if positionList[3] == "X"{
			return "server"
		} else {
			return "client"
		}
	}
	if positionList[6] == positionList[7] && positionList[7] == positionList[8] && positionList[6]!= " "{
		if positionList[6] == "X"{
			return "server"
		} else {
			return "client"
		}
	}
	if positionList[0] == positionList[3] && positionList[3] == positionList[6] && positionList[0]!= " "{
		if positionList[0] == "X"{
			return "server"
		} else {
			return "client"
		}
	}
	if positionList[1] == positionList[4] && positionList[4] == positionList[7] && positionList[1]!= " "{
		if positionList[1] == "X"{
			return "server"
		} else {
			return "client"
		}
	}
	if positionList[2] == positionList[5] && positionList[5] == positionList[8] && positionList[2]!= " "{
		if positionList[2] == "X"{
			return "server"
		} else {
			return "client"
		}
	}
	if positionList[0] == positionList[4] && positionList[4] == positionList[8] && positionList[0]!= " "{
		if positionList[0] == "X"{
			return "server"
		} else {
			return "client"
		}
	}
	if positionList[2] == positionList[4] && positionList[4] == positionList[6] && positionList[2]!= " "{
		if positionList[2] == "X"{
			return "server"
		} else {
			return "client"
		}
	}
	return "none"

}
