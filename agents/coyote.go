package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

// A mighty minotaur
var Coyote = Agent{
	// Name and Description
	Name:        "Coyote",
	Description: "Thin and desperate, the coyote shows it's teeth.",
	// Stats
	Str: Stat{"Strength", 10},
	Int: Stat{"Intelligence", 5},
	Dex: Stat{"Dexterity", 20},
	// Health and Wellness
	MxHp: Stat{"Max Health", 20},
	Hp:   Stat{"Current Health", 20},
	Dead: false,
	// Equiped items
	Weap:  Bite,
	Armor: "Hide",
	Trink: "Amulet",
	// Special Abilities
	Abl1: "Bite",
	Abl2: "kick",
	Abl3: "howl",
	// Inventory
	Inv: []Item{},
}
