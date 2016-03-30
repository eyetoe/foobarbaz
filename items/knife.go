package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Knife = Item{
	Name:        "Knife",
	Description: "a greasy bloodstained shiv!",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      3,
	Damage:      3,
	Crit:        10,
}
