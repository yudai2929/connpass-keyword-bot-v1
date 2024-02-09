package slice

func Map[T1 comparable, T2 any](slice []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(slice))
	for i, s := range slice {
		result[i] = f(s)
	}
	return result
}
