package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Shield = Item{
	Name:        "Shield",
	Description: "heavy barrier of steel and wood.",
	Slot:        "Trinket",
	Affects:     []Affect{},
	Resist:      2,
	Dodge:       25,
	DropChance:  20,
}
