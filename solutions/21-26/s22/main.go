package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a, b, значение которых > 2^20.

func main() {
	// Определимся с типом данных - будем использовать пакет
	// math/big. Чтобы пользоваться возможностями данного пакеты
	// мы должны определять переменные с помощью него же:
	a := big.NewInt(25e10)  // km/s
	b := big.NewInt(299792) // light year
	// если бы b (int64) была > 9.3e18, то объявлять
	// переменную следовало бы так:
	a = new(big.Int)
	a.SetString("10000000000000000000", 10) // 10 - т.к. десятичное значение.

	bn := bigNums{a, b}

	fmt.Println(bn.mult()) // 2997920000000000000000000
	fmt.Println(bn.div())  // 33356460479265
	fmt.Println(bn.add())  // 10000000000000299792
	fmt.Println(bn.sub())  // 9999999999999700208
}

type bigNums struct {
	a *big.Int
	b *big.Int
}

func (bn *bigNums) mult() *big.Int {
	r := new(big.Int)
	return r.Mul(bn.a, bn.b)
}

func (bn *bigNums) div() *big.Int {
	r := new(big.Int)
	return r.Div(bn.a, bn.b)
}

func (bn *bigNums) add() *big.Int {
	r := new(big.Int)
	return r.Add(bn.a, bn.b)
}

func (bn *bigNums) sub() *big.Int {
	r := new(big.Int)
	return r.Sub(bn.a, bn.b)
}
