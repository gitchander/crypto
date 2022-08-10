package enigma

import (
	"fmt"
)

type ReflectorConfig struct {
	Wiring string
}

type Reflector struct {
	convertTable
}

func NewReflector(rc ReflectorConfig) (*Reflector, error) {
	dr, err := parseWiring(rc.Wiring)
	if err != nil {
		return nil, err
	}
	return &Reflector{convertTable: dr.direct}, nil
}

func NewReflectorByID(id string) (*Reflector, error) {
	rc, ok := historicalReflectors[id]
	if !ok {
		return nil, fmt.Errorf("invalid reflector id %q", id)
	}
	return NewReflector(rc)
}
