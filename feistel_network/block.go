package main

import (
	"crypto/cipher"
	"fmt"
)

func NewCipher(key []byte) (cipher.Block, error) {
	return newBlockCipher(key)
}

type blockCipher struct {
	we *wordEncoder
	ks []Word
	rb *roundBlock
	rf RoundFunc
}

func newBlockCipher(key []byte) (cipher.Block, error) {

	we := bigEndian

	ks, err := expandKey(key, we)
	if err != nil {
		return nil, err
	}

	block := &blockCipher{
		we: we,
		ks: ks,
		rb: new(roundBlock),
		rf: RoundFuncXOR,
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

	writeBlock(p.rb, p.we, src)
	encrypt(p.rb, p.ks, p.rf)
	readBlock(p.rb, p.we, dst)
}

func (p *blockCipher) Decrypt(dst, src []byte) {

	err := p.checkParams(dst, src)
	if err != nil {
		panic(err)
	}

	writeBlock(p.rb, p.we, src)
	decrypt(p.rb, p.ks, p.rf)
	readBlock(p.rb, p.we, dst)
}
