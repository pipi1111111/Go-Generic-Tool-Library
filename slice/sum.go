package slice

type Num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func SumSlice[T Num](sli []T) T {
	var sum T
	for _, v := range sli {
		sum = sum + v
	}
	return sum
}
