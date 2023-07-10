package ecore

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	regexpPosition      = regexp.MustCompile("^[A-Z]$")
	regexpWiring        = regexp.MustCompile(fmt.Sprintf("^([A-Z]{%d})$", totalIndexes))
	regexpPlugboardPair = regexp.MustCompile("^([A-Z]{2})$")
)

func ValidateWiring(wiring string) error {
	if !(regexpWiring.MatchString(wiring)) {
		return fmt.Errorf("wiring is invalid %q", wiring)
	}
	// Check duplicates:
	if d, ok := findDuplicate([]rune(wiring)); ok {
		return fmt.Errorf("wiring has duplicates %q", d)
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
	d, ok := findDuplicate([]rune(strings.Join(pairs, "")))
	if ok {
		return fmt.Errorf("plugboard has duplicates %q", d)
	}
	return nil
}
