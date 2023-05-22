package main

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	i int
}

func main() {
	doCount(2)
}

// В течение отведенного времени конкурентно увеличивает счетчик
// Мьютекс использовал для предотвращения эффекта гонки.
func doCount(sec int) {
	var c counter

	start := time.Now()

	var now time.Duration
	et := time.Duration(sec) * time.Second

	var mux sync.Mutex

	for now < et {
		go func() {
			mux.Lock()
			c.i++
			mux.Unlock()
		}()

		now = time.Since(start)
	}

	fmt.Println(c.i)
}

// Интересно узнать, если тут подводные камни, например,
// то что мы запускаем слишком много горутин. Возмножно
// это контролируется под капотом.
