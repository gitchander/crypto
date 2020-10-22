package magma

import (
	"bytes"
	"testing"
)

func TestCipherRand(t *testing.T) {

	r := newRand()

	key := make([]byte, KeySize)

	for i := 0; i < 100; i++ {

		r.FillBytes(key)

		b, err := NewCipher(key)
		if err != nil {
			t.Error(err)
			return
		}

		var (
			plaintext    = make([]byte, b.BlockSize())
			ciphertext   = make([]byte, b.BlockSize())
			plaintextDec = make([]byte, b.BlockSize())
		)

		for j := 0; j < 1000; j++ {

			r.FillBytes(plaintext)

			b.Encrypt(ciphertext, plaintext)
			if bytes.Compare(plaintext, ciphertext) == 0 {
				t.Error("Encrypt compare true")
				return
			}

			b.Decrypt(plaintextDec, ciphertext)
			if bytes.Compare(plaintext, plaintextDec) != 0 {
				t.Error("Decrypt compare false")
				return
			}
		}
	}
}

func TestSamples(t *testing.T) {

	key := []byte{
		0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88,
		0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f, 0x80,
		0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xd8,
		0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xde, 0xdf, 0xd0,
	}

	type Sample struct {
		RT         ReplaceTable
		Key        []byte
		Plaintext  []byte
		Ciphertext []byte
	}

	var samples = []Sample{
		{
			RT:         RT2,
			Key:        key,
			Plaintext:  []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			Ciphertext: []byte{0xce, 0x5a, 0x5e, 0xd7, 0xe0, 0x57, 0x7a, 0x5f},
		},
		{
			RT:         RT2,
			Key:        key,
			Plaintext:  []byte{0xf1, 0xf2, 0xf3, 0xf4, 0xf5, 0xf6, 0xf7, 0xf8},
			Ciphertext: []byte{0xd0, 0xcc, 0x85, 0xce, 0x31, 0x63, 0x5b, 0x8b},
		},
	}

	for _, sample := range samples {

		b, err := NewCipherRT(sample.RT, sample.Key)
		if err != nil {
			t.Error(err)
			return
		}

		ciphertext := make([]byte, b.BlockSize())
		b.Encrypt(ciphertext, sample.Plaintext)
		if bytes.Compare(ciphertext, sample.Ciphertext) != 0 {
			t.Error("Encrypt error: not compare")
		}

		plaintext := make([]byte, b.BlockSize())
		b.Decrypt(plaintext, ciphertext)
		if bytes.Compare(plaintext, sample.Plaintext) != 0 {
			t.Error("Decrypt error: not compare")
		}
	}
}

func TestTables(t *testing.T) {

	r := newRand()

	type Replacers struct {
		r1 replacer
		r2 replacer
	}

	rrs := []Replacers{
		{
			makeReplaceTable8x16(RT1),
			makeReplaceTable4x256(RT1),
		},
		{
			makeReplaceTable8x16(RT2),
			makeReplaceTable4x256(RT2),
		},
	}

	for _, rs := range rrs {
		for i := 0; i < 1000000; i++ {
			u := r.Uint32()
			if rs.r1.replace(u) != rs.r2.replace(u) {
				t.Error("wrong mix")
				return
			}
		}
	}
}
