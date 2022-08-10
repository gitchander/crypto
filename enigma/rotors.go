package enigma

// Historical rotors:
var historicalRotors = map[string]RotorConfig{
	"I": RotorConfig{
		Wiring:    "EKMFLGDQVZNTOWYHXUSPAIBRCJ",
		Turnovers: "Q",
	},
	"II": RotorConfig{
		Wiring:    "AJDKSIRUXBLHWTMCQGZNPYFVOE",
		Turnovers: "E",
	},
	"III": RotorConfig{
		Wiring:    "BDFHJLCPRTXVZNYEIWGAKMUSQO",
		Turnovers: "V",
	},
	"IV": RotorConfig{
		Wiring:    "ESOVPZJAYQUIRHXLNFTGKDCMWB",
		Turnovers: "J",
	},
	"V": RotorConfig{
		Wiring:    "VZBRGITYUPSDNHLXAWMJQOFECK",
		Turnovers: "Z",
	},
	"VI": RotorConfig{
		Wiring:    "JPGVOUMFYQBENHZRDKASXLICTW",
		Turnovers: "ZM",
	},
	"VII": RotorConfig{
		Wiring:    "NZJHGRCXMYSWBOUFAIVLPEKQDT",
		Turnovers: "ZM",
	},
	"VIII": RotorConfig{
		Wiring:    "FKQHTLXOCBJSPDZRAMEWNIUYGV",
		Turnovers: "ZM",
	},
	"Beta": RotorConfig{
		Wiring:    "LEYJVCNIXWPBQMDRTAKZGFUHOS",
		Turnovers: "",
	},
	"Gamma": RotorConfig{
		Wiring:    "FSOKANUERHMBTIYCWLQPZXVGJD",
		Turnovers: "",
	},
}
