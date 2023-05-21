package examples

import (
	"fmt"
	"sync"
	"time"
)

// Пример с использованием WaitGroup
// Результат функции будут квадраты чисел массива, но не упорядоченные;
// Мьютекс обязательно нужен, т.к. из-за гонки вывод разнится на больших
// данных - на примере из ТЗ из 5 чисел данные корректны, на 1000
// нет - проверено через неконкурентное сложение всех сумм квадратов
// и сравнение результата с результатом по формуле n(n+1)(2n+1)/6
// где n - последнее число в последовательности от 1.
func CalcSqrs(nums []int) {
	c := len(nums)
	r := make([]int, 0, c)

	var wg sync.WaitGroup
	wg.Add(c)

	var mux sync.Mutex

	for _, v := range nums {
		go func(n int) {
			defer wg.Done()

			mux.Lock()
			r = append(r, n*n)
			mux.Unlock()

			// Для демонстрации работы конкурентности
			time.Sleep(1 * time.Millisecond)
		}(v)
	}
	wg.Wait()

	// Расчет суммы квадратов для проверки корректности данных;
	// var result int
	// for _, i := range r {
	// 	result += i
	// }

	// fmt.Println(result)

	fmt.Println(r)
}
