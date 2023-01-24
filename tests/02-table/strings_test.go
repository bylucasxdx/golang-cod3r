package strings

import (
	"strings"
	"testing"
)

const errorMessage = "%s (part: %s) - expected index: (%d) <> founded (%d)"

func TestIndex(t *testing.T) {
	tests := []struct {
		text     string
		part     string
		expected int
	}{
		{"Cod3r course", "Cod3r", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Lucas", "c", 2},
	}

	for _, test := range tests {
		t.Logf("Data: %v", test)
		actual := strings.Index(test.text, test.part)

		if actual != test.expected {
			t.Errorf(errorMessage, test.text, test.part, test.expected, actual)
		}
	}
}
