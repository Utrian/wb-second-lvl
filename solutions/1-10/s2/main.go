package main

import (
	"fmt"
	"time"

	"github.com/wb-second-lvl/solutions/1-10/s2/examples"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	// Генератор массива чисел
	count := 1
	nums := make([]int, 0, 1000)
	for i := 1; i <= cap(nums); i++ {
		nums = append(nums, i)
	}

	// nums := []int{2, 4, 6, 7, 10}
	// count := 1

	dur0 := timeIt(CalcSqrsNotConc, nums, count)
	dur1 := timeIt(examples.CalcSqrs, nums, count)
	dur2 := timeIt(examples.CalcSqrs2, nums, count)
	dur3 := timeIt(examples.CalcSqrs3, nums, count)

	// // Средние значения за 100 повторений по 1000 элементов
	fmt.Println("not concurrency", dur0)      // 1.4s
	fmt.Println("concurrency first", dur1)    // 2.3ms
	fmt.Println("'concurrency' second", dur2) // 1.4s
	fmt.Println("concurrency third", dur3)    // 1.8ms
}

func CalcSqrsNotConc(nums []int) {
	r := make([]int, 0, len(nums))
	for _, v := range nums {
		r = append(r, v*v)

		// Для демонстрации работы конкурентности
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println(r)
}

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
