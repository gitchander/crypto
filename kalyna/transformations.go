package kalyna

import (
	"encoding/binary"
)

// Words
func bytesPerWords(ws []uint64) int {
	return len(ws) * bytesPerWord
}

func wordsSet(as []uint64, a uint64) {
	for i := range as {
		as[i] = a
	}
}

func wordsToBytes(ws []uint64, bs []byte) {
	for _, w := range ws {
		binary.LittleEndian.PutUint64(bs, w)
		bs = bs[bytesPerWord:]
	}
}

func bytesToWords(ws []uint64, bs []byte) {
	for i := range ws {
		ws[i] = binary.LittleEndian.Uint64(bs)
		bs = bs[bytesPerWord:]
	}
}

//------------------------------------------------------------------------------
// func sboxesTransform(ws []uint64, sboxes *[4][256]uint8) {
// 	for i, w := range ws {
// 		var v uint64
// 		v |= uint64(sboxes[0][(w&0x00000000_000000FF)>>0]) << 0
// 		v |= uint64(sboxes[1][(w&0x00000000_0000FF00)>>8]) << 8
// 		v |= uint64(sboxes[2][(w&0x00000000_00FF0000)>>16]) << 16
// 		v |= uint64(sboxes[3][(w&0x00000000_FF000000)>>24]) << 24
// 		v |= uint64(sboxes[0][(w&0x000000FF_00000000)>>32]) << 32
// 		v |= uint64(sboxes[1][(w&0x0000FF00_00000000)>>40]) << 40
// 		v |= uint64(sboxes[2][(w&0x00FF0000_00000000)>>48]) << 48
// 		v |= uint64(sboxes[3][(w&0xFF000000_00000000)>>56]) << 56
// 		ws[i] = v
// 	}
// }

func sboxesTransform(ws []uint64, sboxes *[4][256]uint8) {
	nsbox := len(sboxes)
	for i, w := range ws {
		var v uint64
		for j := 0; j < bitsPerByte; j++ {
			v |= uint64(sboxes[j%nsbox][w&0xff]) << (j * bitsPerByte)
			w >>= bitsPerByte
		}
		ws[i] = v
	}
}

//------------------------------------------------------------------------------
// subBytes substitutes each byte of the cipher state using corresponding S-Boxes.
func subBytes(ws []uint64) {
	sboxesTransform(ws, &sboxesEnc)
}

// invSubBytes inverses subBytes transformation.
func invSubBytes(ws []uint64) {
	sboxesTransform(ws, &sboxesDec)
}

// shiftLeft:
// Shift each word one bit to the left.
// The shift of each word is independent of other array words.
// params:
// @param state_value State represented as 64-bit words array.
// Note that this state Nk words long and differs from
// the cipher state used during enciphering.
func shiftLeft(ws []uint64) {
	for i := range ws {
		ws[i] <<= 1
	}
}

// rotate:
// Rotate words of a state.
// The state is processed as 64-bit words array {w_{0}, w_{1}, ..., w_{nk-1}}
// and rotation is performed so the resulting state is
// {w_{1}, ..., w_{nk-1}, w_{0}}.
// params:
// @param state_value A state represented by 64-bit words array of length Nk.
// It is not the cipher state that is used during enciphering.
func rotate(ws []uint64) {
	n := len(ws)
	if n == 0 {
		return
	}
	temp := ws[0]
	for i := 1; i < n; i++ {
		ws[i-1] = ws[i]
	}
	ws[n-1] = temp
}

// rotateLeft:
// Rotate the state (2 * Nb + 3) bytes to the left.
// The state is interpreted as bytes string in little endian. Big endian
// architectures are also correctly processed by this function.
// params:
// @param state_value A state represented by 64-bit words array of length Nk.
// It is not the cipher state that is used during enciphering.
func rotateLeft(state []uint64) {

	rotateSize := 2*len(state) + 3

	stateBytes := poolGetBytes(bytesPerWords(state))
	wordsToBytes(state, stateBytes)

	// Rotate bytes in memory.
	rotateBytes(stateBytes, rotateSize)

	bytesToWords(state, stateBytes)
	poolPutBytes(stateBytes)
}

func rotateBytes(bs []byte, rotateSize int) {

	rbs := poolGetBytes(rotateSize)

	copy(rbs, bs)
	copy(bs, bs[rotateSize:])
	copy(bs[len(bs)-rotateSize:], rbs)

	poolPutBytes(rbs)
}

// multiplyGF multiply bytes in Finite Field GF(2^8).
// Parameter x is multiplicand element of GF(2^8).
// Parameter y is multiplier element of GF(2^8) from MDS matrix.
// It returns product of multiplication in GF(2^8).
func multiplyGF(x, y uint8) uint8 {
	var (
		r    uint8 = 0
		hbit uint8 = 0
	)
	for i := 0; i < bitsPerByte; i++ {
		if (y & 0x1) == 1 {
			r ^= x
		}
		hbit = x & 0x80
		x <<= 1
		if hbit == 0x80 {
			x ^= reductionPolynomialUint8
		}
		y >>= 1
	}
	return r
}

func indexByRowCol(row, col int) int {
	return row + (col * bytesPerWord)
}

// matrixMultiply:
// Multiply cipher state by specified MDS matrix.
// Used to avoid code repetition for MixColumn and Inverse MixColumn.
// params:
// @param current state and round keys precomputed.
// @param matrix MDS 8x8 byte matrix.
func matrixMultiply(state []uint64, matrix *[8][8]uint8) {

	stateLen := len(state)

	stateBytes := poolGetBytes(bytesPerWords(state))
	wordsToBytes(state, stateBytes)

	for col := 0; col < stateLen; col++ {
		var result uint64
		for row := bytesPerWord - 1; row >= 0; row-- {
			var product uint8
			for b := bytesPerWord - 1; b >= 0; b-- {
				v := stateBytes[indexByRowCol(b, col)]
				product ^= multiplyGF(v, matrix[row][b])
			}
			result |= uint64(product) << (row * bytesPerWord)
		}
		state[col] = result
	}
	poolPutBytes(stateBytes)
}

// mixColumns:
// Perform MixColumn transformation to the cipher state.
// params:
// @param current state and round keys precomputed.
func mixColumns(state []uint64) {
	matrixMultiply(state, &mdsMatrix)
}

// invMixColumns:
// Inverse MixColumn transformation.
// params:
// @param current state and round keys precomputed.
func invMixColumns(state []uint64) {
	matrixMultiply(state, &mdsInvMatrix)
}

// shiftRows:
// Shift cipher state rows according to specification.
// params:
// @param current state and round keys precomputed.
func shiftRows(state []uint64) {

	stateLen := len(state)

	var (
		stateBytes0 = poolGetBytes(bytesPerWords(state))
		stateBytes1 = poolGetBytes(bytesPerWords(state))
	)

	wordsToBytes(state, stateBytes0)

	shift := -1
	for row := 0; row < bytesPerWord; row++ {

		if (row % (bytesPerWord / stateLen)) == 0 {
			shift += 1
		}

		for col := 0; col < stateLen; col++ {
			val := stateBytes0[indexByRowCol(row, col)]
			stateBytes1[indexByRowCol(row, (col+shift)%stateLen)] = val
		}
	}

	bytesToWords(state, stateBytes1)

	poolPutBytes(stateBytes0)
	poolPutBytes(stateBytes1)
}

// invShiftRows:
// Inverse ShiftRows transformation.
// params:
// @param current state and round keys precomputed.
func invShiftRows(state []uint64) {

	stateLen := len(state)

	var (
		stateBytes0 = poolGetBytes(bytesPerWords(state))
		stateBytes1 = poolGetBytes(bytesPerWords(state))
	)

	wordsToBytes(state, stateBytes0)

	shift := -1

	for row := 0; row < bytesPerWord; row++ {

		if (row % (bytesPerWord / stateLen)) == 0 {
			shift += 1
		}

		for col := 0; col < stateLen; col++ {
			val := stateBytes0[indexByRowCol(row, (col+shift)%stateLen)]
			stateBytes1[indexByRowCol(row, col)] = val
		}
	}

	bytesToWords(state, stateBytes1)

	poolPutBytes(stateBytes0)
	poolPutBytes(stateBytes1)
}

// reverseWord:
// Reverse bytes ordering that form the word.
// params:
// @param word 64-bit word that needs its bytes to be reversed (perhaps for
// converting between little and big endian).
// @return 64-bit word with reversed bytes.
func reverseWord(x uint64) uint64 {
	var r uint64
	for i := 0; i < bytesPerWord; i++ {
		r = (r << 8) | (x & 0xFF)
		x >>= 8
	}
	return r
}
