package slice

import (
	"fmt"
	"testing"
)

func TestMergeSlice(t *testing.T) {
	sli := MergeSlice([]int{1, 3, 5, 7}, []int{2, 4, 6, 8, 1, 3})
	fmt.Println(sli)
}
