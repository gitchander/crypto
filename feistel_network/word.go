package main

import (
	"encoding/binary"
)

//------------------------------------------------------------------------------
// const sizeOfWord = 4

// type Word uint32

// func xor(a, b Word) Word {
// 	return a ^ b
// }

// // byteOrder
// type wordEncoder struct {
// 	order binary.ByteOrder
// }

// func (we *wordEncoder) getWord(b []byte) Word {
// 	return Word(we.order.Uint32(b))
// }

// func (we *wordEncoder) putWord(b []byte, w Word) {
// 	we.order.PutUint32(b, uint32(w))
// }

// var (
// 	bigEndian    = &wordEncoder{order: binary.BigEndian}
// 	littleEndian = &wordEncoder{order: binary.LittleEndian}
// )

//------------------------------------------------------------------------------
const sizeOfWord = 8

type Word uint64

func xor(a, b Word) Word {
	return a ^ b
}

// byteOrder
type wordEncoder struct {
	order binary.ByteOrder
}

func (we *wordEncoder) getWord(b []byte) Word {
	return Word(we.order.Uint64(b))
}

func (we *wordEncoder) putWord(b []byte, w Word) {
	we.order.PutUint64(b, uint64(w))
}

var (
	bigEndian    = &wordEncoder{order: binary.BigEndian}
	littleEndian = &wordEncoder{order: binary.LittleEndian}
)

//------------------------------------------------------------------------------
