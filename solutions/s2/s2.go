package main

import (
	"fmt"
	"sync"
	"time"
)

// Функция для измерения среднего времени выполнения
func timeIt(f func([]int), data []int, count int) time.Duration {
	arr := make([]time.Duration, 0, count)

	for i := 1; i <= count; i++ {
		start := time.Now()
		f(data)
		dur := time.Since(start)
		arr = append(arr, dur)
	}

	var r time.Duration

	for _, v := range arr {
		r += v
	}

	return r / time.Duration(count)
}

// Главная функция для запуска модуля
func main() {
	// Генератор массива чисел
	count := 1
	nums := make([]int, 0, 1000)
	for i := 1; i <= cap(nums); i++ {
		nums = append(nums, i)
	}

	// nums := []int{2, 4, 6, 7, 10}
	// count := 1

	// dur0 := timeIt(CalcSqrsNotComp, nums, count)
	dur1 := timeIt(CalcSqrs, nums, count)
	dur2 := timeIt(CalcSqrs2, nums, count)
	dur3 := timeIt(CalcSqrs3, nums, count)

	// // Средние значения за 100 повторений по 1000 элементов
	// fmt.Println("not concurrency", dur0)    // 1.4s
	fmt.Println("concurrency first", dur1)  // 2.3ms
	fmt.Println("concurrency second", dur2) // 1.4s
	fmt.Println("concurrency third", dur3)  // 1.8ms

}

// Контрольная группа
func CalcSqrsNotComp(nums []int) {
	r := make([]int, 0, len(nums))
	for _, v := range nums {
		r = append(r, v*v)

		// Для демонстрации работы конкурентности
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println(r)
}

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

	var result int
	for _, i := range r {
		result += i
	}

	fmt.Println(result)

	// fmt.Println(r)
}

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

	// var mux sync.Mutex

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

	var result int
	for _, i := range r {
		result += i
	}

	fmt.Println(result)

	// fmt.Println(r)
}

// Исправленная версия, где мы используем 2 горутины
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
	// Пока что думаю над решением через каналы.
	time.Sleep(100 * time.Microsecond)

	var result int
	for _, i := range r {
		result += i
	}

	fmt.Println(result)
}
