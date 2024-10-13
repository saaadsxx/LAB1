package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	counter++
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Финальное значение счётчика:", counter)
}
