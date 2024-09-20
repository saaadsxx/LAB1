package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Вывод текущего времени и даты
	currentTime := time.Now()
	fmt.Println("Текущая дата и время:", currentTime)

	// 2. Создание и вывод переменных различных типов
	a := 52
	b := 99.99
	c := "aboba"
	d := false

	fmt.Println("Int:", a)
	fmt.Println("Float64:", b)
	fmt.Println("String:", c)
	fmt.Println("Bool:", d)

	// 4. Арифметические операции с двумя целыми числами
	num1, num2 := 15, 4
	fmt.Println("Сумма:", num1+num2)
	fmt.Println("Разность:", num1-num2)
	fmt.Println("Произведение:", num1*num2)
	fmt.Println("Частное:", num1/num2)

	// 5. Функция для вычисления суммы и разности двух чисел с плавающей запятой
	floatSum, floatDiff := floatOperations(10.5, 3.2)
	fmt.Println("Сумма с плавающей запятой:", floatSum)
	fmt.Println("Разность с плавающей запятой:", floatDiff)

	// 6. Вычисление среднего значения трех чисел
	numA, numB, numC := 10.0, 20.0, 30.0
	average := (numA + numB + numC) / 3
	fmt.Println("Среднее значение трех чисел:", average)
}

// Функция для вычисления суммы и разности двух чисел с плавающей запятой
func floatOperations(a, b float64) (float64, float64) {
	return a + b, a - b
}
