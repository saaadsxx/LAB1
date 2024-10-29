package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User структура для пользователей
type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users []User
var nextID uint = 1

// handler для получения всех пользователей
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// handler для добавления нового пользователя
func addUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = nextID
	nextID++
	users = append(users, user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// main функция
func main() {
	http.HandleFunc("/api/users", getUsers)
	http.HandleFunc("/api/users", addUser)

	fmt.Println("Server is running on port 8088...")
	http.ListenAndServe(":8088", nil)
}
