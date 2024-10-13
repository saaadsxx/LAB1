package main

import (
	"fmt"
)

// Функция для генерации первых n чисел Фибоначчи
func generateFibonacci(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a       // Отправляем текущее число Фибоначчи в канал
		a, b = b, a+b // Обновляем значения a и b
	}
	close(ch) // Закрываем канал после завершения
}

func main() {
	ch := make(chan int)

	go generateFibonacci(10, ch)

	for num := range ch { // Чтение из канала до его закрытия
		fmt.Println(num)
	}

	fmt.Println("Все числа Фибоначчи сгенерированы и выведены")
}
