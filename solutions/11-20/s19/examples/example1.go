package examples

// Решение через конкатенацию строк.
// Не самая дешевая реализация, т.к. при каждой конкатенации
// мы создаем новый объект.
func ReverseItConcat(s string) (r string) {
	for _, letter := range s {
		r = string(letter) + r
	}
	return
}
