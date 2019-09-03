package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
	"os"
)

func onMessageClient(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('\n')

		fmt.Println(msg)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Your name: ")
	nameReader := bufio.NewReader(os.Stdin)
	nameInput, _ := nameReader.ReadString('\n')
	nameInput = nameInput[:len(nameInput) - 1]

	fmt.Println("****** Message ******")
	go onMessageClient(conn)
	for {
		msgReader := bufio.NewReader(os.Stdin)
		msg, err := msgReader.ReadString('\n')
		if err != nil {
			break
		}

		msg = fmt.Sprintf("%s: %s\n", nameInput, msg[:len(msg) - 1])
		conn.Write([]byte(msg))
	}
	conn.Close()
}