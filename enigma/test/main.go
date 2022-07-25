package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gitchander/crypto/enigma"
)

func main() {
	testMarshalJSON()
}

func testMarshalJSON() {
	c := enigma.Config{
		Plugboard: "BQ CR DI EJ KW MT OS PX UZ GH",
		Rotors: []enigma.RotorInfo{
			{
				ID:       "I",
				Ring:     "A",
				Position: "A",
			},
			{
				ID:       "II",
				Ring:     "A",
				Position: "A",
			},
			{
				ID:       "III",
				Ring:     "A",
				Position: "A",
			},
		},
		ReflectorID: "A",
	}

	data, err := json.MarshalIndent(c, "", "\t")
	checkError(err)

	fmt.Println(string(data))

	e, err := enigma.New(c)
	checkError(err)

	plaintext := "ABCDEFGH"
	ciphertext := e.FeedString(plaintext)

	fmt.Println("plaintext: ", plaintext)
	fmt.Println("ciphertext:", ciphertext)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
