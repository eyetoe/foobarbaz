package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Tack = Item{
	Name:        "Tack",
	Description: "provides a well placed poke!",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      10,
	Damage:      3,
	Crit:        10,
}
