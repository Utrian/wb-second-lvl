package examples

// Делаем свап возможностями языка x, y = y, x
func WorkLittle(x, y int) (int, int) {
	x, y = y, x
	return x, y
}
