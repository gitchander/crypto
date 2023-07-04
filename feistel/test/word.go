package main

import (
	"encoding/binary"

	"github.com/gitchander/crypto/feistel"
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

// ------------------------------------------------------------------------------
const sizeOfWord = 8

const blockSize = 2 * sizeOfWord

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

func (we *wordEncoder) getBlock(data []byte, p *feistel.RoundBlock[Word]) {
	p.R = we.getWord(data[0*sizeOfWord:])
	p.L = we.getWord(data[1*sizeOfWord:])
}

func (we *wordEncoder) putBlock(data []byte, p *feistel.RoundBlock[Word]) {
	we.putWord(data[0*sizeOfWord:], p.R)
	we.putWord(data[1*sizeOfWord:], p.L)
}

var (
	bigEndian    = &wordEncoder{order: binary.BigEndian}
	littleEndian = &wordEncoder{order: binary.LittleEndian}
)

//------------------------------------------------------------------------------
