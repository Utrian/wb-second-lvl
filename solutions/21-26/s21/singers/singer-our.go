package singers

import "fmt"

// В этом файле у нас есть другая базовая структура, которую
// мы хотим использовать с интерфейсом singer, но из-за
// разного типа данных нам придется реализовать адатер.
type SongRune struct {
	SongString
	text []rune
}

func (sr SongRune) ConvertToString() string {
	return string(sr.text)
}

func NewSongRune(author string, text []rune) SongRune {
	var sr SongRune
	sr.author = author
	sr.text = text

	return sr
}

func GetSongDataRns() (string, []rune) {
	return "Lil Nas X", []rune("I'm gonna ride 'til I can't no more")
}

// Создаем адаптер для связи текущей структуры с интерфейсом singers
type SongRuneAdapter struct {
	SongRune *SongRune
}

// Реализуем интерфейс singers
func (a SongRuneAdapter) SingString() {
	s := a.SongRune.ConvertToString()
	fmt.Println(s)
}
