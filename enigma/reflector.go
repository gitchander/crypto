package enigma

import (
	"fmt"
)

type ReflectorConfig struct {
	Name   string
	Wiring string
}

type Reflector struct {
	direct convertTable
}

func NewReflector(rc ReflectorConfig) (*Reflector, error) {
	dr, err := parseWiring(rc.Wiring)
	if err != nil {
		return nil, err
	}
	return &Reflector{direct: dr.direct}, nil
}

func NewReflectorByID(id string) (*Reflector, error) {
	rc, ok := historicalReflectors[id]
	if !ok {
		return nil, fmt.Errorf("invalid reflector id %q", id)
	}
	return NewReflector(rc)
}

func (r *Reflector) Direct(index int) int {
	return r.direct[index]
}
