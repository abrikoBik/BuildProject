package util

import (
	// "mp1/db"

	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTIssuer string = "mp1-auth"

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
		"iss":  JWTIssuer,
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

// Validates a JWT based on current time.
func ValidateJWT(token string, key string) (bool, error) {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	},
		jwt.WithIssuer(JWTIssuer),
		jwt.WithIssuedAt(),
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return false, err
	}
	return parsed.Valid, nil
}
