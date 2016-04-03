package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

// A mighty minotaur
var Minotaur = Agent{
	// Name and Description
	Name:        "Minotaur",
	Description: "half man half bull, it has hatred in it's eyes.",
	// Stats
	Str: Stat{"Strength", 30},
	Int: Stat{"Intelligence", 5},
	Dex: Stat{"Dexterity", 5},
	// Health and Wellness
	MxHp: Stat{"Max Health", 40},
	Hp:   Stat{"Current Health", 40},
	Dead: false,
	// Equiped items
	Weap:  Battleaxe,
	Armor: "Hide",
	Trink: "Amulet",
	// Special Abilities
	Abl1: "Roar",
	Abl2: "Stomp",
	Abl3: "Snort",
	// Inventory
	DropChance: 20,
	Inv:        []Item{},
	Art: `    ,           ,   ` + "\n" +
		`   /             \  ` + "\n" +
		`  ((__-^^-,-^^-__)) ` + "\n" +
		`   '-_---' '---_-'  ` + "\n" +
		`    <__|o' 'o|__>   ` + "\n" +
		`       \  '  /      ` + "\n" +
		`        ): :(       ` + "\n" +
		`        :o_o:       ` + "\n" +
		`         "-"        ` + "\n",
}
