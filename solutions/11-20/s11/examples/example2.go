package examples

import (
	"fmt"
	"sync"
)

func RunEx2() {
	set1 := map[string]void{
		"cat": v, "gap": v, "mag": v,
		"dog": v, "cap": v, "cup": v,
	}
	set2 := map[string]void{
		"cat": v, "gap": v, "mag": v,
		"flag": v, "drug": v, "mac": v,
	}

	r := Intersection2(set1, set2)
	fmt.Println(r)
}

type void struct{}

var v void

// Решение похожее на первое - с некоторыми изменениями:
// Выводим теперь слайс; к качестве значения ключей используем
// более легковесный struct{}. В целом данное решение
// должно занимать меньше памяти из-за этих двух изменений.
func Intersection2(set1, set2 map[string]void) []string {
	var smSet map[string]void
	var bgSet map[string]void

	if len(set1) <= len(set2) {
		smSet = set1
		bgSet = set2
	} else {
		smSet = set2
		bgSet = set1
	}

	r := []string{}

	var wg sync.WaitGroup
	wg.Add(len(smSet))

	var mux sync.Mutex

	for k := range smSet {
		go func(key string) {
			defer wg.Done()

			if _, ok := bgSet[key]; ok {
				mux.Lock()
				r = append(r, key)
				mux.Unlock()
			}
		}(k)
	}
	wg.Wait()

	return r
}
