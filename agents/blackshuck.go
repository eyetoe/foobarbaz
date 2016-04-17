package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

// A mighty minotaur
var Blackshuck = Agent{
	// Name and Description
	Name:        "Black Shuck",
	Description: "dark embers glow in the shadows as the beast emits a fearsome growl.",
	// Stats
	Str: Stat{"Strength", 25},
	Int: Stat{"Intelligence", 5},
	Dex: Stat{"Dexterity", 25},
	// Health and Wellness
	MxHp: Stat{"Max Health", 25},
	Hp:   Stat{"Current Health", 25},
	Dead: false,
	// Equiped items
	Weap:  Bite,
	Armor: Empty,
	Trink: Empty,
	// Special Abilities
	Abl1: "Growl",
	Abl2: "Sneak",
	Abl3: "Leap",
	// Inventory
	DropChance: 0,
	Inv:        []Item{},
	Art: `                                            ` + "\n" +
		`                     /\                     ` + "\n" +
		`                    ( ;'~v/~~~ ;._          ` + "\n" +
		`                 ,/'"/^) ' < o\  '".~'\\\--,` + "\n" +
		`               ,/",/W  u ''. ~  >,._..,   )'` + "\n" +
		`              ,/'  w  ,U^v  ;//^)/')/^\;~)' ` + "\n" +
		`           ,/"'/   W' ^v  W |;         )/'  ` + "\n" +
		`         ;''  |  v' v'" W }  \\             ` + "\n" +
		`        "    .'\    v  'v/^W,) '\)\.)\/)    ` + "\n" +
		`                 '\   ,/,)'   ''')/^"-;'    ` + "\n" +
		`                      \                     ` + "\n" +
		`                    ".                      ` + "\n" +
		`                   \    Robert Kirkpatrick  ` + "\n",
}
