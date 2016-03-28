package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Bite = Item{
	Name:        "Rabid Bite",
	Description: "looms foamy and wet. The mouth craves blood!",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      10,
	Damage:      4,
	Crit:        10,
}
