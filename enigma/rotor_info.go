package enigma

import (
	"fmt"
)

type RotorInfo struct {
	ID string `json:"id"`

	// Ring settings
	// ['A'..'Z']
	Ring string `json:"ring"`

	// Initial position
	// ['A'..'Z']
	Position string `json:"position"`
}

func NewRotorByInfo(ri RotorInfo) (*Rotor, error) {

	rc, ok := defaultRotors[ri.ID]
	if !ok {
		return nil, fmt.Errorf("invalid rotor id %q", ri.ID)
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
	r.SetRing(ring)

	//--------------------------------------------------------------------------
	// Parse position:
	position, err := parseIndex(ri.Position)
	if err != nil {
		return nil, fmt.Errorf("parse position: %s", err)
	}
	r.SetPosition(position)

	return r, nil
}
