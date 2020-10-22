package magma

import (
	"crypto/cipher"
)

const (
	c0 = 0x01010101
	c1 = 0x01010104
)

type streamCipher struct {
	b        cipher.Block
	s        [2]uint32
	out      []byte
	outIndex int
}

func NewStreamCipher(b cipher.Block, syn []byte) (cipher.Stream, error) {

	size := b.BlockSize()

	if len(syn) != size {
		return nil, ErrorSynLen
	}

	sc := &streamCipher{
		b:        b,
		out:      make([]byte, size),
		outIndex: 0,
	}

	synEnc := make([]byte, size)
	b.Encrypt(synEnc, syn)
	getTwoUint32(synEnc, &(sc.s))

	sc.nextFill()

	return sc, nil
}

func (sc *streamCipher) nextFill() {

	s := &(sc.s)

	s[0] = s[0] + c0
	s[1] = add_mod32m1(s[1], c1)

	putTwoUint32(sc.out, s)

	sc.b.Encrypt(sc.out, sc.out)
	sc.outIndex = 0
}

func (sc *streamCipher) XORKeyStream(dst, src []byte) {

	for len(src) > 0 {

		if sc.outIndex >= len(sc.out) {
			sc.nextFill()
		}

		n := safeXORBytes(dst, src, sc.out[sc.outIndex:])

		src = src[n:]
		dst = dst[n:]
		sc.outIndex += n
	}
}
