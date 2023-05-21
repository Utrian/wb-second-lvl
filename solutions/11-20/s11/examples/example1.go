package examples

import (
	"fmt"
	"sync"
)

func RunEx1() {
	set1 := map[string]bool{
		"cat": true, "gap": true, "mag": true,
		"dog": true, "cap": true, "cup": true,
	}
	set2 := map[string]bool{
		"cat": true, "gap": true, "mag": true,
		"flag": true, "drug": true, "mac": true,
	}

	r := Intersection(set1, set2)
	fmt.Println(r)
}

// В качестве сета будем использовать мапу - для начала определим
// наименьшую из полученных мап (сетов) чтобы минимизировать
// количество проходов в цикле с проверками наличия ключей в
// обоих мапах.
// И т.к. итоговая структура не должна быть упорядоченная - используем
// горутины.
func Intersection(set1, set2 map[string]bool) map[string]bool {
	var smSet map[string]bool
	var bgSet map[string]bool

	if len(set1) <= len(set2) {
		smSet = set1
		bgSet = set2
	} else {
		smSet = set2
		bgSet = set1
	}

	r := make(map[string]bool, len(smSet))

	var wg sync.WaitGroup
	wg.Add(len(smSet))

	var mux sync.Mutex

	for k := range smSet {
		go func(key string) {
			defer wg.Done()

			if _, ok := bgSet[key]; ok {
				mux.Lock()
				r[key] = true
				mux.Unlock()
			}
		}(k)
	}
	wg.Wait()

	return r
}
