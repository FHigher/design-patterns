package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	productA := &AProduct{"产品A"}
	productB := &BProduct{"产品B", 0}
	builderA := &ActBuilder{productA}
	builderB := &BBuilder{productB}

	director := NewDirector(builderA)
	director.Builder.Construct()
	director.Builder.(*ActBuilder).AddFeture()

	director2 := NewDirector(builderB)
	director2.Builder.Construct()
	product := director2.Builder.(*BBuilder).GetProduct()
	product.(*BProduct).limitSale(5)
	fmt.Println(product.(*BProduct).getLimitSale())
}
