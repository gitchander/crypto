package magma

import (
	"crypto/cipher"

	"github.com/gitchander/crypto/feistel"
)

const (
	c0 = 0x01010101
	c1 = 0x01010104
)

type streamCipher struct {
	b        cipher.Block
	we       *wordEncoder
	s        *feistel.RoundBlock[word]
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
		we:       defaultWordEncoder,
		s:        new(feistel.RoundBlock[word]),
		out:      make([]byte, size),
		outIndex: 0,
	}

	synEnc := make([]byte, size)
	b.Encrypt(synEnc, syn)
	sc.we.getBlock(synEnc, sc.s)

	sc.nextFill()

	return sc, nil
}

func (sc *streamCipher) nextFill() {

	s := sc.s

	s.R = s.R + c0
	s.L = word(add_mod32m1(uint32(s.L), c1))

	sc.we.putBlock(sc.out, sc.s)

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
