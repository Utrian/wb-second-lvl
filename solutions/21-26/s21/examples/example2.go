package examples

import "fmt"

// Пример на основе функций.
func RunExFunc() {
	singIt(stringToRune(getSong()))
}

// Еще один пример, но не знаю на сколько он корректен.
// Тут идея про те случаи, например, когда раньше мы получали
// в ответе XML и обрабатываели его, но стали получать
// JSON, и из-за невозможности изменить обрабатывающую
// функцию вводим функцию-адаптер.

// Есть функция возвращающая строки песни
func getSong() string {
	return "I'm gonna take my horse to the old town road"
}

// И есть функция поющая, но она принимает только rune
func singIt(rns []rune) {
	fmt.Println(string(rns))
}

// Чтобы подружить их будем использовать адаптер
func stringToRune(s string) []rune {
	return []rune(s)
}
