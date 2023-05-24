package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.

func main() {
	sUnique := "ajnlg,m1234t"
	sNotUnique := "daspfopIDHADFOHpsadfo"

	fmt.Println(isUnique(sUnique))
	fmt.Println(isUnique(sNotUnique))
}

func isUnique(s string) bool {
	s = strings.ToLower(s)

	var v void
	counter := make(map[rune]void)

	// Начинаем обход строки
	for _, char := range s {
		_, ok := counter[char]

		// Если символ уже есть - заканчиваем выполнение
		if ok {
			return false
		}
		// Если нет, добавляем в счетчик
		counter[char] = v
	}
	return true
}

type void interface{}
