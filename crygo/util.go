package crygo

func safeXORBytes(dst, a, b []byte) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		dst[i] = a[i] ^ b[i]
	}
	return n
}

func duplicateBytes(a []byte) []byte {
	b := make([]byte, len(a))
	copy(b, a)
	return b
}

func fillBytes(bs []byte, b byte) {
	for i, _ := range bs {
		bs[i] = b
	}
}

func quoRem(x, y int) (q, r int) {
	q = x / y
	r = x - q*y
	return
}

func mod(x, y int64) int64 {
	t := x % y
	if t < 0 {
		t += y
	}
	return t
}
