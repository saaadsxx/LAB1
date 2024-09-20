package main

import (
	"fmt"
	"strings"
)

// Функция для вычисления среднего возраста всех людей в карте
func averageAge(people map[string]int) float64 {
	totalAge := 0
	for _, age := range people {
		totalAge += age
	}
	return float64(totalAge) / float64(len(people))
}

func main() {
	// 1. Создаем карту с именами людей и их возрастами
	people := map[string]int{
		"John":  25,
		"Anna":  30,
		"Mike":  40,
		"Susan": 35,
	}

	// Добавляем нового человека
	people["Tom"] = 50

	// Выводим все записи на экран
	fmt.Println("Карта людей и их возрастов:")
	for name, age := range people {
		fmt.Printf("%s: %d лет\n", name, age)
	}

	// 2. Вычисляем и выводим средний возраст
	avgAge := averageAge(people)
	fmt.Printf("\nСредний возраст: %.2f лет\n", avgAge)

	// 3. Удаление записи по имени
	var nameToDelete string
	fmt.Print("\nВведите имя для удаления: ")
	fmt.Scan(&nameToDelete)

	if _, exists := people[nameToDelete]; exists {
		delete(people, nameToDelete)
		fmt.Printf("Запись о %s удалена.\n", nameToDelete)
	} else {
		fmt.Println("Такого имени нет в карте.")
	}

	// Выводим обновленную карту
	fmt.Println("\nОбновленная карта людей и их возрастов:")
	for name, age := range people {
		fmt.Printf("%s: %d лет\n", name, age)
	}

	// 4. Считывание строки с ввода и вывод в верхнем регистре
	var inputString string
	fmt.Print("\nВведите строку: ")
	fmt.Scan(&inputString)

	upperString := strings.ToUpper(inputString)
	fmt.Printf("Строка в верхнем регистре: %s\n", upperString)

	// 5. Считывание нескольких чисел и вычисление их суммы
	var n int
	fmt.Print("\nСколько чисел вы хотите ввести? ")
	fmt.Scan(&n)

	numbers := make([]int, n)
	sum := 0
	fmt.Println("Введите числа:")
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
		sum += numbers[i]
	}
	fmt.Printf("Сумма введенных чисел: %d\n", sum)

	// 6. Считывание массива целых чисел и вывод в обратном порядке
	fmt.Println("\nМассив чисел в обратном порядке:")
	for i := len(numbers) - 1; i >= 0; i-- {
		fmt.Printf("%d ", numbers[i])
	}
	fmt.Println()
}
