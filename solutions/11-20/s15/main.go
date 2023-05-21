package main

import (
	"fmt"
	"strings"
)

// Откажемся от глобальной переменной и будем возвращать
// строку из функции.
func correctSomeFunc() string {
	var justString string

	v := createHugeString(1 << 10)

	// Крутим-вертим руны-стринги
	// Чтобы получить именно первые 100 символов требуется
	// преобразовать строку в слайс рун - коды символов
	// кодировки UTF-8. Получили коды - преобразуем их
	// в строковое представление - string может и в UTF8.
	// Если бы мы знали, что в строке будут только ASCII
	// символы, то можно было бы оставить изначальный вариант
	// т.к. 1 символ 1 байт.
	// В нашем случае мы используем руны для поддержки
	// символов от кириллицы весом в 2 байта, до иероглифов
	// весом в 3 байта.
	justString = string([]rune(v)[:100])
	l := len(string([]rune(v)[:100]))

	fmt.Println(justString, l)

	return justString
}

// Для создания строки используем strings.Builder, т.к.
// это база в конкатенации строк.
func createHugeString(size int) string {
	var b strings.Builder

	for i := 0; i < size; i++ {
		fmt.Fprint(&b, "кис ")
	}

	return b.String()
}

func main() {
	correctSomeFunc()
}
