package factory

import "testing"

func TestNewAnimalFactory(t *testing.T) {
	NewAnimalFactory().createDog().dance()
	NewAnimalFactory().createCat().dance()
}
