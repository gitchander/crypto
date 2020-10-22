package magma

import (
	"crypto/cipher"
	"hash"
)

type digest struct {
	b        cipher.Block
	out      []byte
	src      []byte
	srcIndex int
}

func NewHash(b cipher.Block) hash.Hash64 {
	return &digest{
		b:   b,
		out: make([]byte, b.BlockSize()),
		src: make([]byte, b.BlockSize()),
	}
}

func (d *digest) nextFill() {

	safeXORBytes(d.out, d.out, d.src)

	for i := 0; i < 16; i++ {
		d.b.Encrypt(d.out, d.out)
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

	hash := make([]byte, d.b.BlockSize())

	safeXORBytes(hash, d.out, d.src)

	for i := 0; i < 16; i++ {
		d.b.Encrypt(hash, hash)
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

func (d *digest) Size() int {
	return d.b.BlockSize()
}

func (d *digest) BlockSize() int {
	return d.b.BlockSize()
}
