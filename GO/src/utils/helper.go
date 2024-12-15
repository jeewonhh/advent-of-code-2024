package utils

func MapSum[T int](ts []T, fn func(T) T) T {
	var result T
	for _, t := range ts {
		result += fn(t)
	}
	return result
}
