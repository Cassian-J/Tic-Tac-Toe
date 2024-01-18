package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Voulez vous etre client ou serveur?")
		fmt.Println("1 serveur")
		fmt.Println("2 client")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)
		switch option {
		case "1":
			Server()
		case "2":
			Client()
		}
	}
}

func Server() {
	ln, err := net.Listen("tcp", ":7680")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 7680")
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Message Received:", string(message))
		fmt.Print("Text to send: ")
		newmessage, _ := reader.ReadString('\n')
		conn.Write([]byte(newmessage + "\n"))
	}
}

func Client() {	
	adresseip := os.Args[1]+":7680"
	conn, err := net.Dial("tcp", adresseip)
	if err != nil {
		log.Fatal(err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
