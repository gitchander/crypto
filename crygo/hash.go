package crygo

import (
	"crypto/cipher"
	"hash"
)

type digest struct {
	block    cipher.Block
	out      []byte
	src      []byte
	srcIndex int
}

func NewHash(block cipher.Block) hash.Hash64 {
	return &digest{
		block: block,
		out:   make([]byte, block.BlockSize()),
		src:   make([]byte, block.BlockSize()),
	}
}

func (d *digest) nextFill() {

	safeXORBytes(d.out, d.out, d.src)

	for i := 0; i < 16; i++ {
		d.block.Encrypt(d.out, d.out)
	}

	fillBytes(d.src, 0)
	d.srcIndex = 0
}

func (d *digest) Write(src []byte) (count int, err error) {
	for len(src) > 0 {

		if d.srcIndex >= len(d.src) {
			d.nextFill()
		}

		n := len(src)
		if m := (len(d.src) - d.srcIndex); n > m {
			n = m
		}

		copy(d.src[d.srcIndex:d.srcIndex+n], src[:n])

		src = src[n:]
		d.srcIndex += n
		count += n
	}
	return
}

func (d *digest) checkSum() []byte {

	hash := make([]byte, d.block.BlockSize())

	safeXORBytes(hash, d.out, d.src)

	for i := 0; i < 16; i++ {
		d.block.Encrypt(hash, hash)
	}

	return hash
}

func (d *digest) Sum(in []byte) []byte {
	hash := d.checkSum()
	return append(in, hash...)
}

func (d *digest) Sum64() uint64 {
	hash := d.checkSum()
	return byteOrder.Uint64(hash)
}

func (d *digest) Reset() {
	fillBytes(d.out, 0)
	fillBytes(d.src, 0)
	d.srcIndex = 0
}

func (digest) Size() int {
	return BlockSize
}

func (digest) BlockSize() int {
	return BlockSize
}
