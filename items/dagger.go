package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Dagger = Item{
	Name:        "Dagger",
	Description: "a well balanced ancient piercer!",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      5,
	Damage:      4,
	Crit:        10,
}
