package slice

import (
	"fmt"
	"testing"
)

func TestReviseSlice(t *testing.T) {
	sli, err := ReviseSlice([]int{1, 2, 3, 4, 5}, 2, 6)
	fmt.Println(sli, err)
}
func TestReviseSlice2(t *testing.T) {
	sli, err := ReviseSlice([]string{"a", "b", "c", "d"}, 1, "e")
	fmt.Println(sli, err)
}
