package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

var Phantom = Agent{
	// Name and Description
	Name:        "Phantom",
	Description: "withered, unatural and ethereal the horror loiters.",
	// Stats
	Str: Stat{"Strength", 2},
	Int: Stat{"Intelligence", 15},
	Dex: Stat{"Dexterity", 25},
	// Health and Wellness
	MaxHealth: Stat{"Max Health", 1},
	Health:   Stat{"Current Health", 1},
	Dead: false,
	// Equipped items
	Weap:  Shriek,
	Armor: Empty,
	Trink: Empty,
	// Special Abilities
	Abl1: "Bite",
	Abl2: "kick",
	Abl3: "howl",
	// Inventory
	Inv: []Item{},
	Art: `                                ` + "\n" +
		`      .''''.      ...           ` + "\n" +
		`     :o  o '....''  ;           ` + "\n" +
		`     '. O         :'            ` + "\n" +
		`       '':          '.          ` + "\n" +
		`         ':.          '.        ` + "\n" +
		`          : '.         '.       ` + "\n" +
		`         '..''...       '.      ` + "\n" +
		`                 '...     '.    ` + "\n" +
		`                     ''...  '.  ` + "\n" +
		`                          '''''.` + "\n",
}
