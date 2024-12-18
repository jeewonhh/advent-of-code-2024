package utils

func MapSum[T int](ts []T, fn func(T) T) T {
	var result T
	for _, t := range ts {
		result += fn(t)
	}
	return result
}

func Map[T any, S any](ts []T, f func(T) (S, error)) []S {
	r := make([]S, len(ts))
	for i, t := range ts {
		r[i], _ = f(t)
	}
	return r
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func RemoveAtIndex[T any](slice []T, i int) []T {
	if i < 0 || i >= len(slice) {
		panic("index out of range")
	}
	// Create a new slice without the element at index `i`
	new_slice := append([]T(nil), slice[:i]...)
	return append(new_slice, slice[i+1:]...)
}
