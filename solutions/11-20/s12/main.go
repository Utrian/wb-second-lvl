package main

import (
	"fmt"

	"github.com/wb-second-lvl/solutions/11-20/s12/examples"
)

// Создать множество из последовательности строк

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}

	r := examples.Set(s)
	fmt.Println(r)

	r2 := examples.Set2(s)
	fmt.Println(r2)
}
