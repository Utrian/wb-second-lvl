package main

import "github.com/wb-second-lvl/solutions/1-10/s5/examples"

// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.

func main() {
	examples.RunEx1(1)

	examples.RunEx2([]int{1, 2, 3, 4, 5}, 2)
}
