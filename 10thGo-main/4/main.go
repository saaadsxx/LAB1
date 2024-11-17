package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Server представляет собой TCP-сервер, который управляет соединениями и их завершением.
type Server struct {
	listener net.Listener   // Слушатель для принятия входящих соединений
	wg       sync.WaitGroup // Группа ожидания для отслеживания горутин
}

// NewServer создает новый сервер, который слушает указанный порт.
func NewServer(port string) (*Server, error) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, err
	}
	return &Server{listener: listener}, nil
}

// Start запускает сервер и ожидает входящих соединений.
func (s *Server) Start() {
	fmt.Println("Сервер запущен, ожидает подключения...")
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Ошибка при принятии соединения:", err)
			continue
		}
		s.wg.Add(1)
		go s.handleConnection(conn)
	}
}

// handleConnection обрабатывает сообщение от клиента.
func (s *Server) handleConnection(conn net.Conn) {
	defer s.wg.Done()
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Ошибка при чтении:", err)
			return
		}
		message := string(buffer[:n])
		fmt.Printf("Получено сообщение: %s", message)

		_, err = conn.Write([]byte("Сообщение получено\n"))
		if err != nil {
			fmt.Println("Ошибка при отправке:", err)
			return
		}
	}
}

// Shutdown корректно завершает работу сервера и закрывает все соединения.
func (s *Server) Shutdown() {
	fmt.Println("Завершение работы сервера...")
	s.listener.Close()
	s.wg.Wait()
	fmt.Println("Все соединения закрыты.")
}

// main - точка входа в программу.
func main() {
	mode := flag.String("mode", "", "Укажите 'server' или 'client'")
	flag.Parse()

	if *mode == "server" {
		server, err := NewServer("8080")
		if err != nil {
			fmt.Println("Не удалось создать сервер:", err)
			return
		}

		// Обработка сигнала завершения
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-signalChan
			server.Shutdown()
			os.Exit(0)
		}()

		server.Start()
	} else if *mode == "client" {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			fmt.Println("Ошибка подключения к серверу:", err)
			return
		}
		defer conn.Close()

		for {
			var message string
			fmt.Print("Введите сообщение (или 'exit' для выхода): ")
			_, err := fmt.Scanln(&message)
			if err != nil || message == "exit" {
				break
			}

			_, err = conn.Write([]byte(message))
			if err != nil {
				fmt.Println("Ошибка при отправке сообщения:", err)
				return
			}

			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Ошибка при чтении ответа:", err)
				return
			}
			fmt.Printf("Ответ сервера: %s", string(buffer[:n]))
		}
	} else {
		fmt.Println("Необходимо указать режим: -mode=server или -mode=client")
		os.Exit(1)
	}
}

// openssl genpkey -algorithm RSA -out ca.key Создать корневой (CA) сертификат
// openssl req -x509 -new -key ca.key -out ca.crt -days 365 -subj "//CN=My CA"

// openssl genpkey -algorithm RSA -out server.key Создать серверный ключ и запрос на подпись (CSR)
// openssl req -new -key server.key -out server.csr -subj "//CN=localhost"

// openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365 Подписать серверный сертификат корневым сертификатом (CA):

// openssl genpkey -algorithm RSA -out client.keyopenssl req -new -key client.key -out client.csr -subj "/CN=client" Создать клиентский ключ и запрос на подпись (CSR)
// openssl req -new -key client.key -out client.csr -subj "//CN=client"

// openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365 Подписать клиентский сертификат корневым сертификатом (CA)

// go run main.go -mode=server

// go run main.go -mode=client
