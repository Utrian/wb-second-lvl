package examples

import (
	"fmt"
	"sync"
	"time"
)

// Исправленная версия example2, где мы используем 2 горутины
// и при этом мьютекс позволяет нам детерминировать вывод данных.

func CalcSqrs3(nums []int) {
	c := len(nums)
	r := make([]int, 0, c)

	ch := make(chan int, c)

	var mux sync.Mutex

	for _, v := range nums {
		go func(n int) {
			ch <- n * n

			// Для демонстрации работы конкурентности
			time.Sleep(1 * time.Millisecond)
		}(v)

		go func() {
			mux.Lock()
			r = append(r, <-ch)
			mux.Unlock()
		}()
	}

	// Позволяет дождаться всех горутин
	// Но это скорее костыль, чем решение, т.к. если горутина
	// не выполнится за этот сон, то и ее результат пропадет
	// с завершением функции.
	// Пока что думаю над решением через каналы - можно добавить
	// WaitGroup, но смысл тогда - если есть первое решение.
	time.Sleep(150 * time.Microsecond)

	// Расчет суммы квадратов для проверки корректности данных;
	// var result int
	// for _, i := range r {
	// 	result += i
	// }

	// fmt.Println(result)

	fmt.Println(r)
}
