package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required"`
}

var db *gorm.DB

func initDB() {
	dsn := "host=localhost user=useradmin password=password dbname=userdb port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&User{})
}

func main() {
	initDB()
	router := gin.Default()

	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", createUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	router.Run(":8080")
}

func getUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&user)
	c.JSON(http.StatusOK, user)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	db.Delete(&user)
	c.JSON(http.StatusNoContent, nil)
}
