package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Broadsword = Item{
	Name:        "Broadsword",
	Description: "the finest steel in the hands of an expert is neigh unstoppable.",
	Slot:        "Weapon",
	Affects:     []Affect{},
	Attack:      12,
	Damage:      12,
	Crit:        12,
	DropChance:  20,
}
