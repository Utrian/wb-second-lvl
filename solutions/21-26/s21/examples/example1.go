package examples

import "fmt"

// Пример на основе интерфейса.
func RunExIntf() {
	s := Song{
		author: "Lil Nas X",
		text:   "I'm gonna take my horse to the old town road",
	}
	sr := NewSongRune("Lil Nas X", []rune("I'm gonna ride 'til I can't no more"))

	s.WhoAuthor() // Lil Nas X
	singThis(s)   // I'm gonna take my horse to the old town road

	sr.WhoAuthor() // Lil Nas X
	singThis(sr)   // I'm gonna ride 'til I can't no more
}

// У нас есть интерфейс для пения
type singer interface {
	Sing()
}

// Тут мы используем данный интерфейс, если структура
// его поддерживает.
func singThis(s singer) {
	s.Sing()
}

// У нас есть структура которая умеет работать со строками.
// При этом важным условием для адаптера является то, что
// модифицировать данную структуру и ее методы мы не можем;
type Song struct {
	text string
	// тут могут быть еще поля, но они, например, не мешают
	// работе с рунами, поэтому в дальнейшем мы их не
	// переопределяем:
	author string
}

// И она имеет метод Sing
func (s Song) Sing() {
	fmt.Println(s.text)
}

func (s *Song) WhoAuthor() {
	fmt.Println(s.author)
}

// Но нам теперь нужно уметь петь песни не из строк,
// а из последовательности рун - используем адаптер,
// в котором мы переопределим поля и методы для работы
// с рунами, чтобы поддерживать интерфейс.
type SongRune struct {
	Song
	text []rune
}

func (a SongRune) Sing() {
	fmt.Println(string(a.text))
}

func NewSongRune(author string, text []rune) SongRune {
	var sr SongRune
	sr.author = author
	sr.text = text

	return sr
}

// В итоге получилось, что мы адаптировали новые требования
// под работу со старым интерфейсом.
