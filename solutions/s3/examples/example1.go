package examples

import (
	"fmt"
	"sync"
)

// Стартавая функция с генератором чисел и с примером
// последовательности из задания.
func RunEx1() {
	// Генератор
	nums := make([]int, 0, 1000)
	for i := 1; i <= cap(nums); i++ {
		nums = append(nums, i)
	}
	// nums := []int{2, 4, 6, 7, 10}
	sumSqrsNums(nums)
}

// Функция расчета суммы квадратов.
func sumSqrsNums(nums []int) {
	c := len(nums)

	// Используем WG для завершения горутин.
	var wg sync.WaitGroup
	// Т.к. мы зарание знаем количество чисел в последовательности
	// то и добавим все сразу в счетчик.
	wg.Add(c)

	// Переменная с результатом и мьютекс для избавления от гонки
	var res int
	var mux sync.Mutex

	for _, v := range nums {
		go func(n int) {
			defer wg.Done()

			mux.Lock()
			res += n * n
			mux.Unlock()
		}(v)
	}
	wg.Wait()

	fmt.Println(res)
}
