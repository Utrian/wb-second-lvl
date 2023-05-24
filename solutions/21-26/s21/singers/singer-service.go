package singers

import "fmt"

// У нас есть сервис чтобы выводить в stdout песни
// Этот сервис нельзя изменять

// Есть интерфейс, который умеет работать со строками
type Singer interface {
	SingString()
}

// Тут мы используем данный интерфейс, если структура
// его поддерживает.
func SingThis(s Singer) {
	s.SingString()
}

// Структура с методами, реализующая интерфейс singer
type SongString struct {
	text   string
	author string
}

// Методы интерфейса
func (s SongString) SingString() {
	fmt.Println(s.text)
}

// Конструктор для базовой структуры сервиса
func NewSongStr(author, text string) *SongString {
	return &SongString{
		text:   text,
		author: author,
	}
}

func GetSongData() (string, string) {
	return "Lil Nas X", "I'm gonna take my horse to the old town road"
}
