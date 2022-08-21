package enigma

import (
	"fmt"
)

// Ring and Position possible values: ["A".."Z"]

type RotorInfo struct {
	ID       string `json:"id"`
	Ring     string `json:"ring"`
	Position string `json:"position"`
}

func NewRotorByInfo(ri RotorInfo) (*Rotor, error) {

	rc, ok := historicalRotors[ri.ID]
	if !ok {
		return nil, fmt.Errorf("Unknown rotor by id %q", ri.ID)
	}
	r, err := NewRotor(rc)
	if err != nil {
		return nil, err
	}

	//--------------------------------------------------------------------------
	// Parse ring:
	ring, err := parseIndex(ri.Ring)
	if err != nil {
		return nil, fmt.Errorf("parse ring: %s", err)
	}
	r.setRing(ring)

	//--------------------------------------------------------------------------
	// Parse position:
	position, err := parseIndex(ri.Position)
	if err != nil {
		return nil, fmt.Errorf("parse position: %s", err)
	}
	r.setPosition(position)

	return r, nil
}
