package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

var Spider = Agent{
	Name:        "Spider",
	Description: "poised and ready, you feel something on your shoulder.",
	// Stats
	Str: Stat{"Strength", 1},
	Int: Stat{"Intelligence", 10},
	Dex: Stat{"Dexterity", 40},
	// Health and Wellness
	MxHp: Stat{"Max Health", 2},
	Hp:   Stat{"Current Health", 2},
	Dead: false,
	// Equiped items
	Weap:  Fang,
	Armor: "hair",
	Trink: "eggs",
	// Special Abilities
	Abl1: "Bite",
	Abl2: "kick",
	Abl3: "howl",
	// Inventory
	Inv: []Item{},

	Art: `  ___________________________________________` + "\n" +
		` |\'-._(   /                                 |` + "\n" +
		` | \  .'-._\                            ,   ,|` + "\n" +
		` |-.\'    .-;                         .'\'-' |` + "\n" +
		` |   \  .' (                       _.'   \   |` + "\n" +
		` |.--.\'   _)                   ;-;       \._|` + "\n" +
		` |    ' _\(_)/_                  \ ''-,_,-'\ |` + "\n" +
		` |______ /(O)\  _________________/____)_'-._\|` + "\n",
}
