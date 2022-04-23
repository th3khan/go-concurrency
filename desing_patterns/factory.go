package main

import (
	"errors"
	"fmt"
)

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getName() string {
	return c.name
}

func (c *Computer) getStock() int {
	return c.stock
}

type Laptop struct {
	Computer
}

func NewLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop",
			stock: 25,
		},
	}
}

type Desktop struct {
	Computer
}

func NewDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop",
			stock: 50,
		},
	}
}

func GetComputerFactory(computerType string) (IProduct, error) {
	if computerType == "laptop" {
		return NewLaptop(), nil
	}

	if computerType == "desktop" {
		return NewDesktop(), nil
	}
	return nil, errors.New("Type not Found...")
}

func PrintNameAndStock(p IProduct) {
	fmt.Printf("Product name: %s, with stock: %d\n", p.getName(), p.getStock())
}

func main() {
	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")

	PrintNameAndStock(laptop)
	PrintNameAndStock(desktop)
}
