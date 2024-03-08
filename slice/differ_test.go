package slice

import (
	"fmt"
	"testing"
)

func TestDifferSlice(t *testing.T) {
	sli := DifferSlice([]int{1, 2, 3, 4, 7, 8}, []int{3, 4, 5, 6, 7})

	fmt.Println(sli)
}
func TestDifferSlice2(t *testing.T) {
	sli := DifferSlice([]string{"a", "b", "c", "e", "g", "h"}, []string{"c", "e", "f", "g"})

	fmt.Println(sli)
}
