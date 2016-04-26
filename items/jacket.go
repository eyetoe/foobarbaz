package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Jacket = Item{
	Name:        "Jacket",
	Description: "old durable weather resistant.",
	Slot:        "Armor",
	Affects:     []Affect{},
	Resist:      1,
	Dodge:       2,
	DropChance:  20,
}
