package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

// A mighty minotaur
var FlyingPig = Agent{
	// Name and Description
	Name:        "Flying Pig",
	Description: "... yes, you rub your eyes.  It IS a flying pig!",
	// Stats
	Str: Stat{"Strength", 60},
	Int: Stat{"Intelligence", 40},
	Dex: Stat{"Dexterity", 30},
	// Health and Wellness
	MxHp: Stat{"Max Health", 40},
	Hp:   Stat{"Current Health", 40},
	Dead: false,
	// Equiped items
	Weap:  Snort,
	Armor: "none",
	Trink: "none",
	// Special Abilities
	Abl1: "Flying Fart",
	Abl2: "Snorting Swoop",
	Abl3: "Trash Gobble",
	// Inventory
	DropChance: 0,
	Inv:        []Item{},
	Art: `           ____ _____    ` + "\n" +
		`          (    /,----'   ` + "\n" +
		`           \  //==--'    ` + "\n" +
		`            \//==--'     ` + "\n" +
		`         _//|Y.-~~~~-,   ` + "\n" +
		`       _/66  \        \_@` + "\n" +
		`      (")_   /   /    |  ` + "\n" +
		`        '--'|| |-\   /   ` + "\n" +
		`   jgs      //_/  /_/    ` + "\n",
}
