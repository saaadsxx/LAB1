package main

import (
	"fmt"
	"sync"
)

// Воркер для реверсирования строки
func worker(id int, tasks <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		reversed := reverseString(task)
		fmt.Printf("Воркер %d обработал задачу: %s -> %s\n", id, task, reversed)
		results <- reversed
	}
}

// Функция реверса строки
func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	tasks := make(chan string, 10)
	results := make(chan string, 10)
	var wg sync.WaitGroup

	// Создаём пул из 3 воркеров
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Отправляем задачи воркерам
	stringsToReverse := []string{"apple", "banana", "cherry", "date", "fig"}
	for _, s := range stringsToReverse {
		tasks <- s
	}

	close(tasks)
	wg.Wait()
	close(results)

	// Выводим результаты
	for res := range results {
		fmt.Println("Результат:", res)
	}
}
