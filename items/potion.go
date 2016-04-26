package items

import . "github.com/eyetoe/foobarbaz/affects"

var Potion = Item{
	Name:        "Potion",
	Description: "A mysterious milky blue elixir!",
	Slot:        "Inventory",
	Affects:     []Affect{},
	Attack:      0,
	Damage:      0,
	Crit:        0,
	DropChance:  20,
}
