package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var GoatKick = Item{
	Name:        "Goat Kick",
	Description: "it's razor sharp hooves are ready for action!",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      10,
	Damage:      8,
	Crit:        10,
}
