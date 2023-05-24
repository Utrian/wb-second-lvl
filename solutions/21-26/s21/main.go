package main

import "github.com/wb-second-lvl/solutions/21-26/s21/singers"

// Реализовать паттерн «адаптер» на любом примере.

func main() {
	// Пусть singer-service будет "внешним" сервисом, который мы хотим
	// использовать в своей программе, где у нас есть "локальная"
	// структура, которую мы будем использовать с внешним сервисом

	// Создаем создадим структуру из внешнего сервиса для работы с его
	// интерфейсом, и посмотрим как она работает:
	s := singers.NewSongStr(singers.GetSongData())
	singers.SingThis(s) // I'm gonna take my horse to the old town road

	// Теперь создадим локальную структуру
	sr := singers.NewSongRune(singers.GetSongDataRns())

	// Чтобы она могла взаимодейстовать с интерфейсом - используем адаптер:
	srAdapter := singers.SongRuneAdapter{SongRune: &sr}

	// Теперь мы можем взаимодействовать с интерфейсом
	singers.SingThis(srAdapter) // I'm gonna ride 'til I can't no more
}
