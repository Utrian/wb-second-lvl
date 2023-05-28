package main

import (
	"github.com/wb-second-lvl/solutions/11-20/s11/examples"
)

// Реализовать пересечение двух неупорядоченных множеств.

// Пересечения множеств
func main() {
	// Пример с реализацией множеств через map[string]bool
	// На место string можно поставить и другие типы.
	examples.RunEx1()

	// Пример как и первый, только вместо bool используем
	// пустые структуры, и вывод в виде слайса, что
	// позволяет тратить меньше памяти.
	examples.RunEx2()

	// Уже готовая реализация множества, которая поддерживает
	// в том числе операцию пересечения
	// https://github.com/deckarep/golang-set
	examples.RunEx3()

	// И есть еще в стандартной библиотеке множество int
	// "golang.org/x/tools/container/intsets"
	// оно еще не является стабильным, но последнее обновление
	// было недавно так что в принципе помнить о нем стоит.
}
