package examples

import (
	"fmt"
	"sync"
)

// Пример с завершением горутины, которая бесконечно случает
// канал. Ее завершение происходит через закрытие канала.
func RunEx2(count int) {
	ch := make(chan int)

	// Используем WaitGroup для корректного завершения горутины.
	var wg sync.WaitGroup
	wg.Add(1)

	go printer(count, ch, &wg)

	for i := 0; i <= count; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}

func printer(count int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for s := range ch { // по закрытию канала завершается
		fmt.Println(s) // цикл for range и горутина закрывается;
	}
}
