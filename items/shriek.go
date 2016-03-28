package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Shriek = Item{
	Name:        "Shriek",
	Description: "A blood curdling shriek damages your soul!",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      20,
	Damage:      20,
	Crit:        10,
}