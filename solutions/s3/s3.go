package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := make([]int, 0, 1000)
	for i := 1; i <= cap(nums); i++ {
		nums = append(nums, i)
	}
	// nums := []int{2, 4, 6, 7, 10}
	sumSqrsNums(nums)
}

func sumSqrsNums(nums []int) {
	c := len(nums)

	var wg sync.WaitGroup
	wg.Add(c)

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
