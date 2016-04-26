package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Tack = Item{
	Name:        "Tack",
	Description: "provides a well placed poke!",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      4,
	Damage:      4,
	Crit:        6,
	DropChance:  20,
}
