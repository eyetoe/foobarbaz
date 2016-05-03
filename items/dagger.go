package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Dagger = Item{
	Name:        "Dagger",
	Description: "a well balanced ancient piercer!",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      11,
	Damage:      5,
	Crit:        5,
	DropChance:  20,
}
