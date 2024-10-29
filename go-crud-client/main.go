package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const baseURL = "http://127.0.0.1:8088/api/users" // Убедитесь, что URL совпадает с вашим сервером

var sessionToken string

// Обновленная структура User
type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,gte=0,lte=130"` // Минимум 0, максимум 130
}

// Функция для авторизации (замените на вашу логику)
func login() error {
	// Здесь можно реализовать вашу логику авторизации, если необходимо
	// Например, просто сохраняем токен статически для тестирования
	sessionToken = "example_token"
	return nil
}

// Функция для добавления пользователя
func addUser(name string, age int) error {
	user := User{Name: name, Age: age}
	data, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+sessionToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to add user: %s", resp.Status)
	}

	fmt.Println("User added successfully.")
	return nil
}

// Функция для получения пользователей
func getUsers() {
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+sessionToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to get users:", resp.Status)
		return
	}

	var users []User
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &users)

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}
}

// Функция для обновления пользователя
func updateUser(id uint, name string, age int) error {
	user := User{Name: name, Age: age}
	data, _ := json.Marshal(user)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%d", baseURL, id), bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+sessionToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update user: %s", resp.Status)
	}

	fmt.Println("User updated successfully.")
	return nil
}

// Функция для удаления пользователя
func deleteUser(id uint) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%d", baseURL, id), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+sessionToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete user: %s", resp.Status)
	}

	fmt.Println("User deleted successfully.")
	return nil
}

// Функция для отображения меню
func showMenu() {
	fmt.Println("1. Login")
	fmt.Println("2. Add User")
	fmt.Println("3. Get Users")
	fmt.Println("4. Update User")
	fmt.Println("5. Delete User")
	fmt.Println("6. Exit")
}

func main() {
	var choice int

	for {
		showMenu()
		fmt.Print("Choose an option: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input:", err)
			continue
		}

		switch choice {
		case 1:
			if err := login(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Logged in successfully.")
			}

		case 2:
			var name string
			var age int
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			fmt.Print("Enter age: ")
			fmt.Scan(&age)

			if err := addUser(name, age); err != nil {
				fmt.Println(err)
			}

		case 3:
			getUsers()

		case 4:
			var id uint
			var name string
			var age int
			fmt.Print("Enter user ID to update: ")
			fmt.Scan(&id)
			fmt.Print("Enter new name: ")
			fmt.Scan(&name)
			fmt.Print("Enter new age: ")
			fmt.Scan(&age)

			if err := updateUser(id, name, age); err != nil {
				fmt.Println(err)
			}

		case 5:
			var id uint
			fmt.Print("Enter user ID to delete: ")
			fmt.Scan(&id)

			if err := deleteUser(id); err != nil {
				fmt.Println(err)
			}

		case 6:
			os.Exit(0)

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
