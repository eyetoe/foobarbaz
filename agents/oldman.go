package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

// Old Man
var Oldman = Agent{
	// Name and Description
	Name:        "Old Man",
	Description: "appears feeble and insane. He has a knowing glint in his eye.",
	// Stats
	Str: Stat{"Strength", 2},
	Int: Stat{"Intelligence", 48},
	Dex: Stat{"Dexterity", 4},
	// Health and Wellness
	MxHp: Stat{"Max Health", 25},
	Hp:   Stat{"Current Health", 25},
	Dead: false,
	// Equiped items
	Weap:  Staff,
	Armor: "Tattered Robes",
	Trink: "Black Cat",
	// Special Abilities
	Abl1: "Fireball",
	Abl2: "Shocking Touch",
	Abl3: "Heal",
	// Inventory
	DropChance: 20,
	Inv:        []Item{},
}
