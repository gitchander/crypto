package base26

import (
	"fmt"
)

const bitsPerByte = 8

const encodeTable = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	mask4bit = 0x0f // 0000_1111
	mask5bit = 0x1f // 0001_1111
)

func decodeChar(b byte) (digit int, ok bool) {
	if ('A' <= b) && (b <= 'Z') {
		return int(b - 'A'), true
	}
	return 0, false
}

func Encode(dst, src []byte) int {

	j := 0

	var (
		ba uint32 // bits accumulator
		bn int    // bits length
	)

	for _, b := range src {

		ba |= uint32(b) << bn
		bn += bitsPerByte

		for bn >= 5 {

			v := int(ba & mask4bit) // 4 bits

			if v > mask4bitValueMax {
				ba >>= 4
				bn -= 4
			} else {
				v = int(ba & mask5bit) // 5 bits
				ba >>= 5
				bn -= 5
			}

			dst[j] = encodeTable[v]
			j++
		}
	}

	if bn > 0 {
		v := int(ba) // ba has 4 or less bits

		dst[j] = encodeTable[v]
		j++
	}

	return j
}

func Decode(dst, src []byte) (int, error) {

	j := 0
	var (
		ba uint32 // bits accumulator
		bn int    // bits length
	)

	for _, b := range src {

		x, ok := decodeChar(b)
		if !ok {
			return j, fmt.Errorf("base26: invalid byte: %#U", rune(b))
		}

		v := int(x & mask4bit)

		if v > mask4bitValueMax {
			ba |= uint32(v) << bn
			bn += 4
		} else {
			v = int(x & mask5bit)
			ba |= uint32(v) << bn
			bn += 5
		}

		for bn >= bitsPerByte {
			dst[j] = byte(ba)
			j++

			ba >>= bitsPerByte
			bn -= bitsPerByte
		}
	}

	if (bn > 4) || (ba != 0) {
		return j, fmt.Errorf("base26: invalid source (bits: length %d, accumulator %b)", bn, ba)
	}

	return j, nil
}
