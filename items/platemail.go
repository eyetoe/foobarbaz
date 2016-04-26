package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Platemail = Item{
	Name:        "Platemail",
	Description: "full plate armor designed for heavy combat, with significant compromise to dodge.",
	Slot:        "Armor",
	Affects:     []Affect{},
	Resist:      3,
	Dodge:       15,
	DropChance:  20,
}
