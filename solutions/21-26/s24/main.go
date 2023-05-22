package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(6, -1)

	fmt.Println(Distance(p1, p2)) // 5.830951894845301
	// Корректность данных можно проверить тут
	// https://mnogoformul.ru/rasstoyanie-mezhdu-tochkami?result=384
}

// Определим структуру Pointer
type Point struct {
	x float64
	y float64
}

// И конструктор для нее
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Используем библиотеку math для реализации формулы
// расчета расстояния между точками.
func Distance(p1, p2 *Point) float64 {
	return math.Sqrt(
		math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2),
	)
}
