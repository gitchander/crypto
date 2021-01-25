package magma

type roundBlock struct {
	L, R word
}

func (p *roundBlock) Swap() {
	p.L, p.R = p.R, p.L
}

func writeBlock(p *roundBlock, we *wordEncoder, data []byte) {
	p.R = we.getWord(data[0*bytesPerWord:])
	p.L = we.getWord(data[1*bytesPerWord:])
}

func readBlock(p *roundBlock, we *wordEncoder, data []byte) {
	we.putWord(data[0*bytesPerWord:], p.R)
	we.putWord(data[1*bytesPerWord:], p.L)
}

// round function
type roundFunc func(k, r word) word

func round(b *roundBlock, k word, rf roundFunc) {
	b.L = xor(b.L, rf(k, b.R))
	b.Swap()
}

func encrypt(b *roundBlock, ks []word, rf roundFunc) {
	n := len(ks)
	for i := 0; i < n; i++ {
		round(b, ks[i], rf)
	}
	b.Swap()
}

func decrypt(b *roundBlock, ks []word, rf roundFunc) {
	n := len(ks)
	for i := n; i > 0; i-- {
		round(b, ks[i-1], rf)
	}
	b.Swap()
}
