package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Battleaxe = Item{
	Name:        "Battle Axe",
	Description: "is caked in blood and shows signs of recent use.",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      8,
	Damage:      8,
	Crit:        8,
	DropChance:  20,
}
