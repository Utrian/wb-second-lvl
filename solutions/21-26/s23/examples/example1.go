package examples

// Отсечение элемента с сохранением последовательности.
// pos - отсчет от 0
func RemoveItem(sl []int, pos int) []int {
	r := make([]int, 0, len(sl)-1)

	// Берем нужные участки
	r = append(r, sl[:pos]...)
	r = append(r, sl[pos+1:]...)

	return r
}
