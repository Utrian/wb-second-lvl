package main

// Реализовать быструю сортировку массива (quicksort) встроенными методами
// языка.

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	quicksort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice)
	fmt.Println()
}

func quicksort(sl []int) []int {
	// Точка выхода из рекурсии, либо из функции
	// если изначально передан список с 0 или 1 элементом
	if len(sl) < 2 {
		return sl
	}

	// Назначаем границы сортировки
	left, right := 0, len(sl)-1

	// Выбираем опорный элемент
	pivot := rand.Int() % len(sl)
	// Переместим опорный элемент в конец слайса
	sl[pivot], sl[right] = sl[right], sl[pivot]

	// Начинаем сортировку элементов в левую часть
	// от опорного элемента
	for i := range sl {
		// Если текущий элемент меньше опорного, то
		// меняем его местами с левым граничным элементом
		if sl[i] < sl[right] {
			sl[left], sl[i] = sl[i], sl[left]
			left++
		}
	}

	sl[left], sl[right] = sl[right], sl[left]

	// На данный момент мы получили слайс следующего вида:
	// [
	//	неупорядоченные элементы, которые меньше опорного;
	//  опорный элемент и равные ему;
	//  неупорядоченные элементы, которые больше опорного;
	// ]

	// Теперь мы повторяем сортировку для каждой части
	// с неупорядоченными элементами. Пока длина части
	// не станет равна 1.
	quicksort(sl[:left])
	quicksort(sl[left+1:])

	return sl
}

func generateSlice(size int) []int {

	slice := make([]int, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		slice[i] = r.Intn(999) - r.Intn(999)
	}
	return slice
}
