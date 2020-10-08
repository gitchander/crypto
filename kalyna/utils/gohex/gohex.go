package gohex

import (
	// "encoding/hex"
	// "fmt"
	"strings"
)

// func printState(prefix string, state []uint64) {
// 	v := wordsToString(state)
// 	fmt.Printf("%s %q\n", prefix, v)
// }

// func wordsToString(ws []uint64) string {
// 	bs := makeBytesForWords(ws)
// 	wordsToBytes(ws, bs)
// 	var br strings.Builder
// 	for _, b := range bs {
// 		br.WriteString(byteToString(b))
// 	}
// 	return br.String()
// }

func writeHexByte(br *strings.Builder, b byte) {

	lo, hi := byteToNibbles(b)

	const upper = true

	var bs [2]byte

	bs[0], _ = nibbleToHex(hi, upper)
	bs[1], _ = nibbleToHex(lo, upper)

	br.WriteByte(bs[0])
	br.WriteByte(bs[1])
}

func byteToString(b byte) string {

	lo, hi := byteToNibbles(b)

	const upper = true

	var bs [2]byte

	bs[0], _ = nibbleToHex(hi, upper)
	bs[1], _ = nibbleToHex(lo, upper)

	return string(bs[:])
}

func byteToNibbles(b byte) (lo, hi byte) {
	lo = b & 0xF
	hi = b >> 4
	return
}

func nibblesToByte(lo, hi byte) byte {
	return (hi << 4) | (lo & 0xF)
}

func nibbleToHex(n byte, upper bool) (byte, bool) {
	if (0 <= n) && (n < 10) {
		return (n + '0'), true
	}
	if upper {
		if (10 <= n) && (n < 16) {
			return (n + 'A' - 10), true
		}
	} else {
		if (10 <= n) && (n < 16) {
			return (n + 'a' - 10), true
		}
	}
	return 0, false
}
