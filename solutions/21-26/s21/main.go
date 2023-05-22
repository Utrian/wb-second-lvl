package main

import (
	"fmt"

	"github.com/wb-second-lvl/solutions/21-26/s21/examples"
)

// Реализовать паттерн «адаптер» на любом примере.

func main() {
	// Пример с интерфейсом.
	examples.RunExIntf()
	fmt.Println()

	// Пример с фукнцией.
	examples.RunExFunc()
}
