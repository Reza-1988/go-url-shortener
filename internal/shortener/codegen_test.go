package shortener

import "testing"

// TestRandomBase62_LengthAndCharset checks two basic rules:
// 1) the output has the requested length
// 2) all characters are in the Base62 alphabet
func TestRandomBase62_LengthAndCharset(t *testing.T) {
	const length = 7

	// Run multiple times to reduce the chance of missing a random edge case.
	for i := 0; i < 200; i++ {
		s, err := RandomBase62(length)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		// Length must match exactly.
		if len(s) != length {
			t.Fatalf("expected length %d, got %d", length, len(s))
		}

		// Every character must be Base62.
		for _, ch := range s {
			if !isBase62Char(byte(ch)) {
				t.Fatalf("unexpected char: %q in %q", ch, s)
			}
		}
	}
}

// isBase62Char returns true if b exists in the Base62 alphabet.
func isBase62Char(b byte) bool {
	// Simple linear scan is fine here (alphabet is tiny: 62 chars).
	for i := 0; i < len(base62Alphabet); i++ {
		if base62Alphabet[i] == b {
			return true
		}
	}
	return false
}
