package slice

// DeleteSlice 删除index处的元素
func DeleteSlice[T comparable](sli []T, index int) ([]T, error) {
	length := len(sli)
	if length < 0 || index > length {
		return sli, ErrIndexOutRange(length, index)
	}
	for i := index; i < length; i++ {
		sli[i] = sli[index+1]
	}
	sli = sli[:length-1]
	return sli, nil
}
