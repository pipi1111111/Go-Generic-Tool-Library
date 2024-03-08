package slice

// Same sliced的交集
func Same[T comparable](sli1, sli2 []T) []T {
	var newSlice []T
	for _, v1 := range sli1 {
		for _, v2 := range sli2 {
			if v1 == v2 {
				newSlice = append(newSlice, v1)
			}
		}
	}
	return newSlice
}
