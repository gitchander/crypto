package enigma

import (
	"fmt"
	"sync"

	"github.com/gitchander/crypto/enigma/ecore"
)

// type safeReflectors struct {
// 	guard sync.RWMutex
// 	reflectors map[string]ecore.ReflectorConfig
// }

var guardReflectors sync.RWMutex

// Historical reflectors:
var globalReflectors = map[string]ecore.ReflectorConfig{
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

func RegisterReflector(reflectorID string, rc ecore.ReflectorConfig) {

	guardReflectors.Lock()
	defer guardReflectors.Unlock()

	globalReflectors[reflectorID] = rc
}

func reflectorConfigByID(id string) (ecore.ReflectorConfig, error) {

	guardReflectors.Lock()
	defer guardReflectors.Unlock()

	rc, ok := globalReflectors[id]
	if !ok {
		var zeroValue ecore.ReflectorConfig
		return zeroValue, fmt.Errorf("invalid reflector id %q", id)
	}
	return rc, nil
}

func newReflectorByID(id string) (*ecore.Reflector, error) {

	rc, err := reflectorConfigByID(id)
	if err != nil {
		return nil, err
	}

	return ecore.NewReflector(rc)
}
