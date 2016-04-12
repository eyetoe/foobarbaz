package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var DrakeBite = Item{
	Name:        "Drake Bite",
	Description: "teeth of molten rock snap at your torso!",
	Slot:        "Weapon",
	Affects:     []Affect{},
	Attack:      70,
	Damage:      5,
	Crit:        10,
}
