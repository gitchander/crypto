package feistel

import (
	"golang.org/x/exp/constraints"
)

// Feistel cipher:
// https://en.wikipedia.org/wiki/Feistel_cipher

type Unsigned = constraints.Unsigned

type RoundBlock[T Unsigned] struct {
	L, R T
}

func (p *RoundBlock[T]) Swap() {
	p.L, p.R = p.R, p.L
}

type RoundFunc[T Unsigned] func(k, r T) T

// Network, Cipher
type Cipher[T Unsigned] struct {
	rf      RoundFunc[T]
	encKeys []T
	decKeys []T
}

func NewCipher[T Unsigned](ks []T, rf RoundFunc[T]) *Cipher[T] {

	var (
		encKeys = cloneSlice(ks)
		decKeys = cloneSlice(ks)
	)
	reverseSlice(decKeys)

	return &Cipher[T]{
		rf:      rf,
		encKeys: encKeys,
		decKeys: decKeys,
	}
}

func (p *Cipher[T]) round(k T, b *RoundBlock[T]) {
	b.L = xor(b.L, p.rf(k, b.R))
	b.Swap()
}

func (p *Cipher[T]) Encrypt(b *RoundBlock[T]) {
	ks := p.encKeys
	for _, k := range ks {
		p.round(k, b)
	}
	b.Swap()
}

func (p *Cipher[T]) Decrypt(b *RoundBlock[T]) {
	ks := p.decKeys
	for _, k := range ks {
		p.round(k, b)
	}
	b.Swap()
}
