package factory

import "testing"

func TestNewAnimal(t *testing.T) {
	NewAnimal("d").canEat()
	NewAnimal("c").canEat()
}
