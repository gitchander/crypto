package magma

import "crypto/cipher"

// CFB - Cipher Feedback Mode

func NewCFBEncrypter(b cipher.Block, syn []byte) cipher.Stream {
	return newCFBCipher(b, syn, false)
}

func NewCFBDecrypter(b cipher.Block, syn []byte) cipher.Stream {
	return newCFBCipher(b, syn, true)
}

type cfb struct {
	b        cipher.Block
	out      []byte
	outIndex int

	decrypt bool
}

func newCFBCipher(b cipher.Block, syn []byte, decrypt bool) cipher.Stream {

	blockSize := b.BlockSize()

	if len(syn) != blockSize {
		panic("magma.newCFBCipher: SYN length must equal block size")
	}

	out := make([]byte, blockSize)
	copy(out, syn)

	return &cfb{
		b:        b,
		out:      out,
		outIndex: blockSize,
		decrypt:  decrypt,
	}
}

func (p *cfb) XORKeyStream(dst, src []byte) {

	for len(src) > 0 {

		if p.outIndex >= len(p.out) {
			p.b.Encrypt(p.out, p.out)
			p.outIndex = 0
		}

		n := safeXORBytes(dst, src, p.out[p.outIndex:])

		if p.decrypt {
			copy(p.out[p.outIndex:p.outIndex+n], src[:n])
		} else {
			copy(p.out[p.outIndex:p.outIndex+n], dst[:n])
		}

		src = src[n:]
		dst = dst[n:]
		p.outIndex += n
	}
}
