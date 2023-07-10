package ecore

import (
	"fmt"
)

type RotorBlockConfig struct {
	Rotors    []RotorConfig
	Rings     string
	Positions string
}

type RotorBlock struct {
	rotors    []*Rotor
	rings     []int // initial rings
	positions []int // initial positions
}

func NewRotorBlock(config RotorBlockConfig) (*RotorBlock, error) {

	rotors := make([]*Rotor, len(config.Rotors))
	for i, rc := range config.Rotors {
		r, err := NewRotor(rc)
		if err != nil {
			return nil, err
		}
		rotors[i] = r
	}

	// parse rings:
	rings, err := ParseIndexes(config.Rings)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", "rings", err)
	}
	if len(rings) != len(rotors) {
		return nil, fmt.Errorf("invalid numbers of %s: have %d, want %d", "rings", len(rings), len(rotors))
	}

	//  parse positions:
	positions, err := ParseIndexes(config.Positions)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", "positions", err)
	}
	if len(positions) != len(rotors) {
		return nil, fmt.Errorf("invalid numbers of %s: have %d, want %d", "positions", len(positions), len(rotors))
	}

	rb := &RotorBlock{
		rotors:    rotors,
		rings:     rings,
		positions: positions,
	}
	rb.Reset()

	return rb, nil
}

func (p *RotorBlock) Rotors() []*Rotor {
	return p.rotors
}

func (p *RotorBlock) Reset() {
	for i, rotor := range p.rotors {
		rotor.SetRing(p.rings[i])
	}
	for i, rotor := range p.rotors {
		rotor.SetPosition(p.positions[i])
	}
}

func (p *RotorBlock) Rotate() {
	RotateRotors(p.rotors)
}

func (p *RotorBlock) Forward(index int) int {
	n := len(p.rotors)
	for i := n - 1; i >= 0; i-- {
		index = p.rotors[i].Forward(index)
	}
	return index
}

func (p *RotorBlock) Backward(index int) int {
	n := len(p.rotors)
	for i := 0; i < n; i++ {
		index = p.rotors[i].Backward(index)
	}
	return index
}

func RotateRotors(rs []*Rotor) {
	hasPrev := true // Turnover last rotor.
	for i := len(rs) - 1; i >= 0; i-- {
		r := rs[i]
		ok := r.hasTurnover()
		if hasPrev || ok {
			r.rotate()
		}
		hasPrev = ok
	}
}
