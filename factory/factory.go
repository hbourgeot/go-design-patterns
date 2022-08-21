package main

import "fmt"

// interfaz que cumple la estructura Computer

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

//Las estructuras son literalmente clases

type Computer struct {
	name  string
	stock int
}

//metodos de Computer para satisfacer la interfaz

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

// Podríamos decir que la estructura Laptop es hija de la
// estructura Computer

type Laptop struct {
	Computer //composition over inheritance
}

// Constructor para crear una nueva laptop

func newLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

// Estructura hija de Computer, por lo que obtiene todos sus
// métodos

type Desktop struct {
	Computer
}

// Constructor para crear una instancia de la estructura Desktop

func newDesktop() IProduct {
	return &Desktop{
		Computer{
			name:  "Desktop Computer",
			stock: 35,
		},
	}
}

// Creamos función para determinar cuál estructura se debe
// instanciar

func GetComputerFactory(computerType string) (IProduct, error) {
	if computerType == "laptop" {
		return newLaptop(), nil
	} else if computerType == "desktop" {
		return newDesktop(), nil
	}

	return nil, fmt.Errorf("Invalid computer type")
}

// imprimimos los productos deseados utilizando la interfaz

func printNameAndStoc(p IProduct) {
	fmt.Printf("PRoduct name: %s, with stock %d\n", p.getName(), p.getStock())
}

func main() {
	laptop, _ := GetComputerFactory("laptop")
	// instancia de la estructura Laptop

	desktop, _ := GetComputerFactory("desktop")
	// instancia de la estructura Desktop

	printNameAndStoc(laptop)
	printNameAndStoc(desktop)
}
