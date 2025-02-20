package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func EncryptFile(key []byte, inputPath, outputPath string) {

	content, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Ошибка создания шифра:", err)
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Ошибка создания GCM:", err)
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("Ошибка генерации nonce:", err)
		return
	}

	encrypted := gcm.Seal(nonce, nonce, content, nil)

	err = os.WriteFile(outputPath, encrypted, 0777)
	if err != nil {
		fmt.Println("Ошибка записи зашифрованного файла:", err)
	}
}
