package main

import (
	"crypto/cipher"
	"fmt"
)

func NewCipher(key []byte) (cipher.Block, error) {
	ks, err := expandKey(key, bigEndian)
	if err != nil {
		return nil, err
	}
	return newBlockCipher(ks, RoundFuncXOR)
}

type blockCipher struct {
	we *wordEncoder
	fn *FeistelNetwork[Word]
	rb *RoundBlock[Word]
}

func newBlockCipher(ks []Word, rf RoundFunc[Word]) (cipher.Block, error) {
	we := bigEndian
	block := &blockCipher{
		we: we,
		fn: NewFeistelNetwork(ks, rf),
		rb: new(RoundBlock[Word]),
	}

	return block, nil
}

func (blockCipher) BlockSize() int {
	return blockSize
}

func (blockCipher) checkParams(dst, src []byte) error {
	if len(src) < blockSize {
		return fmt.Errorf("insufficient src size: have %d, want %d",
			len(src), blockSize)
	}
	if len(dst) < blockSize {
		return fmt.Errorf("insufficient dst size: have %d, want %d",
			len(dst), blockSize)
	}
	return nil
}

func (p *blockCipher) Encrypt(dst, src []byte) {

	err := p.checkParams(dst, src)
	if err != nil {
		panic(err)
	}

	p.we.getBlock(src, p.rb)
	p.fn.Encrypt(p.rb)
	p.we.putBlock(dst, p.rb)
}

func (p *blockCipher) Decrypt(dst, src []byte) {

	err := p.checkParams(dst, src)
	if err != nil {
		panic(err)
	}

	p.we.getBlock(src, p.rb)
	p.fn.Decrypt(p.rb)
	p.we.putBlock(dst, p.rb)
}
