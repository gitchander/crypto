package enigma

import (
	"fmt"

	"github.com/gitchander/crypto/enigma/ecore"
)

// Historical reflectors:
var historicalReflectors = map[string]ecore.ReflectorConfig{
	"A": ecore.ReflectorConfig{
		Wiring: "EJMZALYXVBWFCRQUONTSPIKHGD",
	},
	"B": ecore.ReflectorConfig{
		Wiring: "YRUHQSLDPXNGOKMIEBFZCWVJAT",
	},
	"C": ecore.ReflectorConfig{
		Wiring: "FVPJIAOYEDRZXWGCTKUQSBNMHL",
	},
	"B-thin": ecore.ReflectorConfig{
		Wiring: "ENKQAUYWJICOPBLMDXZVFTHRGS",
	},
	"C-thin": ecore.ReflectorConfig{
		Wiring: "RDOBJNTKVEHMLFCWZAXGYIPSUQ",
	},
}

func reflectorByID(id string) (*ecore.Reflector, error) {
	rc, ok := historicalReflectors[id]
	if !ok {
		return nil, fmt.Errorf("invalid reflector id %q", id)
	}
	return ecore.NewReflector(rc)
}
