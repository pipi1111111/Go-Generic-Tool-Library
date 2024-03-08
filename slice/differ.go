package slice

// DifferSlice 两个切片的差集(A对B的差集）
func DifferSlice[T comparable](sliA, sliB []T) []T {
	var newSlice []T
	for i := 0; i < len(sliA); i++ {
		var count int = 0
		for j := 0; j < len(sliB); j++ {
			if sliA[i] == sliB[j] {
				count++
				break
			}
		}
		if count == 0 {
			newSlice = append(newSlice, sliA[i])
		}
	}
	return newSlice
}
