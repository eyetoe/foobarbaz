package items

import (
	. "github.com/eyetoe/foobarbaz/affects"
)

var Malaise = Item{
	Name:        "Malaise",
	Description: "you are overwrought with a feeling of dread and dispair!",
	Slot:        "Weapon",
	Affects:     []Affect{Charged},
	Attack:      0,
	Damage:      18,
	Crit:        1,
}
