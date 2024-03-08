package slice

// AddSlice 向Slice特定位置增加元素
func AddSlice[T comparable](sli []T, index int, res T) ([]T, error) {
	length := len(sli)
	if length < 0 || length < index {
		return sli, ErrIndexOutRange(length, index)
	}
	var zero T
	sli = append(sli, zero)
	for i := length; i > index; i-- {
		sli[i] = sli[i-1]
	}
	sli[index] = res
	return sli, nil
}
