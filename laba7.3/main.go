package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg sync.WaitGroup

func handleConnection(conn net.Conn) {
	defer wg.Done()
	defer conn.Close()
	reader := bufio.NewReader(conn)
	message, _ := reader.ReadString('\n')
	fmt.Printf("Message received: %s", message)
	conn.Write([]byte("Message received\n"))
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP Server is listening on port 8080...")

	// Channel to handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting connection:", err)
				continue
			}
			wg.Add(1)
			go handleConnection(conn)
		}
	}()

	<-stop
	fmt.Println("\nShutting down server...")
	listener.Close()
	wg.Wait()
	fmt.Println("Server stopped gracefully.")
}
