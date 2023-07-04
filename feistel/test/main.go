package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	testBlock()
	testStreamCTR()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testBlock() {

	key := make([]byte, 5*sizeOfWord)

	_, err := rand.Read(key)
	checkError(err)

	fmt.Printf("key: %x\n", key)

	b, err := NewCipher(key)
	checkError(err)

	blockSize := b.BlockSize()
	fmt.Println("blockSize:", blockSize)

	text := "Many modern symmetric block ciphers are based on Feistel networks."

	plaintext := []byte(text[:blockSize])
	ciphertext := make([]byte, blockSize)

	// Encryption
	{
		b.Encrypt(ciphertext, plaintext)
		fmt.Print("ciphertext:\n", hex.Dump(ciphertext))
	}

	plaintextDec := make([]byte, blockSize)

	// Decryption
	{
		b.Decrypt(plaintextDec, ciphertext)
		fmt.Printf("plaintextDec: %s\n", plaintextDec)
	}

	if !bytes.Equal(plaintext, plaintextDec) {
		log.Fatalf("error encrypt decrypt: %x != %x", plaintext, plaintextDec)
	}
}

func testStreamCTR() {

	key := make([]byte, 7*sizeOfWord)

	_, err := rand.Read(key)
	checkError(err)

	fmt.Printf("key: %x\n", key)

	b, err := NewCipher(key)
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
