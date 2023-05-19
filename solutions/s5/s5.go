package main

import (
	"fmt"
	"time"
)

func main() {
	WorkNsec(1)

	WorkNsec2([]int{1, 2, 3, 4, 5}, 2)
}

// Пример с бесконечно-генерируемыми данными;
func WorkNsec(sec int) {
	start := time.Now()

	var now time.Duration
	et := time.Duration(sec) * time.Second

	ch := make(chan int)

	for i := 0; now < et; i++ {
		go func(n int) {
			time.Sleep(10 * time.Millisecond)
			ch <- n
		}(i)

		fmt.Println(<-ch)

		now = time.Since(start)
	}
}

// Пример, если нужно работать с определенными данными;
func WorkNsec2(data []int, sec int) {
	start := time.Now()

	var now time.Duration
	et := time.Duration(sec) * time.Second

	ch := make(chan int)

	for _, v := range data {
		now = time.Since(start)

		if now >= et {
			break
		}

		go func(n int) {
			ch <- n
		}(v)

		fmt.Println(<-ch)
	}
}
