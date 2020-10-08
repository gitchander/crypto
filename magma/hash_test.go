package magma

import (
	"bytes"
	"crypto/cipher"
	"testing"
)

func TestHash(t *testing.T) {

	key := []byte{
		0x81, 0x82, 0x83, 0x84, 0x85, 0xB6, 0x87, 0xCC,
		0x89, 0x8a, 0x11, 0x8c, 0x8d, 0x8e, 0x8f, 0x80,
		0xd1, 0xd2, 0xd3, 0xd4, 0xef, 0xd6, 0x90, 0xd8,
		0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xde, 0xdf, 0x01,
	}

	table, err := NewTable(table2)
	if err != nil {
		t.Error(err)
		return
	}

	block, err := NewBlockCipher(table, key)
	if err != nil {
		t.Error(err)
		return
	}

	h := NewHash(block)

	r := newRand()

	for i := 0; i < 1000; i++ {

		src, h1 := hashSample(block, r)

		h.Reset()
		h.Write(src)
		h2 := h.Sum(nil)

		if bytes.Compare(h1, h2) != 0 {
			t.Error("wrong hash")
			return
		}
	}
}

func hashSample(block cipher.Block, r randomer) (src, hash []byte) {

	var fullBlocks = func(x int) int {
		n := x / BlockSize
		if x > n*BlockSize {
			n++
		}
		return n * BlockSize
	}

	const n = 1000
	m := r.Intn(n) + 1

	bs := make([]byte, fullBlocks(m))
	r.FillBytes(bs[:m])
	src = bs[:m]

	hash = make([]byte, BlockSize)

	for len(bs) > 0 {

		safeXORBytes(hash, hash[:BlockSize], bs[:BlockSize])

		for i := 0; i < 16; i++ {
			block.Encrypt(hash, hash)
		}

		bs = bs[BlockSize:]
	}

	return
}
