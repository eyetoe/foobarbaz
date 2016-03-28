package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

// an annoying pixie
var Pixie = Agent{
	Name:        "Pixie",
	Description: "looks at you with a mischievous grin.",
	// Stats
	Str: Stat{"Strength", 2},
	Int: Stat{"Intelligence", 50},
	Dex: Stat{"Dexterity", 20},
	// Health and Wellness
	MxHp: Stat{"Max Health", 5},
	Hp:   Stat{"Current Health", 5},
	Dead: false,
	// Equiped items
	Weap:  Tack,
	Armor: "none",
	Trink: "none",
	// Special Abilities
	Abl1: "Bite",
	Abl2: "kick",
	Abl3: "howl",
	// Inventory
	DropChance: 20,
	Inv:        []Item{},
}
