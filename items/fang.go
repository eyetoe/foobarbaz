package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Fang = Item{
	Name:        "Fang",
	Description: "conseals a tiny drop of venom!",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      0,
	Damage:      2,
	Crit:        10,
}
