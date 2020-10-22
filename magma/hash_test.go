package magma

import (
	"bytes"
	"crypto/cipher"
	"testing"
)

func TestHash(t *testing.T) {

	rt := RT2

	key := []byte{
		0x81, 0x82, 0x83, 0x84, 0x85, 0xB6, 0x87, 0xCC,
		0x89, 0x8a, 0x11, 0x8c, 0x8d, 0x8e, 0x8f, 0x80,
		0xd1, 0xd2, 0xd3, 0xd4, 0xef, 0xd6, 0x90, 0xd8,
		0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xde, 0xdf, 0x01,
	}

	b, err := NewCipherRT(rt, key)
	if err != nil {
		t.Error(err)
		return
	}

	h := NewHash(b)

	r := newRand()

	for i := 0; i < 1000; i++ {

		src, h1 := hashSample(b, r)

		h.Reset()
		h.Write(src)
		h2 := h.Sum(nil)

		if bytes.Compare(h1, h2) != 0 {
			t.Error("wrong hash")
			return
		}
	}
}

func hashSample(b cipher.Block, r randomer) (src, hash []byte) {

	blockSize := b.BlockSize()

	var fullBlocks = func(x int) int {
		n := x / blockSize
		if x > n*blockSize {
			n++
		}
		return n * blockSize
	}

	const n = 1000
	m := r.Intn(n) + 1

	bs := make([]byte, fullBlocks(m))
	r.FillBytes(bs[:m])
	src = bs[:m]

	hash = make([]byte, blockSize)

	for len(bs) > 0 {

		safeXORBytes(hash, hash[:blockSize], bs[:blockSize])

		for i := 0; i < 16; i++ {
			b.Encrypt(hash, hash)
		}

		bs = bs[blockSize:]
	}

	return
}
