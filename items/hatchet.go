package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Hatchet = Item{
	Name:        "Hatchet",
	Description: "chipped and weathered, it seems to have chopped more heads than trees!",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      2,
	Damage:      4,
	Crit:        3,
	DropChance:  20,
}
