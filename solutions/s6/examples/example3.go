package examples

import (
	"fmt"
	"sync"
)

// Пример с завершением горутины, которая содержит бесконечный
// цикл. Ее завершение происходит через проверку канала на
// открытость - если возвращается false выполняется return.
func RunEx3(count int) {
	ch := make(chan int)

	// Используем WaitGroup для корректного завершения горутины.
	var wg sync.WaitGroup
	wg.Add(1)

	go printer2(count, ch, &wg)

	for i := 0; i <= count; i++ {
		ch <- i
	}
	close(ch) // закрываем канал
	wg.Wait()
}

func printer2(count int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		v, ok := <-ch // ok хранит true/false когда канал открыт/закрыт.
		if !ok {      // при выполнении условия - выход из горутины.
			return
		}
		fmt.Println(v)
	}
}
