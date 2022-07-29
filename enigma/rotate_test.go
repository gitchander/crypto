package enigma

import (
	"fmt"
	"testing"
)

// https://en.wikipedia.org/wiki/Enigma_rotor_details

type enigmaPos struct {
	e *Enigma
}

func newEnigmaPos(e *Enigma) *enigmaPos {
	return &enigmaPos{e}
}

func (e *enigmaPos) Rotate() {
	e.e.rotate()
}

func (e *enigmaPos) Positions() string {
	rotors := e.e.rotors
	ls := make([]byte, len(rotors))
	for i, r := range rotors {
		letter, _ := indexToLetter(r.Position())
		ls[i] = letter
	}
	return string(ls)
}

func (e *enigmaPos) SetPositions(s string) {
	bs := []byte(s)
	if len(bs) < len(e.e.rotors) {
		err := fmt.Errorf("insufficient positions length: have %d, want %d",
			len(bs), len(e.e.rotors))
		panic(err)
	}
	for i, r := range e.e.rotors {
		letter := bs[i]
		index, ok := letterToIndex(letter)
		if !ok {
			panic(errInvalidLetter(letter))
		}
		r.SetPosition(index)
	}
}

func TestRotate(t *testing.T) {

	samples := [][]string{
		// Normal sequence:
		{
			"AAU", // normal step of right rotor
			"AAV", // right rotor (III) goes in V—notch position
			"ABW", // right rotor takes middle rotor one step further
			"ABX", // normal step of right rotor
		},
		// Double step sequence:
		{
			"ADU", // normal step of right rotor
			"ADV", // right rotor (III) goes in V—notch position
			"AEW", // right rotor steps, takes middle rotor (II) one step further, which is now in its own E—notch position
			"BFX", // normal step of right rotor, double step of middle rotor, normal step of left rotor
			"BFY", // normal step of right rotor
		},
	}

	c := Config{
		Rotors: []RotorInfo{
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
		ReflectorID: "B",
	}
	e, err := New(c)
	if err != nil {
		t.Fatal(err)
	}

	te := newEnigmaPos(e)

	for _, sample := range samples {
		te.SetPositions(sample[0])
		for _, wantPositions := range sample {
			havePositions := te.Positions()
			if havePositions != wantPositions {
				t.Fatalf("invalid rotors positions: have %s, want %s",
					havePositions, wantPositions)
			}
			te.Rotate()
		}
	}
}
