package wpool

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type ExecFn func(n int)
type GetData func(ch chan int)

// Создадим Worker Pool для работы с int
func WorkerPoolInt(wcount int, genInt GetData, fn ExecFn) {
	ctx, cancel := context.WithCancel(context.Background())

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)

	chInt := make(chan int, wcount)

	var wg sync.WaitGroup

	// Запускаем установленное количество воркеров
	for i := 0; i < wcount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, fn, chInt)
		}()
	}

	// Запуск функции по генерации данных
	go genInt(chInt)

	// Точка начала завершения программы
	<-gracefulShutdown
	fmt.Println("Начинаем закрытие.")

	// Закрываем контекст чтобы воркеры завершили работу
	// и дожидаемся их. Данный способ завершения позволяет
	// безопасно дождаться выполнения горутин.
	// Есть еще вариант через закрытие канала - в воркере
	// мы бы проверяли канал на открытость, и при отрицательном
	// результате закрывали бы воркера. Но в таком случае
	// есть немалая вероятность, что при закрытии канала
	// еще будут существовать горутины которые пишут в канал,
	// что приведет к вызову паники.
	cancel()
	wg.Wait()

	fmt.Println("Программа завершила работу.")
}

// Сами воркеры. Чтобы воркер закончил работу мы должны
// закрыть контекст.
func worker(ctx context.Context, fn ExecFn, chInt <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Воркер закончил работу.")
			return
		case value := <-chInt:
			fn(value)
		}
	}
}
