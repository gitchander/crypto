package gohex

import (
	"strings"
)

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
	hi = (b >> 4)
	lo = (b & 0xF)
	return
}

func nibblesToByte(lo, hi byte) (b byte) {
	b |= (hi << 4)
	b |= (lo & 0xF)
	return b
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
