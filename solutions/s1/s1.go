package main

import "fmt"

type Human struct {
	song string
}

func (h Human) sing() {
	fmt.Println(h.song)
}

type Action struct {
	Human
}

func newAction(song string) Action {
	var a Action
	a.song = song

	return a
}

func main() {
	Human{song: "Ели мясо мужики"}.sing()

	newAction("Пивом запивали!").sing()
}
