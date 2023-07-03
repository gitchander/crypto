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
	we *wordEncoder

	encKeys []word
	decKeys []word

	rb *roundBlock
	rf roundFunc
}

func NewCipher(key []byte) (cipher.Block, error) {
	return NewCipherRT(RT1, key)
}

func NewCipherRT(rt ReplaceTable, key []byte) (cipher.Block, error) {

	err := checkValidReplaceTable(rt)
	if err != nil {
		return nil, err
	}

	we := defaultWordEncoder

	encKeys, err := expandKeyMagma(we, key)
	if err != nil {
		return nil, err
	}

	decKeys := cloneWords(encKeys)
	reverseWords(decKeys)

	b := &block{
		we:      we,
		encKeys: encKeys,
		decKeys: decKeys,
		rb:      new(roundBlock),
		rf:      roundFuncMagma(&rt),
	}

	return b, nil
}

func expandKeyMagma(we *wordEncoder, key []byte) ([]word, error) {

	if len(key) != KeySize {
		return nil, ErrorKeyLen
	}

	var xs [8]word
	for i := range xs {
		xs[i] = we.getWord(key)
		key = key[4:]
	}

	var ks []word

	for j := 0; j < 3; j++ {
		for i := 0; i < 8; i++ {
			ks = append(ks, xs[i])
		}
	}
	for i := 8; i > 0; i-- {
		ks = append(ks, xs[i-1])
	}

	return ks, nil
}

func roundFuncMagma(rt *ReplaceTable) roundFunc {
	return func(k, r word) word {
		s := k + r
		s = substituteMagma(rt, s)
		return wordShift11(s)
	}
}

func substituteMagma(rt *ReplaceTable, s0 word) (s1 word) {
	for i := 0; i < 8; i++ {
		var (
			shift = 4 * i
			j     = ((s0 >> shift) & 0xF)
		)
		s1 |= word(rt[i][j]) << shift
	}
	return s1
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

	b.we.getBlock(src, b.rb)
	encrypt(b.rb, b.encKeys, b.rf)
	b.we.putBlock(dst, b.rb)
}

func (b *block) Decrypt(dst, src []byte) {

	if len(src) < blockSize {
		panic("magma: input not full block")
	}

	if len(dst) < blockSize {
		panic("magma: output not full block")
	}

	b.we.getBlock(src, b.rb)
	decrypt(b.rb, b.decKeys, b.rf)
	b.we.putBlock(dst, b.rb)
}
