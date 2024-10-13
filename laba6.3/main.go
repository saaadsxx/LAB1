package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Функция для генерации случайных чисел
func generateRandomNumbers(numChannel chan int) {
	for {
		num := rand.Intn(100)
		numChannel <- num // Отправка числа в канал
		time.Sleep(time.Second)
	}
}

// Функция для проверки чётности/нечётности
func checkEvenOdd(numChannel chan int, resultChannel chan string) {
	for num := range numChannel {
		if num%2 == 0 {
			resultChannel <- fmt.Sprintf("%d - чётное", num)
		} else {
			resultChannel <- fmt.Sprintf("%d - нечётное", num)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numChannel := make(chan int)
	resultChannel := make(chan string)

	go generateRandomNumbers(numChannel)
	go checkEvenOdd(numChannel, resultChannel)

	// Основной цикл, который использует select для получения данных из resultChannel
	for {
		select {
		case result := <-resultChannel: // Получение результата из канала
			fmt.Println(result)
		case <-time.After(2 * time.Second):
			close(numChannel)
			close(resultChannel)
			return
		}
	}
}
