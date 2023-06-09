package main

import "fmt"

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
// 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
// градусов. Последовательность в подмножноствах не важна.

func main() {
	t := []float64{
		-25.4, -27.0, 13.0, 19.0,
		15.5, 24.5, -21.0, 32.5,
	}
	gt := groupTemperatures(t)
	fmt.Println(gt)
}

// Группировка температорных колебаний с шагов в 10 градусов
func groupTemperatures(t []float64) map[int][]float64 {
	gt := map[int][]float64{}

	// Сначала получаем целочисленное значение от вещественного
	// числа, далее благодаря делению по модулю отнимаем лишнее
	// число для получения шага. Потом просто добавляем его в мапу.
	for _, v := range t {
		vInt := int(v)
		step := vInt - (vInt % 10)

		gt[step] = append(gt[step], v)
	}

	return gt
}
