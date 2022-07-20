package enigma

// +---+----+-----+----+---+-----+-----+------+
// | I | II | III | IV | V | VI  | VII | VIII |
// +---+----+-----+----+---+-----+-----+------+
// | Q | E  | V   | J  | Z | Z,M | Z,M | Z,M  |
// +---+----+-----+----+---+-----+-----+------+

// let id       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// let rotorI   = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
// let rotorII  = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
// let rotorIII = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
// *NewRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ", "I", "Q"),
// *NewRotor("AJDKSIRUXBLHWTMCQGZNPYFVOE", "II", "E"),
// *NewRotor("BDFHJLCPRTXVZNYEIWGAKMUSQO", "III", "V"),
// *NewRotor("ESOVPZJAYQUIRHXLNFTGKDCMWB", "IV", "J"),
// *NewRotor("VZBRGITYUPSDNHLXAWMJQOFECK", "V", "Z"),
// *NewRotor("JPGVOUMFYQBENHZRDKASXLICTW", "VI", "ZM"),
// *NewRotor("NZJHGRCXMYSWBOUFAIVLPEKQDT", "VII", "ZM"),
// *NewRotor("FKQHTLXOCBJSPDZRAMEWNIUYGV", "VIII", "ZM"),

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
}
