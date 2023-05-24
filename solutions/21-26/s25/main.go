package main

import (
	"fmt"
	"time"

	"github.com/wb-second-lvl/solutions/21-26/s25/examples"
)

// Реализовать собственную функцию sleep.
func main() {
	fmt.Println("Начинаем работу")
	fmt.Println("Работаем...")
	examples.Sleep1(2 * time.Second)
	fmt.Println("Заканчиваем работу")
}
