package slice

import (
	"fmt"
	"testing"
)

func TestDeleteSlice(t *testing.T) {
	newSlice, err := DeleteSlice([]int{1, 2, 3, 4, 5}, 3)
	fmt.Println(newSlice, err)
}
func TestDeleteSlice2(t *testing.T) {
	newSlice, err := DeleteSlice([]string{"a", "b", "c", "d"}, 2)
	fmt.Println(newSlice, err)
}
