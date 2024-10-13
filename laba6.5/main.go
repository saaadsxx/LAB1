package main

import (
	"fmt"
)

type Request struct {
	a, b   int
	op     string
	result chan int
}

func calculator(requests chan Request) {
	for req := range requests {
		switch req.op {
		case "+":
			req.result <- req.a + req.b
		case "-":
			req.result <- req.a - req.b
		case "*":
			req.result <- req.a * req.b
		case "/":
			if req.b != 0 {
				req.result <- req.a / req.b
			} else {
				req.result <- 0
			}
		default:
			req.result <- 0
		}
	}
}

func main() {
	requests := make(chan Request)
	go calculator(requests)

	result := make(chan int)
	requests <- Request{a: 5, b: 3, op: "*", result: result}
	fmt.Println("Результат:", <-result)

	close(requests)
}
