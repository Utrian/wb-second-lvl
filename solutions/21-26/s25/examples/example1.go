package examples

import "time"

// Используем time.After
// Идея в том, что на 9 строке блокируется процесс выполнения
// из-за того что мы ожидаем данные из канала, как только
// time.After отправляет данные - разблокировка спадает.
func Sleep1(d time.Duration) {
	<-time.After(d)
}

// time.After сам по себе может применяться в select
// чтобы закрывать слишком долгие горутины

// select {
// case chan <- num:
// 	...
// case <- time.After(time.Nanosecond * 500):
//     fmt.Println("Finish him!")
//     return
// }
