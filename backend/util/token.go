package util

import (
	// "mp1/db"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Generates a JWT token with name and password in Claims.
// Also generates a 64 byte large salt.
func GenerateJWT(name string, password []byte, key string) (string, error) {
	salt, err := RandomBytes(64)
	if err != nil {
		return "", err
	}
	hash, err := GenerateHash([]byte(password), salt)
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iss":  "mp1-auth",
		"iat":  now,
		"exp":  now.Add(time.Hour * 24 * 7 * 5), // 5 weeks
		"sub":  name,
		"hash": hash,
	})
	tokenStr, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
