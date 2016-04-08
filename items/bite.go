package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Bite = Item{
	Name:        "Bite",
	Description: "fetid black teeth glisten in the shadows!",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      6,
	Damage:      6,
	Crit:        6,
}
