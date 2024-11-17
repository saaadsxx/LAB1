package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// Генерация пары ключей RSA и сохранение в файлы
func generateKeys() (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Ошибка генерации ключей:", err)
		os.Exit(1)
	}

	// Сохранение закрытого ключа
	privateFile, err := os.Create("private.pem")
	if err != nil {
		fmt.Println("Ошибка создания файла для закрытого ключа:", err)
		os.Exit(1)
	}
	defer privateFile.Close()

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	pem.Encode(privateFile, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKeyBytes})

	// Сохранение открытого ключа
	publicKey := &privateKey.PublicKey
	publicFile, err := os.Create("public.pem")
	if err != nil {
		fmt.Println("Ошибка создания файла для открытого ключа:", err)
		os.Exit(1)
	}
	defer publicFile.Close()

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println("Ошибка маршалинга открытого ключа:", err)
		os.Exit(1)
	}
	pem.Encode(publicFile, &pem.Block{Type: "PUBLIC KEY", Bytes: publicKeyBytes})

	return privateKey, publicKey
}

// Подпись сообщения с использованием закрытого ключа
func signMessage(privateKey *rsa.PrivateKey, message string) ([]byte, error) {
	hashed := sha256.Sum256([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, err
	}
	return signature, nil
}

// Проверка подписи с использованием открытого ключа
func verifySignature(publicKey *rsa.PublicKey, message string, signature []byte) error {
	hashed := sha256.Sum256([]byte(message))
	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
}

func main() {
	// Генерация ключей
	privateKey, publicKey := generateKeys()

	// Подпись сообщения
	message := "Hello, this is a signed message"
	signature, err := signMessage(privateKey, message)
	if err != nil {
		fmt.Println("Ошибка подписи:", err)
		return
	}
	fmt.Printf("Подпись: %x\n", signature)

	// Проверка подписи
	err = verifySignature(publicKey, message, signature)
	if err != nil {
		fmt.Println("Подпись недействительна")
	} else {
		fmt.Println("Подпись действительна")
	}
}
