package feistel

func cloneSlice[T any](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)
	return b
}

func reverseSlice[T any](a []T) {
	i, j := 0, (len(a) - 1)
	for i < j {
		a[i], a[j] = a[j], a[i]
		i, j = i+1, j-1
	}
}

func xor[T Unsigned](a, b T) T {
	return a ^ b
}
