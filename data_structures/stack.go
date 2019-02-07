package main

import "fmt"

type Stack struct {
	Books []Book
	count int
}

type Book struct {
	Title string
}

func (s *Stack) Push(book Book) {
	s.Books = append(s.Books, book)
	s.count++
}

func (s *Stack) Pop() {
	l := len(s.Books)
	s.Books = s.Books[:l-1]
	s.count--
}

func (s *Stack) Pip() {
	fmt.Println(s)
}

func main() {
	stack := Stack{}
	book1 := Book{"Algebra"}
	book2 := Book{"Literature"}
	book3 := Book{"Physic"}
	stack.Push(book1)
	stack.Push(book2)
	stack.Push(book3)

	stack.Pip()
	stack.Pop()
	stack.Pip()
}
