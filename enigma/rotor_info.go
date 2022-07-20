package enigma

import (
	"fmt"
)

type RotorInfo struct {
	ID string

	// Ring settings
	// ['A'..'Z']
	Ring byte

	// Initial position
	// ['A'..'Z']
	Position byte
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
	// Parse ring
	ring, ok := letterToIndex(ri.Ring)
	if !ok {
		return nil, fmt.Errorf("invalid ring %#U", ri.Ring)
	}
	r.SetRing(ring)

	//--------------------------------------------------------------------------
	// Parse position
	position, ok := letterToIndex(ri.Position)
	if !ok {
		return nil, fmt.Errorf("invalid position %#U", ri.Position)
	}
	r.SetPosition(position)

	return r, nil
}

func parseIndex(b byte) (int, error) {
	index, ok := letterToIndex(b)
	if !ok {
		return 0, errInvalidLetter(b)
	}
	return index, nil
}
