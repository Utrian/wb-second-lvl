package main

import "fmt"

// Создать множество из последовательности строк
func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}

	r := set(s)
	fmt.Println(r)
}

// В качестве множества будем использовать map[string]void
// Так же как с предыдущим заданием можно добавить горутины

type void struct{}

var v void

func set(s []string) map[string]void {
	r := map[string]void{}

	for _, value := range s {
		if _, ok := r[value]; !ok {
			r[value] = v
		}
	}
	return r
}
