package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Pancho = Item{
	Name:        "Pancho",
	Description: "is a dusty old bolt of heavy fabric with a hole in it for your head.",
	Slot:        "Armor",
	Affects:     []Affect{},
	Resist:      0,
	Dodge:       10,
}
