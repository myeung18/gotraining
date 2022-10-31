package main

import "fmt"

type handler func(string)

// function type
func fireEvent(h handler) {
	h("fire in the hole")
}

func fireEvt(f func(string)) {
	f("anonymous")
}

type data struct {
	name string
	age  int
}

func (d *data) event(msg string) {
	fmt.Println(d.name, msg)
}

func event(msg string) {
	fmt.Println(msg)
}

func main() {
	d := data{
		name: "bill",
	}

	fireEvt(event)
	fireEvt(d.event)

	fireEvent(event)
	fireEvent(d.event)

	h1 := handler(event)
	h2 := handler(d.event)

	fireEvent(h1)
	fireEvent(h2)

	fireEvt(h1)
	fireEvt(h2)
}
