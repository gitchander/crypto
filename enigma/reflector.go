package enigma

import (
	"fmt"
)

type ReflectorConfig struct {
	Wiring string
}

type Reflector struct {
	ct coupleTable
}

func NewReflector(rc ReflectorConfig) (*Reflector, error) {
	ct, err := parseWiring(rc.Wiring)
	if err != nil {
		return nil, err
	}
	r := &Reflector{
		ct: ct,
	}
	return r, nil
}

func NewReflectorByID(id string) (*Reflector, error) {
	rc, ok := historicalReflectors[id]
	if !ok {
		return nil, fmt.Errorf("invalid reflector id %q", id)
	}
	return NewReflector(rc)
}

func (r *Reflector) do(index int) int {
	return r.ct.forwardTable[index]
}
