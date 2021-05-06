package magma

import (
	"encoding/binary"
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
// byteOrder
type wordEncoder struct {
	byteOrder binary.ByteOrder
}

func (we *wordEncoder) getWord(b []byte) word {
	return word(we.byteOrder.Uint32(b))
}

func (we *wordEncoder) putWord(b []byte, w word) {
	we.byteOrder.PutUint32(b, uint32(w))
}

var (
	byteOrder = binary.LittleEndian
)

func newWordEncoder() *wordEncoder {
	return &wordEncoder{
		byteOrder: byteOrder,
	}
}

// var (
// 	bigEndian    = &wordEncoder{order: binary.BigEndian}
// 	littleEndian = &wordEncoder{order: binary.LittleEndian}
// )
