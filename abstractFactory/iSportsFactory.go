package abstractFactory

import (
	"fmt"
	"AbstractFactoryPattern/ConcreteFactory"
	"AbstractFactoryPattern/abstract"
)

type ISportsFactory interface {
	MakeShoe() abstract.IShoe
	MakeShirt() abstract.IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &ConcreteFactory.Adidas{}, nil
	}

	if brand == "nike" {
		return &ConcreteFactory.Nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type")
}
