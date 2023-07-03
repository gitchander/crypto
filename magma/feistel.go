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

func round(b *roundBlock, k word, rf roundFunc) {
	b.L = wordXOR(b.L, rf(k, b.R))
	b.Swap()
}

func encrypt(b *roundBlock, ks []word, rf roundFunc) {
	for _, k := range ks {
		round(b, k, rf)
	}
	b.Swap()
}

func decrypt(b *roundBlock, ks []word, rf roundFunc) {
	for _, k := range ks {
		round(b, k, rf)
	}
	b.Swap()
}
