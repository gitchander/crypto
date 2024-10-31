package enigma

import (
	"github.com/gitchander/crypto/enigma/base26"
)

// Cipher

type BytesCrypt struct {
	tf TextFeeder
}

func NewBytesCrypt(e *Enigma) *BytesCrypt {
	return &BytesCrypt{
		tf: TrueTextFeeder(e),
	}
}

func (p *BytesCrypt) Encrypt(plainData []byte) (cipherData []byte, err error) {

	plaintext := base26.EncodeToString(plainData)

	ciphertext, err := p.tf.FeedText(plaintext)
	if err != nil {
		return nil, err
	}

	cipherData = []byte(ciphertext)

	return cipherData, nil
}

func (p *BytesCrypt) Decrypt(cipherData []byte) (plainData []byte, err error) {

	ciphertext := string(cipherData)

	plaintext, err := p.tf.FeedText(ciphertext)
	if err != nil {
		return nil, err
	}

	plainData, err = base26.DecodeString(plaintext)
	if err != nil {
		return nil, err
	}

	return plainData, nil
}
