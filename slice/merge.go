package slice

// MergeSlice 合并两个类型相同的sli(并集）
func MergeSlice[T comparable](sli ...[]T) []T {
	var newSlice []T
	for _, sli1 := range sli {
		for _, v2 := range sli1 {
			newSlice = append(newSlice, v2)
		}
	}
	newSlice = RemoveSameSlice(newSlice)
	return newSlice
}
