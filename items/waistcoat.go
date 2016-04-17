package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Waistcoat = Item{
	Name:        "Waistcoat",
	Description: "work worn solid garment.",
	Slot:        "Armor",
	Affects:     []Affect{},
	Defence:     0,
	Dodge:       5,
	DropChance:  20,
}
