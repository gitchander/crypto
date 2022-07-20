package enigma

var (
	ReflectorA = ReflectorConfig{
		Name:   "Reflector A",
		Wiring: "EJMZALYXVBWFCRQUONTSPIKHGD",
	}
	ReflectorB = ReflectorConfig{
		Name:   "Reflector B",
		Wiring: "YRUHQSLDPXNGOKMIEBFZCWVJAT",
	}
	ReflectorC = ReflectorConfig{
		Name:   "Reflector C",
		Wiring: "FVPJIAOYEDRZXWGCTKUQSBNMHL",
	}
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

func (r *Reflector) Direct(index int) int {
	return r.direct[index]
}
