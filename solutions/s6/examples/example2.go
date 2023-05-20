package examples

import (
	"fmt"
	"sync"
)

// Аналог первого примера, но с одним каналом;

func RunEx2(count int) {
	ch := make(chan string)

	// Используем WaitGroup чтобы дождаться выполнения горутин;
	var wg sync.WaitGroup
	wg.Add(2)

	go sing2(ch, &wg)

	go func(c int) {
		defer wg.Done()

		for i := 0; i < c; i++ {
			fmt.Println(<-ch)
		}
		close(ch) // сигнал о выходе из горутины
	}(count)

	wg.Wait()
}

func sing2(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	w := []string{"Ели", "мясо", "мужики", "пивом", "запивали"}

	for _, v := range w {
		select {
		case ch <- v:
			continue
		case <-ch: // выход из горутины sing2 происходит тут;
			return
		}
	}
}
