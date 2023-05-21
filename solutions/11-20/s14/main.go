package main

import (
	"fmt"
	"reflect"
)

// Программа, способная различать типы положенные в interface{}
func main() {
	var x bool
	findOutType(x)
}

// Для решения понадобится type switches

// Поддерживает int, string, bool, chan
func findOutType(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Printf("This is %T.\n", t)
	case string:
		fmt.Printf("This is %T.\n", t)
	case bool:
		fmt.Printf("This is %T.\n", t)
	default:
		// Обработка паники при получении interface{}
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("This data type isn`t supported.")
			}
		}()

		if st := reflect.TypeOf(t).String(); st[:4] == "chan" {
			fmt.Printf("This is %T.\n", t)
		} else {
			fmt.Println("This data type isn`t supported.")
		}
	}
}
