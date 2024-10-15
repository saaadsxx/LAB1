package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

func handleConnection(conn net.Conn) {
	defer wg.Done()
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			return
		}

		message := string(buffer[:n])
		fmt.Println("Получено сообщение:", message)

		// Отправка подтверждения
		_, err = conn.Write([]byte("Сообщение получено"))
		if err != nil {
			fmt.Println("Ошибка отправки:", err)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка создания TCP-сервера:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("TCP-сервер запущен на порту 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при подключении:", err)
			continue
		}
		wg.Add(1)
		go handleConnection(conn)
	}
}
