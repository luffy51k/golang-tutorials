package main

import (
	"fmt"
)

// Creature is ...
type Creature struct {
	Name string // field
	Real bool
}

// Dump is method of Creature
func (c Creature) Dump() {
	fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}

// FlyingCreature is struct that ke thua tu Creature struct
type FlyingCreature struct {
	Creature
	WingSpan int
}

func main() {
	c := Creature{Name: "Cat", Real: true} // c:= Creature{"Cat", true}
	c.Dump()

	dragon := FlyingCreature{
		Creature{"Dragon", false},
		15,
	}
	fmt.Println(dragon.Name)
	fmt.Println(dragon.Real)
	fmt.Println(dragon.WingSpan)
	dragon.Dump()
}
