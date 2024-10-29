package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User представляет модель пользователя
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required,gte=0,lte=130"` // Минимум 0, максимум 130
	Password string `json:"password" binding:"required"`          // Удалено "json:"-""
}

// LoginRequest представляет запрос для входа пользователя
type LoginRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Session представляет модель сессии
type Session struct {
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

// Хранение ссылки на базу данных и сессий
var db *gorm.DB
var wg sync.WaitGroup
var sessions = make(map[string]Session)

func main() {
	wg.Add(1)
	go startServer() // Запуск сервера в отдельной горутине

	clientMenu() // Запуск клиентского интерфейса
	wg.Wait()    // Ожидание завершения всех горутин
}

// startServer запускает HTTP сервер
func startServer() {
	defer wg.Done() // Уменьшаем счетчик горутин при завершении

	var err error
	// Подключение к базе данных PostgreSQL
	dsn := "host=localhost user=postgres password=1 dbname=userdb port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Автоматическое создание таблицы
	db.AutoMigrate(&User{})

	router := gin.Default()

	router.POST("/register", registerUser)
	router.POST("/login", loginUser)

	authorized := router.Group("/")
	authorized.Use(AuthMiddleware())
	{
		authorized.GET("/users", getUsers)
		authorized.GET("/users/:id", getUserByID)
		authorized.POST("/users", addUser)
		authorized.PUT("/users/:id", updateUser)
		authorized.DELETE("/users/:id", deleteUser)
	}

	// Запуск сервера
	router.Run("localhost:8088")
}

// registerUser позволяет пользователю зарегистрироваться
func registerUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// loginUser позволяет пользователю войти
func loginUser(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dbUser User
	if err := db.Where("name = ?", loginReq.Name).First(&dbUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверное имя пользователя или пароль"})
		return
	}

	// Проверка пароля (поскольку теперь пароли хранятся в открытом виде, эта проверка не нужна)
	if dbUser.Password != loginReq.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверное имя пользователя или пароль"})
		return
	}

	// Генерация токена
	token := fmt.Sprintf("%d", dbUser.ID)
	sessions[token] = Session{UserID: dbUser.ID, Token: token}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// AuthMiddleware проверяет токен сессии
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if session, exists := sessions[token]; exists {
			c.Set("userID", session.UserID) // Сохраняем UserID для дальнейшего использования
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Необходима авторизация"})
			c.Abort()
		}
	}
}

// clientMenu отображает меню для пользователя
func clientMenu() {
	for {
		fmt.Println("\nВыберите операцию:")
		fmt.Println("1. Зарегистрироваться")
		fmt.Println("2. Войти")
		fmt.Println("3. Получить всех пользователей")
		fmt.Println("4. Получить пользователя по ID")
		fmt.Println("5. Добавить пользователя")
		fmt.Println("6. Обновить пользователя")
		fmt.Println("7. Удалить пользователя")
		fmt.Println("8. Выйти")

		var choice int
		fmt.Print("Введите номер операции: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			registerUserClient()
		case 2:
			loginUserClient()
		case 3:
			getUsersClient()
		case 4:
			getUserByIDClient()
		case 5:
			addUserClient()
		case 6:
			updateUserClient()
		case 7:
			deleteUserClient()
		case 8:
			fmt.Println("Выход из программы...")
			os.Exit(0)
		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}

// registerUserClient позволяет пользователю зарегистрироваться через клиент
func registerUserClient() {
	var user User
	fmt.Print("Введите имя пользователя: ")
	fmt.Scan(&user.Name)
	fmt.Print("Введите возраст пользователя: ")
	fmt.Scan(&user.Age)
	fmt.Print("Введите пароль: ")
	fmt.Scan(&user.Password)

	jsonData, _ := json.Marshal(user)
	resp, err := http.Post("http://localhost:8088/register", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Ошибка:", string(body))
		return
	}

	fmt.Println("Пользователь зарегистрирован:", string(body))
}

// loginUserClient позволяет пользователю войти через клиент
func loginUserClient() {
	var loginReq LoginRequest
	fmt.Print("Введите имя пользователя: ")
	fmt.Scan(&loginReq.Name)
	fmt.Print("Введите пароль: ")
	fmt.Scan(&loginReq.Password)

	jsonData, _ := json.Marshal(loginReq)
	resp, err := http.Post("http://localhost:8088/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка:", string(body))
		return
	}

	var response map[string]string
	json.Unmarshal(body, &response)
	fmt.Println("Токен:", response["token"])
}

// getUsersClient получает всех пользователей через клиент
func getUsersClient() {
	var token string
	fmt.Print("Введите токен: ")
	fmt.Scan(&token)

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:8088/users", nil)
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка:", string(body))
		return
	}

	var users []User
	json.Unmarshal(body, &users)

	// Вывод каждого пользователя на отдельной строке
	fmt.Println("Пользователи:")
	for _, user := range users {
		fmt.Printf("ID: %d, Имя: %s, Возраст: %d\n", user.ID, user.Name, user.Age)
	}
}

// getUserByIDClient получает пользователя по ID через клиент
func getUserByIDClient() {
	var id uint
	fmt.Print("Введите ID пользователя: ")
	fmt.Scan(&id)

	var token string
	fmt.Print("Введите токен: ")
	fmt.Scan(&token)

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://localhost:8088/users/%d", id), nil)
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка:", string(body))
		return
	}

	fmt.Println("Пользователь:", string(body))
}

// addUserClient добавляет пользователя через клиент
func addUserClient() {
	var user User
	fmt.Print("Введите имя пользователя: ")
	fmt.Scan(&user.Name)
	fmt.Print("Введите возраст пользователя: ")
	fmt.Scan(&user.Age)
	fmt.Print("Введите пароль: ")
	fmt.Scan(&user.Password)

	jsonData, _ := json.Marshal(user)
	resp, err := http.Post("http://localhost:8088/users", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		fmt.Println("Ошибка:", string(body))
		return
	}

	fmt.Println("Пользователь добавлен:", string(body))
}

// updateUserClient обновляет пользователя через клиент
func updateUserClient() {
	var id uint
	fmt.Print("Введите ID пользователя: ")
	fmt.Scan(&id)

	var user User
	fmt.Print("Введите новое имя пользователя: ")
	fmt.Scan(&user.Name)
	fmt.Print("Введите новый возраст пользователя: ")
	fmt.Scan(&user.Age)
	fmt.Print("Введите новый пароль: ")
	fmt.Scan(&user.Password)

	jsonData, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8088/users/%d", id), bytes.NewBuffer(jsonData))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка:", string(body))
		return
	}

	fmt.Println("Пользователь обновлен:", string(body))
}

// deleteUserClient удаляет пользователя через клиент
// deleteUserClient удаляет пользователя через клиент
func deleteUserClient() {
	var id int
	fmt.Print("Введите ID пользователя для удаления: ")
	fmt.Scan(&id)

	var token string
	fmt.Print("Введите токен: ")
	fmt.Scan(&token)

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8088/users/%d", id), nil)
	req.Header.Set("Authorization", token) // Устанавливаем токен в заголовок

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Ошибка:", string(body))
		return
	}

	fmt.Println("Пользователь удален.")
}

// getUserByID получаем пользователя по ID
func getUserByID(c *gin.Context) {
	var user User
	id := c.Param("id")
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// getUsers получает всех пользователей
func getUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}

// addUser добавляет нового пользователя
func addUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// updateUser обновляет информацию о пользователе
func updateUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	if err := db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&user)
	c.JSON(http.StatusOK, user)
}

// deleteUser удаляет пользователя
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&User{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}
	c.Status(http.StatusNoContent)
}
