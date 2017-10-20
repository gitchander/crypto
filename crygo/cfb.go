package crygo

import "crypto/cipher"

// CFB - Cipher Feedback Mode

func NewCFBEncrypter(block cipher.Block, syn []byte) cipher.Stream {
	return newCFBCipher(block, syn, false)
}

func NewCFBDecrypter(block cipher.Block, syn []byte) cipher.Stream {
	return newCFBCipher(block, syn, true)
}

type cfb struct {
	block    cipher.Block
	out      []byte
	outIndex int

	decrypt bool
}

func newCFBCipher(block cipher.Block, syn []byte, decrypt bool) cipher.Stream {

	blockSize := block.BlockSize()

	if len(syn) != blockSize {
		panic("crygo.newCFBCipher: SYN length must equal block size")
	}

	out := make([]byte, blockSize)
	copy(out, syn)

	return &cfb{
		block:    block,
		out:      out,
		outIndex: blockSize,
		decrypt:  decrypt,
	}
}

func (p *cfb) XORKeyStream(dst, src []byte) {

	for len(src) > 0 {

		if p.outIndex >= len(p.out) {
			p.block.Encrypt(p.out, p.out)
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
