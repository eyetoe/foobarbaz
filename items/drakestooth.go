package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Drakestooth = Item{
	Name:        "Drakestooth",
	Description: "sword tooth of molten rock to divide a torso!",
	Slot:        "Weapon",
	Affects:     []Affect{},
	Attack:      6,
	Damage:      10,
	Crit:        10,
	DropChance:  20,
}
