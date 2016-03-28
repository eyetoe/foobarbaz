package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Fang = Item{
	Name:        "Fang",
	Description: "glistens with such a tiny drop of venom!",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      10,
	Damage:      3,
	Crit:        10,
}
