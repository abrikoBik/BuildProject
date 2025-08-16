package util

import (
	"crypto/rand"
	"crypto/sha256"
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
// Default options are aligned with OWASP's recommendations.
// Also tries to prevent timing attacks (sleeps for 0-50ms).
func GenerateHash(password []byte, salt []byte) ([]byte, error) {
	toSleep, err := rand.Int(rand.Reader, big.NewInt(50))
	if err != nil {
		return nil, err
	}
	time.Sleep(time.Millisecond * time.Duration(toSleep.Int64()))
	return argon2.IDKey(password, salt, 3, 12288, 1, 40), nil
}

func GenerateChecksum(data []byte) ([]byte, error) {
	hash := sha256.New()
	_, err := hash.Write(data)
	if err != nil {
		return nil, err
	}
	return hash.Sum(nil), nil
}
