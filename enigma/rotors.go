package enigma

// Historical rotors:
var historicalRotors = map[string]RotorConfig{
	"I": RotorConfig{
		Name:      "I",
		Wiring:    "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		Turnovers: "Q",
	},
	"II": RotorConfig{
		Name:      "II",
		Wiring:    "AJDKSIRUXBLHWTMCQGZNPYFVOE",
		Turnovers: "E",
	},
	"III": RotorConfig{
		Name:      "III",
		Wiring:    "BDFHJLCPRTXVZNYEIWGAKMUSQO",
		Turnovers: "V",
	},
	"IV": RotorConfig{
		Name:      "IV",
		Wiring:    "ESOVPZJAYQUIRHXLNFTGKDCMWB",
		Turnovers: "J",
	},
	"V": RotorConfig{
		Name:      "V",
		Wiring:    "VZBRGITYUPSDNHLXAWMJQOFECK",
		Turnovers: "Z",
	},
	"VI": RotorConfig{
		Name:      "VI",
		Wiring:    "JPGVOUMFYQBENHZRDKASXLICTW",
		Turnovers: "ZM",
	},
	"VII": RotorConfig{
		Name:      "VII",
		Wiring:    "NZJHGRCXMYSWBOUFAIVLPEKQDT",
		Turnovers: "ZM",
	},
	"VIII": RotorConfig{
		Name:      "VIII",
		Wiring:    "FKQHTLXOCBJSPDZRAMEWNIUYGV",
		Turnovers: "ZM",
	},
	"Beta": RotorConfig{
		Name:      "Beta",
		Wiring:    "LEYJVCNIXWPBQMDRTAKZGFUHOS",
		Turnovers: "",
	},
	"Gamma": RotorConfig{
		Name:      "Gamma",
		Wiring:    "FSOKANUERHMBTIYCWLQPZXVGJD",
		Turnovers: "",
	},
}
