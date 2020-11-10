package main

const blockSize = 2 * sizeOfWord

type roundBlock struct {
	L, R Word
}

func (p *roundBlock) Swap() {
	p.L, p.R = p.R, p.L
}

func writeBlock(p *roundBlock, we *wordEncoder, data []byte) {
	p.R = we.getWord(data[0*sizeOfWord:])
	p.L = we.getWord(data[1*sizeOfWord:])
}

func readBlock(p *roundBlock, we *wordEncoder, data []byte) {
	we.putWord(data[0*sizeOfWord:], p.R)
	we.putWord(data[1*sizeOfWord:], p.L)
}

func round(b *roundBlock, k Word, f RoundFunc) {
	b.L = xor(b.L, f(k, b.R))
	b.Swap()
}

func encrypt(b *roundBlock, ks []Word, f RoundFunc) {
	n := len(ks)
	for i := 0; i < n; i++ {
		round(b, ks[i], f)
	}
	b.Swap()
}

func decrypt(b *roundBlock, ks []Word, f RoundFunc) {
	n := len(ks)
	for i := n; i > 0; i-- {
		round(b, ks[i-1], f)
	}
	b.Swap()
}
