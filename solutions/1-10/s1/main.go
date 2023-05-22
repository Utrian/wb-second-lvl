package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры
// Human (аналог наследования)

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
