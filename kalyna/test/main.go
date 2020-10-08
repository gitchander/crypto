package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/gitchander/crypto/kalyna"
)

func main() {
	testCipher()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testCipher() {

	config := kalyna.Config{
		BlockSize: 128,
		KeySize:   128,
	}

	key, err := hex.DecodeString("000102030405060708090A0B0C0D0E0F")
	checkError(err)

	plaintext, err := hex.DecodeString("101112131415161718191A1B1C1D1E1F")
	checkError(err)

	block, err := config.NewCipher(key)
	checkError(err)

	fmt.Println("BlockSize:", block.BlockSize())

	ciphertext := make([]byte, block.BlockSize())

	block.Encrypt(ciphertext, plaintext)

	fmt.Printf("CIPHERTEXT: %X\n", ciphertext)

	// CIPHERTEXT: 81BF1C7D779BAC20E1C9EA39B4D2AD06
}
