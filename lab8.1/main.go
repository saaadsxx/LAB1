package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User представляет модель пользователя
type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Хранение ссылки на базу данных
var db *gorm.DB

// main функция
func main() {
	var err error
	// Подключение к базе данных PostgreSQL
	dsn := "host=localhost user=postgres password=1 dbname=userdb port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Автоматическое создание таблицы
	db.AutoMigrate(&User{})

	// Проверка, если таблица пользователей пуста, и добавление тестовых пользователей
	var userCount int64
	db.Model(&User{}).Count(&userCount)
	if userCount == 0 {
		users := []User{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 35},
		}
		db.Create(&users)
	}

	router := gin.Default()

	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", addUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	router.Run("localhost:8088")
}

// getUsers возвращает список пользователей
func getUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.IndentedJSON(http.StatusOK, users)
}

// getUserByID возвращает пользователя по ID
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	var user User

	if err := db.First(&user, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

// addUser добавляет нового пользователя
func addUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	db.Create(&newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// updateUser обновляет информацию о пользователе
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user User

	if err := db.First(&user, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		return
	}

	db.Save(&user)
	c.IndentedJSON(http.StatusOK, user)
}

// deleteUser удаляет пользователя
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&User{}, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "user deleted"})
}
