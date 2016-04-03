package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

var Kobold = Agent{
	Name:        "Kobold",
	Description: "small and dexterious it draws it's knife.",

	// Stats
	Str: Stat{"Strength", 5},
	Int: Stat{"Intelligence", 30},
	Dex: Stat{"Dexterity", 10},

	// Health and Wellness
	MxHp: Stat{"Max Health", 5},
	Hp:   Stat{"Current Health", 5},
	Dead: false,

	// Equiped items
	Weap:  Knife,
	Armor: "Hide",
	Trink: "Amulet",

	// Special Abilities
	Abl1:       "Bite",
	Abl2:       "kick",
	Abl3:       "howl",
	DropChance: 20,

	// Inventory
	Inv: []Item{},
}
