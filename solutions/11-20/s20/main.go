package main

import (
	"fmt"
	"strings"

	"github.com/wb-second-lvl/solutions/11-20/s19/examples"
)

func main() {
	fmt.Println(reverseTheseWords("snow dog sun"))
}

func reverseTheseWords(s string) string {
	// Для конкатенации используем strings.Builder
	var b strings.Builder

	// Разбиваем строку на слова в массиве - strings.Fields
	words := strings.Fields(s)

	for _, word := range words {
		// Переиспользуем код из предыдущего задания для реверса
		b.WriteString(examples.ReverseItRune(" " + word))
	}

	// Чтобы избавиться от лишнего пробела отсекаем его
	// с помощью strings.TrimSpace
	return strings.TrimSpace(b.String())
}
