package main

import (
	"fmt"
	"log"
)

type Message string

type Greeter struct {
	Mess Message
}

type Event struct {
	Greeter Greeter
}

func NewMessage(hi string) Message {
	return Message(hi)
}

func NewGreeter(m Message) Greeter {
	return Greeter{Mess: m}
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (g Greeter) Greet() Message {
	return g.Mess
}

func (e Event) start() {
	msg := e.Greeter.Greet()
	fmt.Println("msg: ", msg)
}

func main() {
	event, err := InitializeEvent("hi there, how are you?")
	if err != nil {
		log.Fatal("Can't injection")
	}
	event.start()
}
