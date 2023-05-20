package examples

import (
	"fmt"
	"sync"
)

// Использование отдельного канала (quit) для сообщения
// о закрытии - когда из него приходят данные, то это
// приводит к выходу из функции;

// Функция выводит установленное число слов песни - если
// количество слов меньше count, то будет выведена вся песня;
func RunEx1(count int) {
	ch := make(chan string)
	quit := make(chan string)

	// Используем WaitGroup чтобы дождаться выполнения горутин;
	var wg sync.WaitGroup
	wg.Add(2)

	go func(c int) {
		defer wg.Done()

		for i := 0; i < c; i++ {
			fmt.Println(<-ch)
		}
		quit <- "" // сигнал о выходе из горутины
	}(count)

	go sing(ch, quit, &wg)

	wg.Wait()
}

func sing(ch, quit chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	w := []string{"Ели", "мясо", "мужики", "пивом", "запивали"}

	for _, v := range w {
		select {
		case ch <- v:
			continue
		case <-quit: // выход из горутины sing происходит тут;
			return
		}
	}
}
