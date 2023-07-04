package kalyna

import (
	"crypto/cipher"
	"fmt"
)

// Reference implementation of the Kalyna block cipher (DSTU 7624:2014),
// all block and key length variants.

// info:
// https://en.wikipedia.org/wiki/Kalyna_(cipher)
// https://uk.wikipedia.org/wiki/%D0%9A%D0%B0%D0%BB%D0%B8%D0%BD%D0%B0_(%D1%88%D0%B8%D1%84%D1%80)

// forked from:
// github.com/Roman-Oliynykov/Kalyna-reference

// Kalyna
type Config struct {
	BlockSize int // Block size in bytes.
	KeySize   int // Key size in bytes.
}

func (c Config) NewCipher(key []byte) (cipher.Block, error) {

	k, err := newKalynaContext(c.BlockSize, c.KeySize)
	if err != nil {
		return nil, err
	}

	// Init key
	{
		var (
			haveKeySize = len(key)
			wantKeySize = c.KeySize
		)
		if haveKeySize != wantKeySize {
			return nil, fmt.Errorf("kalyna: invalid key size: have %d, want %d",
				haveKeySize, wantKeySize)
		}

		ws := make([]uint64, c.KeySize/bytesPerWord)
		bytesToWords(ws, key)
		err = k.KeyExpand(ws)
		if err != nil {
			return nil, err
		}
	}

	nb := c.BlockSize / bytesPerWord

	b := &block{
		k:          k,
		blockSize:  c.BlockSize,
		plaintext:  make([]uint64, nb),
		ciphertext: make([]uint64, nb),
	}

	return b, nil
}

type block struct {
	k *kalynaContext

	blockSize  int // Block size in bytes.
	plaintext  []uint64
	ciphertext []uint64
}

var _ cipher.Block = &block{}

func (b *block) BlockSize() int {
	return b.blockSize
}

func (b *block) Encrypt(dst, src []byte) {

	checkBuffersSize(dst, src, b.blockSize)

	bytesToWords(b.plaintext, src)
	b.k.Encipher(b.plaintext, b.ciphertext)
	wordsToBytes(b.ciphertext, dst)
}

func (b *block) Decrypt(dst, src []byte) {

	checkBuffersSize(dst, src, b.blockSize)

	bytesToWords(b.ciphertext, src)
	b.k.Decipher(b.ciphertext, b.plaintext)
	wordsToBytes(b.plaintext, dst)
}

func checkBuffersSize(dst, src []byte, blockSize int) {
	if len(src) < blockSize {
		panic("kalyna: input not full block")
	}
	if len(dst) < blockSize {
		panic("kalyna: output not full block")
	}
}
