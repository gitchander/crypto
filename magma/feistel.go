package magma

// Feistel cipher:
// https://en.wikipedia.org/wiki/Feistel_cipher

type roundBlock struct {
	L, R word
}

func (p *roundBlock) Swap() {
	p.L, p.R = p.R, p.L
}

// round function
type roundFunc func(k, r word) word

type feistelNetwork struct {
	rf      roundFunc
	encKeys []word
	decKeys []word
}

func newFeistelNetwork(ks []word, rf roundFunc) *feistelNetwork {

	var (
		encKeys = cloneWords(ks)
		decKeys = cloneWords(ks)
	)
	reverseWords(decKeys)

	return &feistelNetwork{
		rf:      rf,
		encKeys: encKeys,
		decKeys: decKeys,
	}
}

func (p *feistelNetwork) round(k word, b *roundBlock) {
	b.L = b.L ^ p.rf(k, b.R) // XOR(L, F(K, R))
	b.Swap()
}

func (p *feistelNetwork) encrypt(b *roundBlock) {
	ks := p.encKeys
	for _, k := range ks {
		p.round(k, b)
	}
	b.Swap()
}

func (p *feistelNetwork) decrypt(b *roundBlock) {
	ks := p.decKeys
	for _, k := range ks {
		p.round(k, b)
	}
	b.Swap()
}
