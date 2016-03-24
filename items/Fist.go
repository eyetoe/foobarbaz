package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Fist = Item{
	Name:        "Fist",
	Description: "A trusty knuckle sandwich!",
	Slot:        "Weapon",
	Affects:     []Affect{},
	Attack:      2,
	Damage:      2,
	Crit:        10,
}
