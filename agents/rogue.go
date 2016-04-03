package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

// Rogue
var Rogue = Agent{
	// Name and Description
	Name:        "Rogue",
	Description: "shifty and quick, it has larceny in it's repertoire",
	// Stats
	Str: Stat{"Strength", 8},
	Int: Stat{"Intelligence", 14},
	Dex: Stat{"Dexterity", 20},
	// Health and Wellness
	MxHp: Stat{"Max Health", 18},
	Hp:   Stat{"Current Health", 18},
	Dead: false,
	// Equiped items
	Weap:  Dagger,
	Armor: "Leather",
	Trink: "Coin",
	// Special Abilities
	Abl1: "Pickpocket",
	Abl2: "Hide in shadows",
	Abl3: "Backstab",
	// Inventory
	DropChance: 20,
	Inv:        []Item{},
}
