package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"os"
)

// Функция для хэширования строки
func hashData(input string, algo string) string {
	var h hash.Hash
	switch algo {
	case "md5":
		h = md5.New()
	case "sha256":
		h = sha256.New()
	case "sha512":
		h = sha512.New()
	default:
		fmt.Println("Неподдерживаемый алгоритм хэширования")
		os.Exit(1)
	}
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}

// Функция проверки целостности
func verifyHash(input, expectedHash, algo string) bool {
	calculatedHash := hashData(input, algo)
	return calculatedHash == expectedHash
}

func main() {
	var input, algo, hash string
	fmt.Println("Введите строку для хэширования:")
	fmt.Scanln(&input)
	fmt.Println("Выберите алгоритм (md5, sha256, sha512):")
	fmt.Scanln(&algo)

	// Хэширование
	hashedData := hashData(input, algo)
	fmt.Printf("Хэш: %s\n", hashedData)

	// Проверка целостности
	fmt.Println("Введите строку для проверки хэша:")
	fmt.Scanln(&input)
	fmt.Println("Введите ожидаемый хэш:")
	fmt.Scanln(&hash)

	if verifyHash(input, hash, algo) {
		fmt.Println("Хэш совпадает")
	} else {
		fmt.Println("Хэш не совпадает")
	}
}
