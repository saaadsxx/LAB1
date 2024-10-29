package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Инициализация тестовой базы данных
func setupTestDB() {
	dsn := "host=localhost user=postgres password=1 dbname=test_userdb port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}
	db.AutoMigrate(&User{})
}

// Тестовый экземпляр Gin
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", addUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)
	return router
}

func TestGetUsers(t *testing.T) {
	setupTestDB()
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddUser(t *testing.T) {
	setupTestDB()
	router := setupRouter()

	user := User{Name: "Test User", Age: 28}
	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var responseUser User
	json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.Equal(t, user.Name, responseUser.Name)
	assert.Equal(t, user.Age, responseUser.Age)
}

func TestGetUserByID(t *testing.T) {
	setupTestDB()
	router := setupRouter()

	// Добавляем тестового пользователя
	user := User{Name: "Test User", Age: 28}
	db.Create(&user)

	// Получаем ID пользователя и выполняем запрос
	req, _ := http.NewRequest("GET", "/users/"+strconv.Itoa(int(user.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var responseUser User
	json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.Equal(t, user.Name, responseUser.Name)
	assert.Equal(t, user.Age, responseUser.Age)
}

func TestDeleteUser(t *testing.T) {
	setupTestDB()
	router := setupRouter()

	// Добавляем тестового пользователя
	user := User{Name: "Test User", Age: 28}
	db.Create(&user)

	req, _ := http.NewRequest("DELETE", "/users/"+strconv.Itoa(int(user.ID)), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
