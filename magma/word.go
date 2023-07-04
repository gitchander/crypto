package magma

import (
	"encoding/binary"

	"github.com/gitchander/crypto/feistel"
)

const bytesPerWord = 4

type word uint32

func wordXOR(a, b word) word {
	return a ^ b
}

func wordShift11(w word) word {
	return (w << 11) | (w >> 21)
}

//------------------------------------------------------------------------------

func cloneWords(a []word) []word {
	b := make([]word, len(a))
	copy(b, a)
	return b
}

func reverseWords(a []word) {
	i, j := 0, (len(a) - 1)
	for i < j {
		a[i], a[j] = a[j], a[i]
		i, j = i+1, j-1
	}
}

//------------------------------------------------------------------------------

// byteOrder
type wordEncoder struct {
	byteOrder binary.ByteOrder
}

func newWordEncoder(byteOrder binary.ByteOrder) *wordEncoder {
	return &wordEncoder{
		byteOrder: byteOrder,
	}
}

var defaultWordEncoder = newWordEncoder(binary.LittleEndian)

func (we *wordEncoder) getWord(b []byte) word {
	return word(we.byteOrder.Uint32(b))
}

func (we *wordEncoder) putWord(b []byte, w word) {
	we.byteOrder.PutUint32(b, uint32(w))
}

func (we *wordEncoder) getBlock(data []byte, p *feistel.RoundBlock[word]) {
	p.R = we.getWord(data[0*bytesPerWord:])
	p.L = we.getWord(data[1*bytesPerWord:])
}

func (we *wordEncoder) putBlock(data []byte, p *feistel.RoundBlock[word]) {
	we.putWord(data[0*bytesPerWord:], p.R)
	we.putWord(data[1*bytesPerWord:], p.L)
}
