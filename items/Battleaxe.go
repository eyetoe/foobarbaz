package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Battleaxe = Item{
	Name:        "Battle Axe",
	Description: "is caked in blood and shows signs of recent use.",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      1,
	Damage:      12,
	Crit:        10,
}
