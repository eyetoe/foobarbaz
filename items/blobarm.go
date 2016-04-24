package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
	. "github.com/eyetoe/foobarbaz/util"
)

var Blobarm = Item{
	Name:        "Blob Arm",
	Description: "extends in your direction with a horrifying yearning.",
	Slot:        "Weapon",
	Affects:     []Affect{},
	Attack:      Roll(3, 50),
	Damage:      Roll(3, 50),
	Crit:        1,
}
