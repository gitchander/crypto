package base16

const (
	bitsPerByte   = 8
	valuesPerByte = 1 << bitsPerByte
)

const lastBitMask = 1 << (bitsPerByte - 1)

//------------------------------------------------------------------------------

// Exclude: I, O, Q
const encodeTable = "ABCDEFGHJKLMNPRS"

func encodeNibble(nibble byte) byte {
	return encodeTable[nibble]
}

//------------------------------------------------------------------------------

var decodeTable = func() (dt [valuesPerByte]byte) {
	for i, b := range encodeTable {
		dt[b] = byte(lastBitMask | i)
	}
	return dt
}()

func decodeNibble_v1(b byte) (nibble byte, ok bool) {
	x := decodeTable[b]
	ok = (x & lastBitMask) != 0
	if ok {
		nibble = x &^ lastBitMask
	}
	return nibble, ok
}

func decodeNibble_v2(b byte) (nibble byte, ok bool) {
	x := decodeTable[b]
	nibble = x &^ lastBitMask
	ok = (x & lastBitMask) != 0
	return nibble, ok
}

func decodeNibble(b byte) (nibble byte, ok bool) {
	// return decodeNibble_v1(b)
	return decodeNibble_v2(b)
}

func decodeNibbleErr(b byte) (nibble byte, err error) {
	nibble, ok := decodeNibble(b)
	if !ok {
		return 0, InvalidByteError(b)
	}
	return nibble, nil
}

//------------------------------------------------------------------------------

func EncodedLen(n int) int { return n * 2 }

func DecodedLen(x int) int { return x / 2 }

func Encode(dst, src []byte) int {
	j := 0
	for _, b := range src {
		hi, lo := byteToNibbles(b)

		// dst[j+0] = encodeTable[lo]
		// dst[j+1] = encodeTable[hi]

		dst[j+0] = encodeNibble(lo)
		dst[j+1] = encodeNibble(hi)

		j += 2
	}
	return len(src) * 2
}

func decode_v1(dst, src []byte) (int, error) {
	i, j := 0, 1
	for ; j < len(src); j += 2 {
		lo, ok := decodeNibble(src[j-1])
		if !ok {
			return i, InvalidByteError(src[j-1])
		}
		hi, ok := decodeNibble(src[j])
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

func decode_v2(dst, src []byte) (int, error) {
	i, j := 0, 0
	for ; (j + 1) < len(src); j += 2 {
		lo, err := decodeNibbleErr(src[j+0])
		if err != nil {
			return i, err
		}
		hi, err := decodeNibbleErr(src[j+1])
		if err != nil {
			return i, err
		}
		dst[i] = nibblesToByte(hi, lo)
		i++
	}
	if (len(src) % 2) == 1 {
		// Check for invalid char before reporting bad length,
		// since the invalid char (if present) is an earlier problem.

		if _, err := decodeNibbleErr(src[j+0]); err != nil {
			return i, err
		}
		return i, ErrLength
	}
	return i, nil
}

func Decode(dst, src []byte) (int, error) {
	// return decode_v1(dst, src)
	return decode_v2(dst, src)
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
