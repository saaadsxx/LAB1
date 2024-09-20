package main

import (
	"fmt"
)

// 1. Определение четного или нечетного числа
func checkEvenOdd(num int) {
	if num%2 == 0 {
		fmt.Println("Четное число")
	} else {
		fmt.Println("Нечетное число")
	}
}

// 2. Определение, является ли число положительным, отрицательным или нулем
func checkNumber(num int) string {
	if num > 0 {
		return "Positive"
	} else if num < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}

// 3. Вывод всех чисел от 1 до 10 с помощью цикла for
func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

// 4. Определение длины строки
func stringLength(s string) int {
	return len(s)
}

// 5. Структура Rectangle и метод для вычисления площади
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// 6. Функция для вычисления среднего значения двух целых чисел
func average(a, b int) float64 {
	return float64(a+b) / 2
}

func main() {
	// Ввод для четности и нечетности
	var num int
	fmt.Print("Введите число для проверки четности/нечетности: ")
	fmt.Scan(&num)
	checkEvenOdd(num)

	// Ввод для проверки положительное, отрицательное или ноль
	fmt.Print("Введите число для проверки положительное, отрицательное или ноль: ")
	fmt.Scan(&num)
	fmt.Println(checkNumber(num))

	// Вывод чисел от 1 до 10
	fmt.Println("Числа от 1 до 10:")
	printNumbers()

	// Ввод строки и вывод её длины
	var inputString string
	fmt.Print("Введите строку для определения её длины: ")
	fmt.Scan(&inputString)
	fmt.Printf("Длина строки: %d\n", stringLength(inputString))

	// Пример с прямоугольником
	r := Rectangle{Width: 5, Height: 3}
	fmt.Printf("Площадь прямоугольника (5x3): %.2f\n", r.Area())

	// Ввод двух чисел и вычисление их среднего значения
	var num1, num2 int
	fmt.Print("Введите первое число для среднего значения: ")
	fmt.Scan(&num1)
	fmt.Print("Введите второе число для среднего значения: ")
	fmt.Scan(&num2)
	fmt.Printf("Среднее значение: %.2f\n", average(num1, num2))
}
