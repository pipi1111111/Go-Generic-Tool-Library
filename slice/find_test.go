package slice

import (
	"fmt"
	"testing"
)

func TestFindSlice(t *testing.T) {
	index, ok := FindSlice([]int{1, 2, 3, 4, 5}, 9)
	fmt.Println(index, ok)
}
