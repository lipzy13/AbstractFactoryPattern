package ConcreteFactory

import (
	"AbstractFactoryPattern/abstract"
	"AbstractFactoryPattern/concrete"
)

type Nike struct {
}

func (n *Nike) MakeShoe() abstract.IShoe {
	return &concrete.NikeShoe{
		Shoe: abstract.Shoe{
			Logo: "Adidas",
			Size: 14,
		},
	}
}

func (n *Nike) MakeShirt() abstract.IShirt {
	return &concrete.NikeShirt{Shirt: abstract.Shirt{
		Logo: "Nike",
		Size: 14,
	}}
}
