package enigma

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gitchander/crypto/enigma/ecore"
)

type RotorsConfig struct {
	IDs       string `json:"ids"`
	Rings     string `json:"rings"`
	Positions string `json:"positions"`
}

//------------------------------------------------------------------------------

type syncRotors struct {
	guard  sync.RWMutex
	rotors map[string]ecore.RotorConfig
}

func (p *syncRotors) RegisterRotor(rotorID string, rc ecore.RotorConfig) error {

	// Check rotor config
	_, err := ecore.NewRotor(rc)
	if err != nil {
		return err
	}

	p.guard.Lock()
	defer p.guard.Unlock()

	p.rotors[rotorID] = rc
	return nil
}

func (p *syncRotors) RotorByID(rotorID string) (*ecore.Rotor, error) {

	p.guard.Lock()
	defer p.guard.Unlock()

	rc, ok := p.rotors[rotorID]
	if !ok {
		return nil, fmt.Errorf("There is no rotor ID (%q)", rotorID)
	}

	return ecore.NewRotor(rc)
}

//------------------------------------------------------------------------------

var guardRotors sync.RWMutex

// Historical rotors:
var globalRotors = map[string]ecore.RotorConfig{
	"I": ecore.RotorConfig{
		Wiring:    "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		Turnovers: "Q",
	},
	"II": ecore.RotorConfig{
		Wiring:    "AJDKSIRUXBLHWTMCQGZNPYFVOE",
		Turnovers: "E",
	},
	"III": ecore.RotorConfig{
		Wiring:    "BDFHJLCPRTXVZNYEIWGAKMUSQO",
		Turnovers: "V",
	},
	"IV": ecore.RotorConfig{
		Wiring:    "ESOVPZJAYQUIRHXLNFTGKDCMWB",
		Turnovers: "J",
	},
	"V": ecore.RotorConfig{
		Wiring:    "VZBRGITYUPSDNHLXAWMJQOFECK",
		Turnovers: "Z",
	},
	"VI": ecore.RotorConfig{
		Wiring:    "JPGVOUMFYQBENHZRDKASXLICTW",
		Turnovers: "ZM",
	},
	"VII": ecore.RotorConfig{
		Wiring:    "NZJHGRCXMYSWBOUFAIVLPEKQDT",
		Turnovers: "ZM",
	},
	"VIII": ecore.RotorConfig{
		Wiring:    "FKQHTLXOCBJSPDZRAMEWNIUYGV",
		Turnovers: "ZM",
	},
	"Beta": ecore.RotorConfig{
		Wiring:    "LEYJVCNIXWPBQMDRTAKZGFUHOS",
		Turnovers: "",
	},
	"Gamma": ecore.RotorConfig{
		Wiring:    "FSOKANUERHMBTIYCWLQPZXVGJD",
		Turnovers: "",
	},
}

func RegisterRotor(rotorID string, rc ecore.RotorConfig) {
	guardRotors.Lock()
	defer guardRotors.Unlock()

	globalRotors[rotorID] = rc
}

func rotorConfigByID(rotorID string) (ecore.RotorConfig, error) {
	guardRotors.RLock()
	defer guardRotors.RUnlock()

	rc, ok := globalRotors[rotorID]
	if !ok {
		var zeroValue ecore.RotorConfig
		return zeroValue, fmt.Errorf("There is no rotor by ID (%q)", rotorID)
	}
	return rc, nil
}

//------------------------------------------------------------------------------

func parseRotorConfigs(s string) ([]ecore.RotorConfig, error) {
	rotorIDs := strings.Split(s, " ")
	rcs := make([]ecore.RotorConfig, len(rotorIDs))
	for i, rotorID := range rotorIDs {
		rc, err := rotorConfigByID(rotorID)
		if err != nil {
			return nil, err
		}
		rcs[i] = rc
	}
	return rcs, nil
}

func newRotorBlock(rc RotorsConfig) (*ecore.RotorBlock, error) {
	rcs, err := parseRotorConfigs(rc.IDs)
	if err != nil {
		return nil, err
	}
	rbc := ecore.RotorBlockConfig{
		Rotors:    rcs,
		Rings:     rc.Rings,
		Positions: rc.Positions,
	}
	return ecore.NewRotorBlock(rbc)
}
