package examples

// Вариант через умножение и деление.
func SwapArith2(x, y int) (int, int) {
	x = x * y
	y = x / y
	x = x / y

	return x, y
}
