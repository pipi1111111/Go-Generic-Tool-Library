package slice

import (
	"fmt"
	"testing"
)

func TestRemoveSameSlice(t *testing.T) {
	sli := RemoveSameSlice([]int{1, 1, 3, 4, 5, 6, 7, 4})
	fmt.Println(sli)
}
