package youtube

import (
	"testing"
)

func TestIsValidUrl(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param string
		expected bool
	} {
		{"test", false},
		{"watcher?v=KDPW_g2AhAU", false},
		{"https://www.youtube.com/test", false},
		{"watcher?v=KDPW_g2AhAU", false},
		{"https://www.youtube.com/watch?v=KDPW_g2AhAU", true},
		{"watch?v=KDPW_g2AhAU", true},
	}

	for _, test := range tests {
		result := isValidUrl(test.param)
		if result != test.expected {
			t.Error(
				"For", test.param,
				"expected", test.expected,
				"got", result,
			)
		}
	}
}
