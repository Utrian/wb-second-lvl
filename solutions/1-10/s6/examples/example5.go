package examples

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Завершение горутины через внешний сигнал (ctrl + c)
func RunEx5() {
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)

	go loop()

	<-gracefulShutdown

	fmt.Println("End.")
}

func loop() {
	for {
		fmt.Println("Cycling.")
	}
}
