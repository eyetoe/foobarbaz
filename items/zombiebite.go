package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var ZombieBite = Item{
	Name:        "Zombie Bite",
	Description: "infected, foaming, hungry undead mouth gapes and snaps!",
	Slot:        "Weapon",
	Affects:     []Affect{},
	Attack:      0,
	Damage:      13,
	Crit:        1,
	DropChance:  0,
}
