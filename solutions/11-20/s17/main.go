package main

import (
	"fmt"
)

// Реализовать бинарный поиск встроенными методами языка.

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	num := sl[5]

	fmt.Println(BinarySearch(sl, num))
}

func BinarySearch(a []int, search int) (result int) {
	mid := len(a) / 2

	switch {
	case len(a) == 0: // элемент не найден
		result = -1
	case a[mid] > search: // элемент больше срединного => заходим в
		// рекурсию передав первую половину слайса;
		result = BinarySearch(a[:mid], search)
	case a[mid] < search: // элемент меньше срединного => заходим в
		// рекурсию передав вторую половину слайса;
		result = BinarySearch(a[mid+1:], search)
		if result >= 0 { // если элемент был найден, то нужно добавить
			result += mid + 1 // единицу, т.к. мы передавали mid+1
		}
	default: // если искомый элемент оказался срединным
		result = mid
	}

	return
}
