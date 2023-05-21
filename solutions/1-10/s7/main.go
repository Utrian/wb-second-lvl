package main

import (
	"fmt"
	"sync"
)

// Конкурентная запись данных в map
func main() {
	fmt.Println("map:", do(10))
}

// Конкурентный расчет квадратов диапазона чисел и
// запись результатов в map.
func do(count int) map[int]int {
	r := make(map[int]int, count)

	// Используем мьютекс для лечения гонки
	// И WG для завершения горутин.
	var mux sync.Mutex
	var wg sync.WaitGroup

	wg.Add(count)

	for i := 1; i <= count; i++ {
		go func(n int) {
			defer wg.Done()

			sqrt := n * n // Мы выполняем расчеты даже
			// если карта залочена (можно было совершить ошибку
			// и вычислять внутри блока lock-unlock)

			mux.Lock()
			r[n] = sqrt
			mux.Unlock()
		}(i)
	}

	wg.Wait()

	return r
}
