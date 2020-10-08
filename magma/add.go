package magma

const maxUint32 = (1 << 32) - 1 // 2 ^ 32 - 1

// (a + b) mod (2 ^ 32)
func add_mod32_v1(a, b uint32) uint32 {
	return a + b
}

func add_mod32_v2(a, b uint32) uint32 {
	_a := maxUint32 - a
	if b > _a {
		return (b - _a - 1)
	}
	return (a + b)
}

func add_mod32_v3(a, b uint32) uint32 {
	var (
		A = int64(a)
		B = int64(b)
		C = int64(1 << 32) // 2^32
	)
	return uint32(mod(A+B, C))
}

// (a + b) mod (2^32 - 1)*
func add_mod32m1(a, b uint32) uint32 {
	da := maxUint32 - a
	if b > da {
		return (b - da)
	}
	return (a + b)
}

func add_mod32m1_sample(a, b uint32) uint32 {
	s := int64(a) + int64(b)
	if s < (1 << 32) {
		return uint32(s)
	}
	return uint32(s - (1 << 32) + 1)
}

// wrong variant!
func add_mod32m1_wrong(a, b uint32) uint32 {
	var (
		A = int64(a)
		B = int64(b)
		C = int64((1 << 32) - 1) // (2^32 - 1)
	)
	return uint32(mod(A+B-1, C) + 1)
}
