package main

import "fmt"

// Удалить i-ый элемент из слайса.
func main() {
	sl := []int{0, 0, 0, 1, 0, 0, 0}
	fmt.Println(removeItem(sl, 3))
}

// pos - начинается от 0
func removeItem(sl []int, pos int) []int {
	r := make([]int, 0, len(sl)-1)

	// Берем нужные участки
	r = append(r, sl[:pos]...)
	r = append(r, sl[pos+1:]...)

	return r
}
