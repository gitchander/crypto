package enigma

var defaultRotors = map[string]RotorConfig{
	"I": RotorConfig{
		Name:      "Rotor I",
		Wiring:    "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		Turnovers: "Q",
	},
	"II": RotorConfig{
		Name:      "Rotor II",
		Wiring:    "AJDKSIRUXBLHWTMCQGZNPYFVOE",
		Turnovers: "E",
	},
	"III": RotorConfig{
		Name:      "Rotor III",
		Wiring:    "BDFHJLCPRTXVZNYEIWGAKMUSQO",
		Turnovers: "V",
	},
	"IV": RotorConfig{
		Name:      "Rotor IV",
		Wiring:    "ESOVPZJAYQUIRHXLNFTGKDCMWB",
		Turnovers: "J",
	},
	"V": RotorConfig{
		Name:      "Rotor V",
		Wiring:    "VZBRGITYUPSDNHLXAWMJQOFECK",
		Turnovers: "Z",
	},
	"VI": RotorConfig{
		Name:      "Rotor VI",
		Wiring:    "JPGVOUMFYQBENHZRDKASXLICTW",
		Turnovers: "ZM",
	},
	"VII": RotorConfig{
		Name:      "Rotor VII",
		Wiring:    "NZJHGRCXMYSWBOUFAIVLPEKQDT",
		Turnovers: "ZM",
	},
	"VIII": RotorConfig{
		Name:      "Rotor VIII",
		Wiring:    "FKQHTLXOCBJSPDZRAMEWNIUYGV",
		Turnovers: "ZM",
	},
	"Beta": RotorConfig{
		Name:      "Rotor Beta",
		Wiring:    "LEYJVCNIXWPBQMDRTAKZGFUHOS",
		Turnovers: "",
	},
	"Gamma": RotorConfig{
		Name:      "Rotor Gamma",
		Wiring:    "FSOKANUERHMBTIYCWLQPZXVGJD",
		Turnovers: "",
	},
}
