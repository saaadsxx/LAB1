package main

import (
	"fmt"
	"math"
)

// Структура Person
type Person struct {
	name string
	age  int
}

// Метод для вывода информации о человеке
func (p Person) PrintInfo() {
	fmt.Printf("Имя: %s, Возраст: %d\n", p.name, p.age)
}

// Метод birthday для увеличения возраста на 1 год
func (p *Person) Birthday() {
	p.age++
}

// Структура Circle
type Circle struct {
	radius float64
}

// Метод для вычисления площади круга
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Структура Rectangle
type Rectangle struct {
	width, height float64
}

// Метод для вычисления площади прямоугольника
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Интерфейс Shape с методом Area
type Shape interface {
	Area() float64
}

// Функция для вывода площади каждого объекта Shape
func PrintAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("Площадь: %.2f\n", shape.Area())
	}
}

// Структура Book с полями title и author
type Book struct {
	title  string
	author string
}

// Интерфейс Stringer
type Stringer interface {
	String() string
}

// Реализация интерфейса Stringer для структуры Book
func (b Book) String() string {
	return fmt.Sprintf("Книга: %s, Автор: %s", b.title, b.author)
}

func main() {
	// Работа со структурой Person
	person := Person{name: "Alice", age: 30}
	person.PrintInfo()
	person.Birthday() // Увеличение возраста на 1 год
	person.PrintInfo()

	// Работа с интерфейсом Shape
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 4, height: 6}
	shapes := []Shape{circle, rectangle}

	// Вывод площадей объектов
	fmt.Println("\nПлощади объектов:")
	PrintAreas(shapes)

	// Работа с интерфейсом Stringer
	book := Book{title: "Преступление и наказание", author: "Фёдор Достоевский"}
	fmt.Println("\nИнформация о книге:")
	fmt.Println(book.String())
}
