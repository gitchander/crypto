package kalyna

import (
	"fmt"
	"strconv"
)

type BlockSizeError int

func (k BlockSizeError) Error() string {
	return "kalyna: invalid block size " + strconv.Itoa(int(k))
}

type KeySizeError int

func (k KeySizeError) Error() string {
	return "kalyna: invalid key size " + strconv.Itoa(int(k))
}

type kalynaContext struct {
	nb     int // Number of 64-bit words in enciphering block.
	nk     int // Number of 64-bit words in key.
	rounds int // Number of enciphering rounds.

	state     []uint64   // Current cipher state.
	roundKeys [][]uint64 // Round key computed from enciphering key.
}

// NewKalyna:
// Initialize Kalyna parameters and create cipher context.
// params:
// @param block_size Enciphering block bit size (128, 256 or 512 bit sizes are
// allowed).
// @param block_size Enciphering key bit size. Must be equal or double the
// block bit size.
// @return Pointer to Kalyna context containing cipher instance
// parameters and allocated memory for state and round keys. NULL in case of
// error.
func newKalynaContext(blockSize, keySize int) (*kalynaContext, error) {

	var nb, nk, rounds int

	if blockSize == bitsPerBlock128 {
		nb = wordsPerBlock128
		if keySize == bitsPerKey128 {
			nk = wordsPerKey128
			rounds = roundsForKey128
		} else if keySize == bitsPerKey256 {
			nk = wordsPerKey256
			rounds = roundsForKey256
		} else {
			return nil, KeySizeError(keySize)
		}
	} else if blockSize == bitsPerBlock256 {
		nb = wordsPerBlock256
		if keySize == bitsPerKey256 {
			nk = wordsPerKey256
			rounds = roundsForKey256
		} else if keySize == bitsPerKey512 {
			nk = wordsPerKey512
			rounds = roundsForKey512
		} else {
			return nil, KeySizeError(keySize)
		}
	} else if blockSize == bitsPerBlock512 {
		nb = wordsPerBlock512
		if keySize == bitsPerKey512 {
			nk = wordsPerKey512
			rounds = roundsForKey512
		} else {
			return nil, KeySizeError(keySize)
		}
	} else {
		return nil, BlockSizeError(blockSize)
	}

	state := make([]uint64, nb)

	roundKeys := make([][]uint64, rounds+1)
	for i := range roundKeys {
		roundKeys[i] = make([]uint64, nb)
	}

	k := &kalynaContext{
		nb:     nb,
		nk:     nk,
		rounds: rounds,

		state:     state,
		roundKeys: roundKeys,
	}

	return k, nil
}

// KeyExpand:
// Compute round keys given the enciphering key and store them
// in cipher context `k`.
// params:
// param key - Kalyna enciphering key.
func (k *kalynaContext) KeyExpand(key []uint64) error {

	if len(key) != k.nk {
		return fmt.Errorf("kalyna: invalid key size: have %d, want %d",
			len(key), k.nk)
	}

	kt := poolGetWords(k.nb)

	k.keyExpandKt(key, kt)
	k.keyExpandEven(key, kt)
	k.keyExpandOdd()

	poolPutWords(kt)

	return nil
}

// Encipher:
// Encipher plaintext using Kalyna symmetric block cipher.
// KalynaInit() function with appropriate block and enciphering key sizes must
// be called beforehand to get the cipher context `k`. After all enciphering
// is completed KalynaDelete() must be called to free up allocated memory.
// params:
// @param plaintext Plaintext of length Nb words for enciphering.
// @param ciphertext The result of enciphering.
func (k *kalynaContext) Encipher(plaintext, ciphertext []uint64) {

	copy(k.state, plaintext)

	k.addRoundKey(0)
	for round := 1; round < k.rounds; round++ {
		k.encipherRound()
		k.xorRoundKey(round)
	}
	k.encipherRound()
	k.addRoundKey(k.rounds)

	copy(ciphertext, k.state)
}

// Decipher:
// Decipher ciphertext using Kalyna symmetric block cipher.
// KalynaInit() function with appropriate block and enciphering key sizes must
// be called beforehand to get the cipher context `k`. After all enciphering
// is completed KalynaDelete() must be called to free up allocated memory.
// params:
// @param ciphertext Enciphered data of length Nb words.
// @param plaintext The result of deciphering.
func (k *kalynaContext) Decipher(ciphertext, plaintext []uint64) {

	copy(k.state, ciphertext)

	k.subRoundKey(k.rounds)
	for round := k.rounds - 1; round > 0; round-- {
		k.decipherRound()
		k.xorRoundKey(round)
	}
	k.decipherRound()
	k.subRoundKey(0)

	copy(plaintext, k.state)
}

// encipherRound:
// Perform single round enciphering routine.
// params:
// @param current state and round keys precomputed.
func (k *kalynaContext) encipherRound() {
	subBytes(k.state)
	shiftRows(k.state)
	mixColumns(k.state)
}

// decipherRound:
// Perform single round deciphering routine.
// params:
// @param current state and round keys precomputed.
func (k *kalynaContext) decipherRound() {
	invMixColumns(k.state)
	invShiftRows(k.state)
	invSubBytes(k.state)
}

// addRoundKey:
// Inject round key into the state using addition modulo 2^{64}.
// params:
// @param round Number of the round on which the key addition is performed in
// order to use the correct round key.
func (k *kalynaContext) addRoundKey(round int) {
	for i := range k.state {
		k.state[i] += k.roundKeys[round][i]
	}
}

// subRoundKey:
// Extract round key from the state using subtraction modulo 2^{64}.
// params:
// @param round Number of the round on which the key subtraction is performed
// in order to use the correct round key.
func (k *kalynaContext) subRoundKey(round int) {
	for i := range k.state {
		k.state[i] -= k.roundKeys[round][i]
	}
}

// addRoundKeyExpand:
// Perform addition of two arbitrary states modulo 2^{64}.
// The operation is identical to simple round key addition but on arbitrary
// state array and addition value (instead of the actual round key). Used in
// key expansion procedure. The result is stored in `state`.
// params:
// @param value Is to be added to the state array modulo 2^{64}.
func (k *kalynaContext) addRoundKeyExpand(value []uint64) {
	for i := range k.state {
		k.state[i] += value[i]
	}
}

// xorRoundKey:
// Inject round key into the state using XOR operation.
// params:
// @param round Number of the round on which the key addition is performed in
// order to use the correct round key.
func (k *kalynaContext) xorRoundKey(round int) {
	for i := range k.state {
		k.state[i] ^= k.roundKeys[round][i]
	}
}

// xorRoundKeyExpand:
// Perform XOR of two arbitrary states.
// The operation is identical to simple round key XORing but on arbitrary
// state array and addition value (instead of the actual round key). Used in
// key expansion procedure. The result is stored in `state`.
// XOR operation is involutive so no inverse transformation is required.
// params:
// @param value Is to be added to the state array modulo 2^{64}.
func (k *kalynaContext) xorRoundKeyExpand(value []uint64) {
	for i := range k.state {
		k.state[i] ^= value[i]
	}
}

// keyExpandKt:
// Generate the Kt value (auxiliary key used in key expansion).
// params:
// @param key Enciphering key of size corresponding to the one stored in cipher
// context `k` (specified via KalynaInit() function).
// @param kt Array for storing generated Kt value.
func (k *kalynaContext) keyExpandKt(key []uint64, kt []uint64) {

	var (
		k0 = poolGetWords(k.nb)
		k1 = poolGetWords(k.nb)
	)

	wordsSet(k.state, 0)
	k.state[0] += uint64(k.nb + k.nk + 1)

	if k.nb == k.nk {
		copy(k0, key)
		copy(k1, key)
	} else {
		// For: (k.nb * 2) == k.nk
		copy(k0, key[:k.nb])
		copy(k1, key[k.nb:])
	}

	k.addRoundKeyExpand(k0)
	k.encipherRound()
	k.xorRoundKeyExpand(k1)
	k.encipherRound()
	k.addRoundKeyExpand(k0)
	k.encipherRound()

	copy(kt, k.state)

	poolPutWords(k0)
	poolPutWords(k1)
}

// keyExpandEven:
// Compute even round keys and store them in cipher context `k`.
// params:
// @param key Kalyna enciphering key of length Nk 64-bit words.
// @param kt Kalyna auxiliary key. The size is equal to enciphering state
// size and equals Nb 64-bit words.
func (k *kalynaContext) keyExpandEven(key []uint64, kt []uint64) {

	var (
		initialData = poolGetWords(k.nk)
		roundKT     = poolGetWords(k.nb)
		tmv         = poolGetWords(k.nb)
	)

	copy(initialData, key)

	for i := range tmv {
		tmv[i] = 0x0001000100010001
	}

	round := 0
	for {
		copy(k.state, kt)

		k.addRoundKeyExpand(tmv)

		copy(roundKT, k.state)
		copy(k.state, initialData)

		k.addRoundKeyExpand(roundKT)
		k.encipherRound()
		k.xorRoundKeyExpand(roundKT)
		k.encipherRound()
		k.addRoundKeyExpand(roundKT)

		copy(k.roundKeys[round], k.state)

		if k.rounds == round {
			break
		}

		if k.nb != k.nk {
			// For: (k.nb * 2) == k.nk
			if (k.nb * 2) != k.nk {
				panic("invalid nb, kn size")
			}

			round += 2

			shiftLeft(tmv)

			copy(k.state, kt)

			k.addRoundKeyExpand(tmv)

			copy(roundKT, k.state)
			copy(k.state, initialData[k.nb:])

			k.addRoundKeyExpand(roundKT)
			k.encipherRound()
			k.xorRoundKeyExpand(roundKT)
			k.encipherRound()
			k.addRoundKeyExpand(roundKT)

			copy(k.roundKeys[round], k.state)

			if k.rounds == round {
				break
			}
		}
		round += 2
		shiftLeft(tmv)
		rotate(initialData)
	}

	poolPutWords(initialData)
	poolPutWords(roundKT)
	poolPutWords(tmv)
}

// keyExpandOdd:
// Compute odd round keys by rotating already generated even ones and
// fill in the rest of the round keys in cipher context `k`.
func (k *kalynaContext) keyExpandOdd() {
	for i := 1; i < k.rounds; i += 2 {
		copy(k.roundKeys[i], k.roundKeys[i-1])
		rotateLeft(k.roundKeys[i])
	}
}
