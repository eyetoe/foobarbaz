package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

// A mighty minotaur
var New = Agent{
	// Name and Description
	Name: "New",
	//Description: "Thin and desperate, the coyote shows it's teeth.",
	Description: "A person with great potential.",
	// Stats
	Str: Stat{"Strength", 10},
	Int: Stat{"Intelligence", 10},
	Dex: Stat{"Dexterity", 10},
	// Health and Wellness
	MxHp: Stat{"Max Health", 10},
	Hp:   Stat{"Current Health", 10},
	Dead: false,
	// Equiped items
	Weap:  Fist,
	Armor: Pancho,
	Trink: Empty,
	// Special Abilities
	Abl1: "",
	Abl2: "",
	Abl3: "",
	// Inventory
	Inv: []Item{},
}
