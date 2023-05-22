package main

import (
	"github.com/wb-second-lvl/solutions/1-10/s9/examples"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.

func main() {
	examples.Conveyor1([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	// examples.Conveyor2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}
