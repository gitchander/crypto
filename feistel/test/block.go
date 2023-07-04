package main

import (
	"crypto/cipher"
	"fmt"

	"github.com/gitchander/crypto/feistel"
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
	fc *feistel.Cipher[Word]
	rb *feistel.RoundBlock[Word]
}

func newBlockCipher(ks []Word, rf feistel.RoundFunc[Word]) (cipher.Block, error) {
	we := bigEndian
	block := &blockCipher{
		we: we,
		fc: feistel.NewCipher(ks, rf),
		rb: new(feistel.RoundBlock[Word]),
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
	p.fc.Encrypt(p.rb)
	p.we.putBlock(dst, p.rb)
}

func (p *blockCipher) Decrypt(dst, src []byte) {

	err := p.checkParams(dst, src)
	if err != nil {
		panic(err)
	}

	p.we.getBlock(src, p.rb)
	p.fc.Decrypt(p.rb)
	p.we.putBlock(dst, p.rb)
}
