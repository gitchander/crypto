package magma

import (
	"crypto/cipher"
)

// GOST block cipher (Magma)
// GOST 28147-89

// info:
// https://en.wikipedia.org/wiki/GOST_(block_cipher)
// https://uk.wikipedia.org/wiki/%D0%93%D0%9E%D0%A1%D0%A2_28147-89

// Electronic Codebook - ECB

const (
	blockSize = 8 // Block size in bytes.
)

const KeySize = 32 // Key size in bytes.

type block struct {
	r  replacer
	xs [8]uint32
	n  [2]uint32
}

func NewCipher(key []byte) (cipher.Block, error) {
	return NewCipherRT(RT1, key)
}

func NewCipherRT(rt ReplaceTable, key []byte) (cipher.Block, error) {

	err := checkValidReplaceTable(rt)
	if err != nil {
		return nil, err
	}

	if len(key) != KeySize {
		return nil, ErrorKeyLen
	}

	var xs [8]uint32
	for i := range xs {
		xs[i] = byteOrder.Uint32(key)
		key = key[4:]
	}

	b := &block{
		r:  makeReplacer(rt),
		xs: xs,
	}

	return b, nil
}

func (block) BlockSize() int {
	return blockSize
}

func (b *block) Encrypt(dst, src []byte) {

	if len(src) < blockSize {
		panic("magma: input not full block")
	}

	if len(dst) < blockSize {
		panic("magma: output not full block")
	}

	encryptBlock(b.r, &(b.n), &(b.xs), dst, src)
}

func (b *block) Decrypt(dst, src []byte) {

	if len(src) < blockSize {
		panic("magma: input not full block")
	}

	if len(dst) < blockSize {
		panic("magma: output not full block")
	}

	decryptBlock(b.r, &(b.n), &(b.xs), dst, src)
}
