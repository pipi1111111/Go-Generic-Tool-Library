package slice

import (
	"fmt"
	"testing"
)

func TestAddSlice(t *testing.T) {
	sli, err := AddSlice([]int{1, 2, 3, 4, 5}, 1, 12)
	fmt.Println(sli, err)
}
func TestAddSlice2(t *testing.T) {
	sli, err := AddSlice([]string{"a", "b", "c", "d"}, 2, "e")
	fmt.Println(sli, err)
}
