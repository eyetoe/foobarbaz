package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Fist = Item{
	Name:        "Fist",
	Description: "is ready. Taste a trusty knuckle sandwich!",
	Slot:        "Weapon",
	Affects:     []Affect{},
	Attack:      2,
	Damage:      2,
	Crit:        1,
}
