package enigma

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	regexpPosition      = regexp.MustCompile("^[A-Z]$")
	regexpWiring        = regexp.MustCompile(fmt.Sprintf("^([A-Z]{%d})$", positions))
	regexpPlugboardPair = regexp.MustCompile("^([A-Z]{2})$")
)

func ValidatePosition(s string) error {
	if !regexpPosition.MatchString(s) {
		return fmt.Errorf("invalid position %q", s)
	}
	return nil
}

func ValidateWiring(wiring string) error {
	if !regexpWiring.MatchString(wiring) {
		return fmt.Errorf("wiring is invalid %q", wiring)
	}
	// Check duplicates:
	if rs := duplicateRunes(wiring); len(rs) > 0 {
		return fmt.Errorf("wiring has duplicates %q", rs)
	}
	return nil
}

func ValidatePlugboard(s string) error {
	if s == "" {
		return nil
	}
	pairs := strings.Split(s, " ")
	for i, pair := range pairs {
		if !regexpPlugboardPair.MatchString(pair) {
			return fmt.Errorf("plugboard pair[%d]: invalid pair value %q", i, pair)
		}
	}
	// Check duplicates:
	rs := duplicateRunes(strings.Join(pairs, ""))
	if len(rs) > 0 {
		return fmt.Errorf("plugboard has duplicates %q", rs)
	}
	return nil
}
