package enigma

import (
	"fmt"
	"strings"
	"sync"
)

type RotorsConfig struct {
	IDs       string `json:"ids"`
	Rings     string `json:"rings"`
	Positions string `json:"positions"`
}

// rotors block
type rotorsCore struct {
	rotors    []*Rotor
	rings     []int // initial rings
	positions []int // initial positions
}

func (p *rotorsCore) reset() {
	for i, rotor := range p.rotors {
		rotor.setRing(p.rings[i])
	}
	for i, rotor := range p.rotors {
		rotor.setPosition(p.positions[i])
	}
}

func (p *rotorsCore) rotorsRotate() {
	rotateRotors(p.rotors)
}

func (p *rotorsCore) rotorsForward(index int) int {
	n := len(p.rotors)
	for i := n - 1; i >= 0; i-- {
		index = p.rotors[i].doForward(index)
	}
	return index
}

func (p *rotorsCore) rotorsBackward(index int) int {
	n := len(p.rotors)
	for i := 0; i < n; i++ {
		index = p.rotors[i].doBackward(index)
	}
	return index
}

func newRotorsCore(config RotorsConfig) (*rotorsCore, error) {

	// parse rotors:
	rotors, err := parseRotors(config.IDs)
	if err != nil {
		return nil, err
	}

	// parse rings:
	rings, err := parseLetters(config.Rings)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", "rings", err)
	}
	if len(rings) != len(rotors) {
		return nil, fmt.Errorf("invalid numbers of %s: have %d, want %d", "rings", len(rings), len(rotors))
	}

	//  parse positions:
	positions, err := parseLetters(config.Positions)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", "positions", err)
	}
	if len(positions) != len(rotors) {
		return nil, fmt.Errorf("invalid numbers of %s: have %d, want %d", "positions", len(positions), len(rotors))
	}

	rc := &rotorsCore{
		rotors:    rotors,
		rings:     rings,
		positions: positions,
	}

	rc.reset()

	return rc, nil
}

func parseRotors(s string) ([]*Rotor, error) {
	rotorIDs := strings.Split(s, " ")
	rs := make([]*Rotor, len(rotorIDs))
	for i, rotorID := range rotorIDs {
		rc, err := rotorByID(rotorID)
		if err != nil {
			return nil, err
		}
		r, err := NewRotor(rc)
		if err != nil {
			return nil, err
		}
		rs[i] = r
	}
	return rs, nil
}

//------------------------------------------------------------------------------

// Historical rotors:

var rotorsGuard sync.RWMutex

var rotorsMap = map[string]RotorConfig{
	"I": RotorConfig{
		Wiring:    "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		Turnovers: "Q",
	},
	"II": RotorConfig{
		Wiring:    "AJDKSIRUXBLHWTMCQGZNPYFVOE",
		Turnovers: "E",
	},
	"III": RotorConfig{
		Wiring:    "BDFHJLCPRTXVZNYEIWGAKMUSQO",
		Turnovers: "V",
	},
	"IV": RotorConfig{
		Wiring:    "ESOVPZJAYQUIRHXLNFTGKDCMWB",
		Turnovers: "J",
	},
	"V": RotorConfig{
		Wiring:    "VZBRGITYUPSDNHLXAWMJQOFECK",
		Turnovers: "Z",
	},
	"VI": RotorConfig{
		Wiring:    "JPGVOUMFYQBENHZRDKASXLICTW",
		Turnovers: "ZM",
	},
	"VII": RotorConfig{
		Wiring:    "NZJHGRCXMYSWBOUFAIVLPEKQDT",
		Turnovers: "ZM",
	},
	"VIII": RotorConfig{
		Wiring:    "FKQHTLXOCBJSPDZRAMEWNIUYGV",
		Turnovers: "ZM",
	},
	"Beta": RotorConfig{
		Wiring:    "LEYJVCNIXWPBQMDRTAKZGFUHOS",
		Turnovers: "",
	},
	"Gamma": RotorConfig{
		Wiring:    "FSOKANUERHMBTIYCWLQPZXVGJD",
		Turnovers: "",
	},
}

func RegisterRotor(rotorID string, rc RotorConfig) {
	rotorsGuard.Lock()
	defer rotorsGuard.Unlock()
	rotorsMap[rotorID] = rc
}

func rotorByID(rotorID string) (RotorConfig, error) {
	rotorsGuard.RLock()
	defer rotorsGuard.RUnlock()
	rc, ok := rotorsMap[rotorID]
	if !ok {
		return rc, fmt.Errorf("There is no rotor ID (%q)", rotorID)
	}
	return rc, nil
}
