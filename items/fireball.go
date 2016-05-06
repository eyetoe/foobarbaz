package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Fireball = Item{
	Name:        "Fireball",
	Description: "the power of sheer will crushes your soul.",
	Slot:        "Weapon",
	Affects:     []Affect{},
	Attack:      0,
	Damage:      40,
	Crit:        20,
	DropChance:  10,
}
