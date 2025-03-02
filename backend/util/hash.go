package util

import (
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

func RandomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Generates a hash using salt.
// Default options: 1 iteration, 64MB, 4 threads, 32 key length
func GenerateHash(password []byte, salt []byte) []byte {
	return argon2.IDKey(password, salt, 1, 64*1024, 4, 32)
}
