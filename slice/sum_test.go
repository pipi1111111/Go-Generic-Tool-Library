package slice

import (
	"fmt"
	"testing"
)

func TestSumSlice(t *testing.T) {
	sum := SumSlice([]int{1, 2, 3, 4, 5})
	fmt.Println(sum)
}
