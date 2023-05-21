package examples

// В качестве множества будем использовать map[string]void
type void struct{}

var v void

func Set(s []string) map[string]void {
	r := map[string]void{}

	for _, value := range s {
		if _, ok := r[value]; !ok {
			r[value] = v
		}

	}
	return r
}
