package main

import (
	"fmt"
	"laba3/mathutils"
	"laba3/stringutils"
)

func main() {
	// 1. Ввод числа и вычисление его факториала
	var num int
	fmt.Print("Введите число для вычисления факториала: ")
	fmt.Scan(&num)
	fmt.Printf("Факториал числа %d: %d\n", num, mathutils.Factorial(num))

	// 2. Ввод строки и её переворот
	var input string
	fmt.Print("Введите строку для переворота: ")
	fmt.Scan(&input)
	fmt.Printf("Перевернутая строка: %s\n", stringutils.Reverse(input))

	// 3. Создание и вывод массива из 5 целых чисел
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Массив:", arr)

	// 4. Создание среза из массива и операции добавления/удаления элементов
	slice := arr[:]
	fmt.Println("Исходный срез:", slice)

	// Добавляем элементы
	slice = append(slice, 6, 7)
	fmt.Println("Срез после добавления элементов:", slice)

	// Удаление элемента (удаляем второй элемент)
	slice = append(slice[:1], slice[2:]...)
	fmt.Println("Срез после удаления элемента:", slice)

	// 5. Создание среза из строк и нахождение самой длинной строки
	strings := []string{"Арбуз", "Машина", "мышь", "АВТОМОБИЛЬ!"}
	longest := findLongestString(strings)
	fmt.Printf("Самая длинная строка: %s\n", longest)
}

// Функция для нахождения самой длинной строки
func findLongestString(strings []string) string {
	longest := ""
	for _, s := range strings {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}
