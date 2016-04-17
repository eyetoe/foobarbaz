package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

// A mighty minotaur
var Goat = Agent{
	// Name and Description
	Name:        "Goat",
	Description: "screaming and furious the goat is enraged!",
	// Stats
	Str: Stat{"Strength", 55},
	Int: Stat{"Intelligence", 2},
	Dex: Stat{"Dexterity", 55},
	// Health and Wellness
	MxHp: Stat{"Max Health", 55},
	Hp:   Stat{"Current Health", 55},
	Dead: false,
	// Equiped items
	Weap:  GoatKick,
	Armor: Platemail,
	Trink: Empty,
	// Special Abilities
	Abl1: "Horns-o-doom",
	Abl2: "Goat Scream",
	Abl3: "Iron Stomach",
	// Inventory
	Inv: []Item{},
	Art: `             / /                            ` + "\n" +
		`          (\/_//')                          ` + "\n" +
		`           /   '/                           ` + "\n" +
		`          0  0   \                          ` + "\n" +
		`         /        \                         ` + "\n" +
		`        /    __/   \                        ` + "\n" +
		`       /,  _/ \     \_                      ` + "\n" +
		`       '-./ )  |     ~^~^~^~^~^~^~^~\~.     ` + "\n" +
		`           (   /                     \_}    ` + "\n" +
		`              ;               /      |      ` + "\n" +
		`               \/ ,/           \    |       ` + "\n" +
		`               / /~~|~|~~~~~~|~|\   |       ` + "\n" +
		`              / /   | |      | | '\ \       ` + "\n" +
		`            / (     | |      | |    \ \     ` + "\n" +
		`     jgs   /,_)    /__)     /__)   /,_/     ` + "\n" +
		`    '''''"""""'''""""""'''""""""''"""""'''''` + "\n",
}
