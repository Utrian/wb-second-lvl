package main

import (
	"fmt"

	"github.com/wb-second-lvl/solutions/11-20/s13/examples"
)

// Свап значений двух переменных без использования третей;
func main() {
	x1, y1 := 1, 2
	x1, y1 = examples.DontWorkHard(x1, y1)
	fmt.Println(x1, y1)

	x2, y2 := 1, 2
	x2, y2 = examples.WorkLittle(x2, y2)
	fmt.Println(x2, y2)

	x3, y3 := 1, 2
	x3, y3 = examples.SwapArith1(x3, y3)
	fmt.Println(x3, y3)

	x4, y4 := 1, 2
	x4, y4 = examples.SwapArith2(x4, y4)
	fmt.Println(x4, y4)

	x5, y5 := 1, 2
	x5, y5 = examples.SwapBit(x5, y5)
	fmt.Println(x5, y5)
}
