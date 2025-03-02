package util

import (
	// "mp1/db"

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
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iss":  "mp1-auth",
		"sub":  name,
		"hash": hash,
	})
	tokenStr, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
