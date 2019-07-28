package assistance

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CompareHashAndPassword(hashed []byte, password []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hashed, password); err != nil {
		return false
	}
	return true
}

func GenerateToken(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
