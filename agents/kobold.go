package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

var Kobold = Agent{
	Name:        "Kobold",
	Description: "small and dexterious it draws it's weapon.",

	// Stats
	Str: Stat{"Strength", 5},
	Int: Stat{"Intelligence", 30},
	Dex: Stat{"Dexterity", 10},

	// Health and Wellness
	MaxHealth: Stat{"Max Health", 5},
	Health:   Stat{"Current Health", 5},
	Dead: false,

	// Equipped items
	Weap:  Hatchet,
	Armor: Waistcoat,
	Trink: Empty,

	// Special Abilities
	Abl1: "Bite",
	Abl2: "kick",
	Abl3: "howl",

	// Inventory
	Inv: []Item{},
	Art: `                   (    )           ` + "\n" +
		`                  ((((()))          ` + "\n" +
		`                  |o\ /o)|          ` + "\n" +
		`                  ( (  _')          ` + "\n" +
		`                   (._.  /\__       ` + "\n" +
		`                  ,\___,/ '  ')     ` + "\n" +
		`    '.,_,,       (  .- .   .    )   ` + "\n" +
		`     \   \\     ( '        )(    )  ` + "\n" +
		`      \   \\    \.  _.__ ____( .  | ` + "\n" +
		`       \  /\\   .(   .'  /\  '.  )  ` + "\n" +
		`        \(  \\.-' ( /    \/    \)   ` + "\n" +
		`         '  ()) _'.-|/\/\/\/\/\|    ` + "\n" +
		`             '\\ .( |\/\/\/\/\/|    ` + "\n" +
		`               '((  \    /\    /    ` + "\n" +
		`               ((((  '.__\/__.')    ` + "\n" +
		`                ((,) /   ((()   )   ` + "\n" +
		`                 "..-,  (()("   /   ` + "\n" +
		`                  _//.   ((() ."    ` + "\n" +
		`          _____ //,/" ___ ((( ', ___` + "\n",
	//		`                           ((  )    ` + "\n" +
	//		`                            / /     ` + "\n" +
	//		`                          _/,/'     ` + "\n" +
	//		`                        /,/,"       ` + "\n",
}
