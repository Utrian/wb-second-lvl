package examples

import (
	"fmt"
	"time"
)

// Пример, если нужно работать с определенными данными
// определенное время.
func RunEx2(data []int, sec int) {
	start := time.Now()

	var now time.Duration
	et := time.Duration(sec) * time.Second

	ch := make(chan int)

	for _, v := range data { // главное отличие тут;
		now = time.Since(start)

		if now >= et { // и тут;
			break
		}

		go func(n int) {
			ch <- n
		}(v)

		fmt.Println(<-ch)
	}
}
