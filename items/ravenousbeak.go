package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var RavenousBeak = Item{
	Name:        "RavenousBeak",
	Description: "razor sharp, chipped and stained with blood, the beak clacks in anticipation!",
	Slot:        "Weapon",
	Affects:     []Affect{},
	Attack:      6,
	Damage:      39,
	Crit:        1,
}
