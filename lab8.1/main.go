package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User представляет модель пользователя
type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,gte=0,lte=130"` // Минимум 0, максимум 130
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

// handleError централизованная обработка ошибок
func handleError(c *gin.Context, err error, statusCode int) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

// getUsers возвращает список пользователей с пагинацией и фильтрацией
func getUsers(c *gin.Context) {
	var users []User
	var count int64

	// Параметры запроса
	name := c.Query("name")
	age := c.Query("age")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Применяем фильтрацию
	query := db.Model(&User{})
	if name != "" {
		query = query.Where("name ILIKE ?", "%"+name+"%")
	}
	if age != "" {
		query = query.Where("age = ?", age)
	}

	// Подсчет общего числа пользователей
	query.Count(&count)

	// Пагинация
	offset := (page - 1) * limit
	query.Offset(offset).Limit(limit).Find(&users)

	c.JSON(http.StatusOK, gin.H{"total": count, "users": users})
}

// getUserByID возвращает пользователя по ID
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	var user User

	if err := db.First(&user, id).Error; err != nil {
		handleError(c, err, http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, user)
}

// addUser добавляет нового пользователя
func addUser(c *gin.Context) {
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		handleError(c, err, http.StatusBadRequest)
		return
	}

	db.Create(&newUser)
	c.JSON(http.StatusCreated, newUser)
}

// updateUser обновляет информацию о пользователе
func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user User

	if err := db.First(&user, id).Error; err != nil {
		handleError(c, err, http.StatusNotFound)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		handleError(c, err, http.StatusBadRequest)
		return
	}

	db.Save(&user)
	c.JSON(http.StatusOK, user)
}

// deleteUser удаляет пользователя
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&User{}, id).Error; err != nil {
		handleError(c, err, http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
