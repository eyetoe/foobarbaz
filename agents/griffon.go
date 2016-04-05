package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

// A mighty minotaur
var Griffon = Agent{
	// Name and Description
	Name:        "Griffon",
	Description: "majestic yet ferocious, the beast is hungry.",
	// Stats
	Str: Stat{"Strength", 20},
	Int: Stat{"Intelligence", 5},
	Dex: Stat{"Dexterity", 5},
	// Health and Wellness
	MxHp: Stat{"Max Health", 30},
	Hp:   Stat{"Current Health", 30},
	Dead: false,
	// Equiped items
	Weap:  Beak,
	Armor: "Feathers",
	Trink: "Quill",
	// Special Abilities
	Abl1: "Bite",
	Abl2: "Stomp",
	Abl3: "Rush",
	// Inventory
	DropChance: 0,
	Inv:        []Item{},
	Art: `                         //           //                     ` + "\n" +
		//		`                        ///          ///                     ` + "\n" +
		`                       ////         ////                     ` + "\n" +
		`                       |////       /////                     ` + "\n" +
		//		`                       |))//;     /)))//;                    ` + "\n" +
		`                      /)))))/;   /)))))/;                    ` + "\n" +
		`                  .---^,))))/;  /)))))))/;                   ` + "\n" +
		`              __--\/6-  \^))/; |)))))))/;                    ` + "\n" +
		`             (----/    \\\^^;  |))))))/;                     ` + "\n" +
		`                ~/-\  \\\\\^^   \))))))/;                    ` + "\n" +
		//		`                    \\\\\\\\^    |)))))/;                    ` + "\n" +
		`                    |\\\\\\\\___/))))))/;__-------.          ` + "\n" +
		`                    //////|  %%_/))))))/;           \___,    ` + "\n" +
		`                   |||||||\   \%%%%%%;:              \_. \   ` + "\n" +
		`                   |\\\\\\\\\                        |  | |  ` + "\n" +
		`                    \\\\\\\                          |  | |  ` + "\n" +
		`                     |\\\\               __|        /   / /  ` + "\n" +
		`                     | \\__\     \___----  |       |   / /   ` + "\n" +
		`                     |    / |     )     \   \      \  / /    ` + "\n" +
		`                     |   /  |    /       \   \      >/ /  ,, ` + "\n" +
		//		`                     |   |  |   |         |   |    // /  //, ` + "\n" +
		`                     |   |  |   |         |   |   /| |   |\\,` + "\n" +
		`                  _--'   _--'   |     _---_---'  |  \ \__/\|/` + "\n" +
		`                 (-(-===(-(-(===/    (-(-=(-(-(==/   \____/  ` + "\n",
}
