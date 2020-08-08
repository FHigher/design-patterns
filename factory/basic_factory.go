package factory

import "fmt"

// Animal animal
type Animal interface {
	canEat()
}

// Dog dog
type Dog struct{}

func (d *Dog) canEat() {
	fmt.Println("Dog eat")
}

// Cat cat
type Cat struct{}

func (c *Cat) canEat() {
	fmt.Println("Cat eat")
}

// NewAnimal newAnimal
func NewAnimal(which string) Animal {
	switch which {
	case "d":
		return &Dog{}
	case "c":
		return &Cat{}
	}

	return nil
}
