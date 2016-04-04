package agents

import (
	//	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/items"
)

// A mighty minotaur
var Coyote = Agent{
	// Name and Description
	Name:        "Coyote",
	Description: "thin and desperate, it shows it's teeth.",
	// Stats
	Str: Stat{"Strength", 5},
	Int: Stat{"Intelligence", 4},
	Dex: Stat{"Dexterity", 20},
	// Health and Wellness
	MxHp: Stat{"Max Health", 15},
	Hp:   Stat{"Current Health", 15},
	Dead: false,
	// Equiped items
	Weap:  Maw,
	Armor: "Hide",
	Trink: "Amulet",
	// Special Abilities
	Abl1: "Bite",
	Abl2: "kick",
	Abl3: "howl",
	// Inventory
	Inv: []Item{},
	Art: `                                   ` + "\n" +
		`      /^\      /^\                 ` + "\n" +
		`      |  \    /  |                 ` + "\n" +
		`      ||\ \../ /||                 ` + "\n" +
		`      )'        '(                 ` + "\n" +
		`     ,;'w,    ,w';,                ` + "\n" +
		//		`     ,;'` + Yellow("w") + `,    ,w';,                ` + "\n" +
		`     ;,  ) __ (  ,;                ` + "\n" +
		`      ;  \(\/)/  ;;                ` + "\n" +
		`     ;|  |vwwv|    ''-...          ` + "\n" +
		`      ;  'lwwl'   ;      '''''-.   ` + "\n" +
		`     ;| ; '""' ; ;              '. ` + "\n" +
		`      ;         ,   ,          , | ` + "\n" +
		`      '  ;      ;   l    .     | | ` + "\n" +
		`      ;    ,  ,    |,-,._|      \; ` + "\n" +
		`       ;  ; '' ;   '    \ '\     \;` + "\n" +
		`       |  |    |  |     |   |    |;` + "\n" +
		`       |  ;    ;  |      \   \   (;` + "\n" +
		`       | |      | l       | | \  | ` + "\n" +
		//		`       | |      | |  pb   | |  ) | ` + "\n" +
		//		`       | |      | ;       | |  | | ` + "\n" +
		`       ; ,      : ,      ,_.'  | | ` + "\n" +
		`      :__'      | |           ,_.' ` + "\n" +
		`               '--'                ` + "\n",
}
