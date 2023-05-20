package examples

import (
	"fmt"
	"time"
)

// Решение через каналы, но данное решение не является конкурентным
// Данный пример был найден на просторах интернета, немного изменил
// его убрав лишний код.

// Он не является конкурентным, т.к. часть с получением инфы из канала
// тоже должна быть запущена как горутина - иначе все работает как
// с бутылочным горлышком - горутин отправляющих данные много (или самих
// данных в буф.канале), но принимает его всего один процесс.
func CalcSqrs2(nums []int) {
	c := len(nums)
	r := make([]int, 0, c)

	ch := make(chan int, c)

	go func() {
		for _, n := range nums {
			ch <- n * n

			// Для демонстрации работы конкурентности
			time.Sleep(1 * time.Millisecond)
		}
		close(ch)
	}()

	for {
		n, ok := <-ch
		if !ok {
			break
		}
		r = append(r, n)
	}

	// Расчет суммы квадратов для проверки корректности данных;
	// var result int
	// for _, i := range r {
	// 	result += i
	// }

	// fmt.Println(result)

	fmt.Println(r)
}
