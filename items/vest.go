package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Vest = Item{
	Name:        "Vest",
	Description: "modest, no sleeves.",
	Slot:        "Armor",
	Affects:     []Affect{},
	Resist:      0,
	Dodge:       1,
	DropChance:  20,
}
