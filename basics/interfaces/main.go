package main

import (
	"fmt"
	"time"
)

type Entry interface {
	Title() string
}

type Book struct {
	Name      string
	Author    string
	Published time.Time
}

func (b Book) Title() string {
	return fmt.Sprintf("%s by %s (%s)", b.Name, b.Author, b.Published.Format("Jan 2006"))
}

type Movie struct {
	Name     string
	Director string
	Year     int
}

func (m Movie) Title() string {
	return fmt.Sprintf("%s (%d)", m.Name, m.Year)
}

func Display(e Entry) string {
	return e.Title()
}

func nilInterfaces() {
	fmt.Println("Nil interfaces")

	var i interface{}
	fmt.Println(i == nil)
	fmt.Printf("%T, %v\n", i, i)
	// interfaces are only nil when both the concrete type
	// and the value are nil.

	var s *string
	fmt.Println("s == nil: ", s == nil)

	i = s
	fmt.Println("i == nil: ", i == nil)
	fmt.Printf("%T, %v\n", i, i)
}

func main() {
	b := Book{Name: "John Adams", Author: "David McCullough", Published: time.Date(2001, time.May, 22, 0, 0, 0, 0, time.UTC)}
	m := Movie{Name: "The Godfather", Director: "Francis Ford Coppola", Year: 1972}
	fmt.Println(Display(b))
	fmt.Println(Display(m))

	// empty interface - can hold a value of any type
	var i interface{}
	i = "asdf"
	fmt.Println(i)

	s, ok := i.(string) // type assertion
	if !ok {
		fmt.Println("s is not type string")
	}
	fmt.Println(s)

	nilInterfaces()

}
