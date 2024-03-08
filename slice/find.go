package slice

func FindSlice[T comparable](sli []T, res T) (int, bool) {
	for k, v := range sli {
		if v == res {
			return k, true
		}
	}
	return 0, false
}
