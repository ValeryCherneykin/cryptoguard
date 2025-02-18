package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func encrypt(key string) error {
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	cipherText := gcm.Seal(nonce, nonce, []byte(key), nil)
	return nil
}
