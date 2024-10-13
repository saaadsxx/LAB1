package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция для вычисления факториала
func factorial(n int) {
	time.Sleep(2 * time.Second)
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Printf("Факториал %d равен %d\n", n, result)
}

// Функция для генерации случайных чисел
func randomNumbers() {
	time.Sleep(1 * time.Second)
	for i := 0; i < 5; i++ {
		fmt.Println("Случайное число:", rand.Intn(100))
		time.Sleep(500 * time.Millisecond)
	}
}

// Функция для вычисления суммы числового ряда
func sumSeries(n int) {
	time.Sleep(3 * time.Second)
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("Сумма числового ряда до %d равна %d\n", n, sum)
}

func main() {
	go factorial(5)
	go randomNumbers()
	go sumSeries(10)

	// Дать горутинам время на выполнение
	time.Sleep(5 * time.Second)
	fmt.Println("Все горутины завершены")
}
