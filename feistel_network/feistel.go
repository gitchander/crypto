package main

import (
	"golang.org/x/exp/constraints"
)

type Unsigned = constraints.Unsigned

func cloneSlice[T any](a []T) []T {
	b := make([]T, len(a))
	copy(b, a)
	return b
}

func reverse[T any](a []T) {
	i, j := 0, (len(a) - 1)
	for i < j {
		a[i], a[j] = a[j], a[i]
		i, j = i+1, j-1
	}
}

//------------------------------------------------------------------------------

type RoundBlock[T Unsigned] struct {
	L, R T
}

func (p *RoundBlock[T]) Swap() {
	p.L, p.R = p.R, p.L
}

// round function
type RoundFunc[T Unsigned] func(k, r T) T

type FeistelNetwork[T Unsigned] struct {
	rf      RoundFunc[T]
	encKeys []T
	decKeys []T
}

func NewFeistelNetwork[T Unsigned](ks []T, rf RoundFunc[T]) *FeistelNetwork[T] {

	var (
		encKeys = cloneSlice(ks)
		decKeys = cloneSlice(ks)
	)
	reverse(decKeys)

	return &FeistelNetwork[T]{
		rf:      rf,
		encKeys: encKeys,
		decKeys: decKeys,
	}
}

func (p *FeistelNetwork[T]) round(k T, b *RoundBlock[T]) {
	b.L = b.L ^ p.rf(k, b.R) // XOR(L, F(K, R))
	b.Swap()
}

func (p *FeistelNetwork[T]) Encrypt(b *RoundBlock[T]) {
	ks := p.encKeys
	for _, k := range ks {
		p.round(k, b)
	}
	b.Swap()
}

func (p *FeistelNetwork[T]) Decrypt(b *RoundBlock[T]) {
	ks := p.decKeys
	for _, k := range ks {
		p.round(k, b)
	}
	b.Swap()
}
