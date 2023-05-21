package examples

// Пример с битовыми операциями (XOR).
func SwapBit(x, y int) (int, int) {
	x = x ^ y
	y = y ^ x
	x = x ^ y

	return x, y
}
