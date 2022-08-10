package base16

import (
	"encoding/hex"
)

const (
	bitsPerByte   = 8
	valuesPerByte = 1 << bitsPerByte
)

func testHex() {
	text1 := "Hello, World!"
	src := []byte(text1)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	res := make([]byte, hex.DecodedLen(len(dst)))
	hex.Decode(res, dst)
	text2 := string(res)
	_ = text2
}

func EncodedLen(n int) int { return n * 2 }

func DecodedLen(x int) int { return x / 2 }

// Exclude: I, O, Q
const encodeTable = "ABCDEFGHJKLMNPRS"

const lastBitMask = 1 << (bitsPerByte - 1)

var decodeTable = func() (dt [valuesPerByte]byte) {
	for i, b := range encodeTable {
		dt[b] = byte(lastBitMask | i)
	}
	return dt
}()

func decodeNibble(b byte) (byte, bool) {
	x := decodeTable[b]
	if (x & lastBitMask) != 0 {
		return x &^ lastBitMask, true
	}
	return 0, false
}

func Encode(dst, src []byte) int {
	j := 0
	for _, b := range src {
		hi, lo := byteToNibbles(b)

		dst[j+0] = encodeTable[hi]
		dst[j+1] = encodeTable[lo]

		j += 2
	}
	return len(src) * 2
}

func Decode(dst, src []byte) (int, error) {
	i, j := 0, 1
	for ; j < len(src); j += 2 {
		hi, ok := decodeNibble(src[j-1])
		if !ok {
			return i, InvalidByteError(src[j-1])
		}
		lo, ok := decodeNibble(src[j])
		if !ok {
			return i, InvalidByteError(src[j])
		}
		dst[i] = nibblesToByte(hi, lo)
		i++
	}
	if (len(src) % 2) == 1 {
		// Check for invalid char before reporting bad length,
		// since the invalid char (if present) is an earlier problem.
		if _, ok := decodeNibble(src[j-1]); !ok {
			return i, InvalidByteError(src[j-1])
		}
		return i, ErrLength
	}
	return i, nil
}

func EncodeToString(src []byte) string {
	dst := make([]byte, EncodedLen(len(src)))
	Encode(dst, src)
	return string(dst)
}

func DecodeString(s string) ([]byte, error) {
	src := []byte(s)
	// We can use the source slice itself as the destination
	// because the decode loop increments by one and then the 'seen' byte is not used anymore.
	n, err := Decode(src, src)
	return src[:n], err
}
