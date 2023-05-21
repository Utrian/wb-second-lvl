package examples

import (
	"fmt"
	"time"
)

// Пример с бесконечно-генерируемыми данными;
// Если посмотреть на вывод, то выглядит так, что все
// данные выводятся последовательно.
func RunEx1(sec int) {
	// Время отсчета;
	start := time.Now()

	// Продолжительность работы в секундах;
	var now time.Duration
	et := time.Duration(sec) * time.Second

	ch := make(chan int)

	// Пока текущее время выполнения меньше et (end time)
	// функция работает.
	for i := 0; now < et; i++ {
		go func(n int) {
			ch <- n
		}(i)

		fmt.Println(<-ch)

		// Обновляем текущее время выполнения.
		now = time.Since(start)
	}
}
