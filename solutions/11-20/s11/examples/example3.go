package examples

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
)

func RunEx3() {
	set1 := mapset.NewSet("cat", "gap", "mag")
	set2 := mapset.NewSet("dog", "gap", "flag")

	fmt.Println(set1.Intersect(set2)) // Set{gap}
}
