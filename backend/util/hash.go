package util

import (
	"crypto/rand"
	"math/big"
	"time"

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
// Default options: 1 iteration, 64MB, 4 threads, 32 key length.
// Also tries to prevent timing attacks.
func GenerateHash(password []byte, salt []byte) ([]byte, error) {
	toSleep, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Millisecond * time.Duration(toSleep.Add(toSleep, big.NewInt(50)).Int64()))
	return argon2.IDKey(password, salt, 1, 64*1024, 4, 32), nil
}
