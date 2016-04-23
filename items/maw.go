package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Maw = Item{
	Name:        "Rabid Maw",
	Description: "glistens foamy and wet. The mouth craves blood!",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      1,
	Damage:      4,
	Crit:        1,
}
