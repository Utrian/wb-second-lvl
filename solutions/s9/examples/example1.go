package examples

import "fmt"

// Пример без конкурентности
func Conveyor1(nums []int) {
	l := len(nums)

	ch1 := make(chan int, l)
	ch2 := make(chan int, l)

	for _, v := range nums {
		ch1 <- v

		doubling(ch1, ch2)

		fmt.Println(<-ch2)
	}
}

func doubling(ch1, ch2 chan int) {
	ch2 <- (2 * <-ch1)
}
