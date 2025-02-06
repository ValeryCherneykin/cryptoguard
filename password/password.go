package password

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-=_+[]{}|;:,.<>?"

func GeneratePassword(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("длина пароля должна быть больше 0")
	}

	password := make([]rune, length)
	charSetLength := big.NewInt(int64(len(charSet)))
	for i := 0; i < length; i++ {
		randIndex, err := rand.Int(rand.Reader, charSetLength)
		if err != nil {
			return "", err
		}
		password[i] = rune(charSet[randIndex.Int64()])
	}

	return string(password), nil
}
