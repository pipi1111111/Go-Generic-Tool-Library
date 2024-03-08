package slice

import "fmt"

func ErrIndexOutRange(length int, index int) error {
	return fmt.Errorf("下标超出范围，长度%v，下标%v", length, index)
}
