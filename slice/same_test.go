package slice

import (
	"fmt"
	"testing"
)

func TestSame(t *testing.T) {
	sli := Same([]int{1, 2, 3, 4, 5}, []int{3, 4, 5})
	fmt.Println(sli)
}
func TestSame2(t *testing.T) {
	sli := Same([]string{"a", "b", "c", "e"}, []string{"c", "e", "f", "g"})
	fmt.Println(sli)
}
