package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Staff = Item{
	Name:        "Oaken Staff",
	Description: "is weathered beyond belief, and seems to be laquered in blood and sweat.  On close inspection it is engraved with ancient runes, and passages.",
	Slot:        "Weapon",
	Affects:     []Affect{OnFire},
	Attack:      2,
	Damage:      8,
	//Damage
	//Dodge
	//DoT
	Crit: 10,
}
