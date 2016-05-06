package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var SoulCrush = Item{
	Name:        "Soul Crush",
	Description: "the power of sheer will crushes your soul.",
	Slot:        "Weapon",
	Affects:     []Affect{},
	Attack:      25,
	Damage:      25,
	Crit:        25,
}
