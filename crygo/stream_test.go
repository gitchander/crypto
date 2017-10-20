package crygo

import (
	"bytes"
	"testing"
)

func TestStream(t *testing.T) {

	key := []byte{
		0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88,
		0x89, 0x8a, 0x11, 0x8c, 0x8d, 0x8e, 0x8f, 0x80,
		0xd1, 0xd2, 0xd3, 0xd4, 0xef, 0xd6, 0xd7, 0xd8,
		0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xde, 0xdf, 0x01,
	}

	table, err := NewTable(table1)
	if err != nil {
		t.Error(err)
		return
	}

	block, err := NewBlockCipher(table, key)
	if err != nil {
		t.Error(err)
		return
	}

	syn := []byte{0xA1, 0x09, 0xDC, 0x11, 0x73, 0x17, 0x04, 0x13}

	r := newRand()

	const n = 1000

	s1 := make([]byte, n)
	s2 := make([]byte, n)
	s3 := make([]byte, n)

	for i := 0; i < 10000; i++ {

		m := r.Intn(n) + 1
		s1 = s1[:m]
		s2 = s2[:m]
		s3 = s3[:m]

		r.FillBytes(s1)

		// Encrypt
		{
			sc, err := NewStreamCipher(block, syn)
			if err != nil {
				t.Error(err)
				return
			}

			sc.XORKeyStream(s2, s1)
		}

		// Decrypt
		{
			sc, err := NewStreamCipher(block, syn)
			if err != nil {
				t.Error(err)
				return
			}

			sc.XORKeyStream(s3, s2)
		}

		if bytes.Compare(s1, s3) != 0 {

			t.Error("not equal")
			return
		}
	}
}
