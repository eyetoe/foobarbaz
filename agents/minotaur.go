package agents

import "github.com/eyetoe/foobarbaz/inv"

// A mighty minotaur
var Minotaur = Agent{
	// Name and Description
	Name:        "Minotaur",
	Description: "Half man half bull, the minotaur has hatred in it's eyes",
	// Stats
	Str: 30,
	Int: 5,
	Dex: 5,
	// Health and Wellness
	MxHp: 40,
	Hp:   40,
	Dead: false,
	// Equiped items
	Weap:  "Axe",
	Armor: "Hide",
	Trink: "Amulet",
	// Special Abilities
	Abl1: "Roar",
	Abl2: "Stomp",
	Abl3: "Snort",
	// Inventory
	Inv: []inv.Item{},
}
