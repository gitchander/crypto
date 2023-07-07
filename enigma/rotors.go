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

// rotors block
type rotorsCore struct {
	rotors    []*ecore.Rotor
	rings     []int // initial rings
	positions []int // initial positions
}

func newRotorsCore(config RotorsConfig) (*rotorsCore, error) {

	// parse rotors:
	rotors, err := parseRotors(config.IDs)
	if err != nil {
		return nil, err
	}

	// parse rings:
	rings, err := ecore.ParseIndexes(config.Rings)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", "rings", err)
	}
	if len(rings) != len(rotors) {
		return nil, fmt.Errorf("invalid numbers of %s: have %d, want %d", "rings", len(rings), len(rotors))
	}

	//  parse positions:
	positions, err := ecore.ParseIndexes(config.Positions)
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

func (p *rotorsCore) reset() {
	for i, rotor := range p.rotors {
		rotor.SetRing(p.rings[i])
	}
	for i, rotor := range p.rotors {
		rotor.SetPosition(p.positions[i])
	}
}

func (p *rotorsCore) rotate() {
	ecore.RotateRotors(p.rotors)
}

func (p *rotorsCore) doForward(index int) int {
	n := len(p.rotors)
	for i := n - 1; i >= 0; i-- {
		index = p.rotors[i].DoForward(index)
	}
	return index
}

func (p *rotorsCore) doBackward(index int) int {
	n := len(p.rotors)
	for i := 0; i < n; i++ {
		index = p.rotors[i].DoBackward(index)
	}
	return index
}

//------------------------------------------------------------------------------

type syncRotors struct {
	guard  sync.RWMutex
	rotors map[string]ecore.RotorConfig
}

func (p *syncRotors) RegisterRotor(rotorID string, rc ecore.RotorConfig) error {

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

var rotorsGuard sync.RWMutex

// Historical rotors:
var rotorsMap = map[string]ecore.RotorConfig{
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
	rotorsGuard.Lock()
	defer rotorsGuard.Unlock()

	rotorsMap[rotorID] = rc
}

func rotorByID(rotorID string) (*ecore.Rotor, error) {
	rotorsGuard.RLock()
	defer rotorsGuard.RUnlock()

	rc, ok := rotorsMap[rotorID]
	if !ok {
		return nil, fmt.Errorf("There is no rotor ID (%q)", rotorID)
	}
	return ecore.NewRotor(rc)
}

func parseRotors(s string) ([]*ecore.Rotor, error) {
	rotorIDs := strings.Split(s, " ")
	rs := make([]*ecore.Rotor, len(rotorIDs))
	for i, rotorID := range rotorIDs {
		r, err := rotorByID(rotorID)
		if err != nil {
			return nil, err
		}
		rs[i] = r
	}
	return rs, nil
}
