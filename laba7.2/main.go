package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Print("Enter message: ")
	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')

	conn.Write([]byte(message))
	reply, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Printf("Server reply: %s", reply)
}
