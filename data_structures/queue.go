package main

import (
	"fmt"
	"sync"
)

type Conveior struct {
	length   int
	Products []Product
	lock     sync.Mutex
}

type Product struct {
	Title string
}

func (c *Conveior) Enqueue(p Product) {
	c.Products = append(c.Products, p)
	c.length++
}

func (c *Conveior) Dequeue() *Product {
	prod := c.Products[0]
	c.Products = c.Products[1:c.length]
	c.length--
	return &prod
}

func (c *Conveior) isEmpty() bool {
	return c.length == 0
}

func (c *Conveior) Front() Product {
	return c.Products[0]
}

func main() {
	conveior := Conveior{}

	prod1 := Product{"Prod 1"}
	conveior.Enqueue(prod1)
	prod2 := Product{"Prod 2"}
	conveior.Enqueue(prod2)
	prod3 := Product{"Prod 3"}
	conveior.Enqueue(prod3)

	fmt.Println(conveior)
	fmt.Println(conveior.Front())
	conveior.Dequeue()
	fmt.Println(conveior)
	fmt.Println(conveior.Front())
}
