package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Beak = Item{
	Name:        "Beak",
	Description: "razor sharp, chipped and stained with blood, the beak clacks in anticipation!",
	Slot:        "Weapon",
	Affects:     []Affect{},
	Attack:      4,
	Damage:      6,
	Crit:        10,
}
