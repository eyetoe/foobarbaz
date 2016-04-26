package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Leathers = Item{
	Name:        "Leathers",
	Description: "full leather armor designed for light combat.",
	Slot:        "Armor",
	Affects:     []Affect{},
	Resist:      1,
	Dodge:       0,
	DropChance:  20,
}
