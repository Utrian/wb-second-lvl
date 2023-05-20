package examples

import (
	"fmt"
	"sync"
)

// Пример с конкурентностью
func Conveyor2(nums []int) {
	l := len(nums)

	ch1 := make(chan int, l)
	ch2 := make(chan int, l)

	var wg sync.WaitGroup
	wg.Add(l * 2) // 2 - т.к. на один цикл 2 горутины;

	for _, v := range nums {
		ch1 <- v

		go func() {
			defer wg.Done()
			ch2 <- (2 * <-ch1)
		}()

		go func() {
			defer wg.Done()
			fmt.Println(<-ch2)
		}()
	}
	wg.Wait()
}
