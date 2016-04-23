package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Snort = Item{
	Name:        "Snort",
	Description: "a flying snort!",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      5,
	Damage:      4,
	Crit:        1,
}
