package toolkit

import (
	"crypto/rand"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+=!@#$%^&*()[]{}|;:,.<>?`~"

// RandomString generates a random string of the given length which is cryptographically secure.
func (t *Tools) RandomString(length int) string {
	s, r := make([]rune, length), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}

// RandomBytes generates a random byte slice of the given length which is cryptographically secure.
func (t *Tools) RandomBytes(length int) []byte {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil
	}

	return b
}
