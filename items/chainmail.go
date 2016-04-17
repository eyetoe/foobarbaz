package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Chainmail = Item{
	Name:        "Chainmail",
	Description: "full chain armor designed for medium combat, with a compromise to dodge.",
	Slot:        "Armor",
	Affects:     []Affect{},
	Defence:     2,
	Dodge:       10,
	DropChance:  20,
}
