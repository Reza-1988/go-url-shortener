package shortener

import (
	"crypto/rand"
	"math/big"
)

// base62Alphabet is the allowed characters for short codes (0-9, A-Z, a-z).
const base62Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// RandomBase62 generates a random Base62 string of the given length.
// It uses crypto/rand, so it is safe for tokens and short links.
func RandomBase62(length int) (string, error) {
	// For invalid length, return an empty string (no error).
	if length <= 0 {
		return "", nil
	}

	// max is the upper bound for random numbers: [0, len(alphabet)).
	max := big.NewInt(int64(len(base62Alphabet)))

	// Pre-allocate the output bytes for speed and simplicity.
	out := make([]byte, length)

	for i := 0; i < length; i++ {
		// Pick a secure random index into the alphabet.
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		out[i] = base62Alphabet[n.Int64()]
	}

	return string(out), nil
}
