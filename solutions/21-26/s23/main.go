package main

import (
	"fmt"

	"github.com/wb-second-lvl/solutions/21-26/s23/examples"
)

// Удалить i-ый элемент из слайса.
func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7}

	// Пример с созданием нового слайса
	new_sl1 := examples.RemoveItem(sl, 3)
	fmt.Println(new_sl1, cap(new_sl1)) // [1 2 3 5 6 7] 6

	new_sl1 = examples.RemoveItem(new_sl1, 3)
	fmt.Println(new_sl1, cap(new_sl1)) // [1 2 3 6 7] 5

	// Пример с заменой, смещением и изменением порядка
	new_sl2 := examples.RemoveItem2(sl, 3)
	fmt.Println(new_sl2, cap(new_sl2)) // [1 2 3 7 5 6] 7

	new_sl2 = examples.RemoveItem2(new_sl2, 3)
	fmt.Println(new_sl2, cap(new_sl2)) // [1 2 3 6 5] 7
}

// Как видим по capacity во втором варианте после
// изменения слайса он все еще ссылается на изначальный массив.
// Данную реализацию можно изменить чтобы оставить порядок
// но смысла в этом особого смысла нет (из-за недостатка выше)
