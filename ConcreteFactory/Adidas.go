package ConcreteFactory

import (
	"AbstractFactoryPattern/abstract"
	"AbstractFactoryPattern/concrete"
)

type Adidas struct {
}

func (a *Adidas) MakeShoe() abstract.IShoe {
	return &concrete.AdidasShoe{
		Shoe: abstract.Shoe{
			Logo: "Adidas",
			Size: 14,
		},
	}
}

func (a *Adidas) MakeShirt() abstract.IShirt {
	return &concrete.AdidasShirt{
		Shirt: abstract.Shirt{
			Logo: "Adidas",
			Size: 14,
		},
	}
}
