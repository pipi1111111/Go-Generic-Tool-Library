package slice

// ReviseSlice 修改Slice特定位置的元素
func ReviseSlice[T comparable](sli []T, index int, res T) ([]T, error) {
	length := len(sli)
	if length < 0 || length < index {
		return sli, ErrIndexOutRange(length, index)
	}
	sli[index] = res
	return sli, nil
}
