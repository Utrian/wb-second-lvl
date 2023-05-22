package main

import (
	"fmt"

	"github.com/wb-second-lvl/solutions/11-20/s19/examples"
)

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	fmt.Println(examples.ReverseItConcat("Привет мой друг"))
	fmt.Println(examples.ReverseItRune("Привет мой друг"))
}
