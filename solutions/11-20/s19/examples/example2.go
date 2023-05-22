package examples

// Решение через руны и их свап.
// Классическое перевертывание массива через два указателя.
func ReverseItRune(s string) string {
	rns := []rune(s)
	// Берем первый и последний элемент;
	// Пока первый элемент меньше последнего;
	// После свапа делаем шаг к середине;
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
