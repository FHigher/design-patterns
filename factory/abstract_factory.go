package factory

import "fmt"

// DanceAnimal animal
type DanceAnimal interface {
	dance()
}

// DanceDog dog
type DanceDog struct{}

func (d *DanceDog) dance() {
	fmt.Println("Dog dance")
}

// DanceCat cat
type DanceCat struct{}

func (c *DanceCat) dance() {
	fmt.Println("Cat dance")
}

// Factory animal factory
type Factory interface {
	createDog() DanceAnimal
	createCat() DanceAnimal
}

// AnimalFactory animal factory
type AnimalFactory struct{}

// NewAnimalFactory new
func NewAnimalFactory() Factory {
	return &AnimalFactory{}
}

func (a *AnimalFactory) createDog() DanceAnimal {
	return &DanceDog{}
}

func (a *AnimalFactory) createCat() DanceAnimal {
	return &DanceCat{}
}
