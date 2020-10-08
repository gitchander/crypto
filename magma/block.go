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

type blockCipher struct {
	xs [8]uint32
	n  [2]uint32
	t  Table
}

func NewBlockCipher(t Table, key []byte) (cipher.Block, error) {

	if len(key) != KeySize {
		return nil, ErrorKeyLen
	}

	bc := &blockCipher{t: t}

	xs := bc.xs[:8]
	for i := range xs {
		xs[i] = byteOrder.Uint32(key)
		key = key[4:]
	}

	return bc, nil
}

func (blockCipher) BlockSize() int {
	return BlockSize
}

func (bc *blockCipher) Encrypt(dst, src []byte) {

	if len(src) < BlockSize {
		panic("crygo: input not full block")
	}

	if len(dst) < BlockSize {
		panic("crygo: output not full block")
	}

	encryptBlock(bc.xs[:8], bc.t, bc.n[:2], dst, src)
}

func (bc *blockCipher) Decrypt(dst, src []byte) {

	if len(src) < BlockSize {
		panic("crygo: input not full block")
	}

	if len(dst) < BlockSize {
		panic("crygo: output not full block")
	}

	decryptBlock(bc.xs[:8], bc.t, bc.n[:2], dst, src)
}
