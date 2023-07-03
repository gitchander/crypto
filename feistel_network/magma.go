package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

// https://en.wikipedia.org/wiki/GOST_(block_cipher)

func NewCipherMagma(key []byte) (cipher.Block, error) {
	return NewCipherMagmaSBox(key, SBox1)
}

func NewCipherMagmaSBox(key []byte, sbox SBoxMagma) (cipher.Block, error) {

	we := littleEndian

	ks, err := expandKeyMagma(key, we)
	if err != nil {
		return nil, err
	}

	block := &blockCipher{
		we: we,
		ks: ks,
		rb: new(roundBlock),
		rf: roundFuncMagma(&sbox),
	}

	return block, nil
}

func roundFuncMagma(sbox *SBoxMagma) RoundFunc {
	return func(k, r Word) Word {
		s := k + r
		s = substituteMagma(sbox, s)
		return shiftWord11(s)
	}
}

func shiftWord11(s Word) Word {
	return (s << 11) | (s >> 21)
}

func expandKeyMagma(key []byte, we *wordEncoder) ([]Word, error) {

	xs, err := expandKey(key, we)
	if err != nil {
		return nil, err
	}

	if len(xs) != 8 {
		return nil, ErrorKeyLen
	}

	var ks []Word

	for j := 0; j < 3; j++ {
		for i := 0; i < 8; i++ {
			ks = append(ks, xs[i])
		}
	}
	for i := 8; i > 0; i-- {
		ks = append(ks, xs[i-1])
	}

	return ks, nil
}

// ------------------------------------------------------------------------------
// SBox
// ------------------------------------------------------------------------------
type SBoxMagma [8][16]byte

func substituteMagma(sbox *SBoxMagma, s0 Word) (s1 Word) {
	for i := 0; i < 8; i++ {
		var (
			shift = 4 * i
			j     = ((s0 >> shift) & 0xF)
		)
		s1 |= Word(sbox[i][j]) << shift
	}
	return s1
}

var SBox1 = SBoxMagma{
	{0x4, 0xA, 0x9, 0x2, 0xD, 0x8, 0x0, 0xE, 0x6, 0xB, 0x1, 0xC, 0x7, 0xF, 0x5, 0x3},
	{0xE, 0xB, 0x4, 0xC, 0x6, 0xD, 0xF, 0xA, 0x2, 0x3, 0x8, 0x1, 0x0, 0x7, 0x5, 0x9},
	{0x5, 0x8, 0x1, 0xD, 0xA, 0x3, 0x4, 0x2, 0xE, 0xF, 0xC, 0x7, 0x6, 0x0, 0x9, 0xB},
	{0x7, 0xD, 0xA, 0x1, 0x0, 0x8, 0x9, 0xF, 0xE, 0x4, 0x6, 0xC, 0xB, 0x2, 0x5, 0x3},
	{0x6, 0xC, 0x7, 0x1, 0x5, 0xF, 0xD, 0x8, 0x4, 0xA, 0x9, 0xE, 0x0, 0x3, 0xB, 0x2},
	{0x4, 0xB, 0xA, 0x0, 0x7, 0x2, 0x1, 0xD, 0x3, 0x6, 0x8, 0x5, 0x9, 0xC, 0xF, 0xE},
	{0xD, 0xB, 0x4, 0x1, 0x3, 0xF, 0x5, 0x9, 0x0, 0xA, 0xE, 0x7, 0x6, 0x8, 0x2, 0xC},
	{0x1, 0xF, 0xD, 0x0, 0x5, 0x7, 0xA, 0x4, 0x9, 0x2, 0x3, 0xE, 0x6, 0xB, 0x8, 0xC},
}

var SBox2 = SBoxMagma{
	{0xC, 0x4, 0x6, 0x2, 0xA, 0x5, 0xB, 0x9, 0xE, 0x8, 0xD, 0x7, 0x0, 0x3, 0xF, 0x1},
	{0x6, 0x8, 0x2, 0x3, 0x9, 0xA, 0x5, 0xC, 0x1, 0xE, 0x4, 0x7, 0xB, 0xD, 0x0, 0xF},
	{0xB, 0x3, 0x5, 0x8, 0x2, 0xF, 0xA, 0xD, 0xE, 0x1, 0x7, 0x4, 0xC, 0x9, 0x6, 0x0},
	{0xC, 0x8, 0x2, 0x1, 0xD, 0x4, 0xF, 0x6, 0x7, 0x0, 0xA, 0x5, 0x3, 0xE, 0x9, 0xB},
	{0x7, 0xF, 0x5, 0xA, 0x8, 0x1, 0x6, 0xD, 0x0, 0x9, 0x3, 0xE, 0xB, 0x4, 0x2, 0xC},
	{0x5, 0xD, 0xF, 0x6, 0x9, 0x2, 0xC, 0xA, 0xB, 0x7, 0x8, 0x1, 0x4, 0x3, 0xE, 0x0},
	{0x8, 0xE, 0x2, 0x5, 0x6, 0x9, 0x1, 0xC, 0xF, 0x4, 0xB, 0x0, 0xD, 0xA, 0x3, 0x7},
	{0x1, 0x7, 0xE, 0xD, 0x0, 0x5, 0x8, 0x3, 0x4, 0xF, 0xA, 0x6, 0x9, 0xC, 0xB, 0x2},
}

// ------------------------------------------------------------------------------
func testBlockMagmaSample() {

	if blockSize != 8 {
		log.Fatal("blockSize must be equal 8")
	}

	key := []byte{
		0x81, 0x82, 0x83, 0x84, 0x85, 0x86, 0x87, 0x88,
		0x89, 0x8a, 0x8b, 0x8c, 0x8d, 0x8e, 0x8f, 0x80,
		0xd1, 0xd2, 0xd3, 0xd4, 0xd5, 0xd6, 0xd7, 0xd8,
		0xd9, 0xda, 0xdb, 0xdc, 0xdd, 0xde, 0xdf, 0xd0,
	}
	//fmt.Printf("key: %x\n", key)

	b, err := NewCipherMagmaSBox(key, SBox2)
	checkError(err)

	blockSize := b.BlockSize()
	fmt.Println("blockSize:", blockSize)

	plaintext := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	fmt.Printf("plaintext: %x\n", plaintext)

	ciphertext := make([]byte, blockSize)

	// Encryption
	{
		b.Encrypt(ciphertext, plaintext)
		fmt.Printf("ciphertext: %x\n", ciphertext)
	}

	plaintextDec := make([]byte, blockSize)

	// Decryption
	{
		b.Decrypt(plaintextDec, ciphertext)
		fmt.Printf("plaintextDec: %x\n", plaintextDec)
	}

	if !bytes.Equal(plaintext, plaintextDec) {
		log.Fatalf("error encrypt decrypt: %x != %x", plaintext, plaintextDec)
	}
}

func testStreamMagmaCTR() {

	if blockSize != 8 {
		log.Fatal("blockSize must be equal 8")
	}

	key := make([]byte, 8*sizeOfWord)

	_, err := rand.Read(key)
	checkError(err)

	fmt.Printf("key: %x\n", key)

	b, err := NewCipherMagma(key)
	checkError(err)

	blockSize := b.BlockSize()
	fmt.Println("blockSize:", blockSize)

	iv := make([]byte, b.BlockSize())
	_, err = rand.Read(iv)
	checkError(err)

	text := "Many modern symmetric block ciphers are based on Feistel networks."

	plaintext := []byte(text)
	ciphertext := make([]byte, len(plaintext))

	// Encryption
	{
		stream := cipher.NewCTR(b, iv)
		stream.XORKeyStream(ciphertext, plaintext)
		fmt.Print("ciphertext:\n", hex.Dump(ciphertext))
	}

	plaintextDec := make([]byte, len(ciphertext))

	// Decryption
	{
		stream := cipher.NewCTR(b, iv)
		stream.XORKeyStream(plaintextDec, ciphertext)
		fmt.Printf("plaintextDec: %s\n", plaintextDec)
	}

	if !bytes.Equal(plaintext, plaintextDec) {
		log.Fatalf("error encrypt decrypt: %x != %x", plaintext, plaintextDec)
	}
}
