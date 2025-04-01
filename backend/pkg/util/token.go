package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	JWTIssuer = "mp1-auth"
	fiveWeeks = time.Hour * 24 * 7 * 5
)

// Generates a JWT token with name and password in Claims.
// Also generates a 16 byte large salt.
func GenerateJWT(name string, password []byte, key []byte) (string, error) {
	salt, err := RandomBytes(16)
	if err != nil {
		return "", err
	}
	hash, err := GenerateHash(password, salt)
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iss":  JWTIssuer,
		"iat":  jwt.NewNumericDate(now),
		"exp":  jwt.NewNumericDate(now.Add(fiveWeeks)),
		"sub":  name,
		"hash": hash,
	})
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

// Validates a JWT based on current time.
func ValidateJWT(token string, key []byte) bool {
	parsed, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	},
		jwt.WithIssuer(JWTIssuer),
		jwt.WithIssuedAt(),
		jwt.WithExpirationRequired(),
	)
	return parsed.Valid
}
