package main

// Реализовать все возможные способы остановки выполнения горутины.

// https://pkg.go.dev/gopkg.in/tomb.v2
// https://habr.com/ru/articles/141536/
// Следует перечитать
// Пакет tomb позволяет завершать горутину чисто

func main() {
	// examples.RunEx1(3) // пример с выделенным каналом quit и select
	// examples.RunEx2(3) // пример с закрытием канала и select
	// examples.RunEx3(3) // пример с закрытием канала и for := range
	// examples.RunEx4(3) // пример c закрытием канала и проверкой value, ok := <-ch
	// examples.RunEx5()  // пример с закрытием через внешний сигнал (ctrl + c)
	// examples.RunEx6() // пример с пакетом tomb
}
