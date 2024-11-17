package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// Шифрование AES
func encryptAES(plaintext, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return hex.EncodeToString(ciphertext), nil
}

// Расшифровка AES
func decryptAES(ciphertextHex, key string) (string, error) {
	ciphertext, _ := hex.DecodeString(ciphertextHex)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}

func main() {
	var text, key string

	fmt.Println("Введите строку для шифрования:")
	fmt.Scanln(&text)
	fmt.Println("Введите ключ (длина 16 символов):")
	fmt.Scanln(&key)

	encrypted, err := encryptAES(text, key)
	if err != nil {
		fmt.Println("Ошибка при шифровании:", err)
		return
	}

	fmt.Println("Зашифрованные данные:", encrypted)

	decrypted, err := decryptAES(encrypted, key)
	if err != nil {
		fmt.Println("Ошибка при расшифровке:", err)
		return
	}

	fmt.Println("Расшифрованные данные:", decrypted)
}
