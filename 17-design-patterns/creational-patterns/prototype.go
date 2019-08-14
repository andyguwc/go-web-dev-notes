/*
Prototype pattern 

While with the Builder pattern, we are dealing with repetitive building algorithms and with
the factories we are simplifying the creation of many types of objects; with the Prototype
pattern, we will use an already created instance of some type to clone it and complete it
with the particular needs of each context.

The aim of the Prototype pattern is to have an object or a set of objects that is already
created at compilation time, but which you can clone as many times as you want at runtime.

*/

package creational

import (
	"errors"
	"fmt"
)

// first a cloner interface and an object that implements it 
type ShirtCloner interface {
	GetClone(m int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

// a function to retrieve a new instance of the cloner
func GetShirtsCloner() ShirtCloner {
	return nil
}

type ShirtsCache struct{}

func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

// object struct ot cline 
type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

// implements an interface to retrive information of its fields 
func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

func (i *Shirt) GetPrice() float32 {
	return i.Price
}
