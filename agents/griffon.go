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
	Str: Stat{"Strength", 40},
	Int: Stat{"Intelligence", 5},
	Dex: Stat{"Dexterity", 5},
	// Health and Wellness
	MaxHealth: Stat{"Max Health", 30},
	Health:   Stat{"Current Health", 30},
	Dead: false,
	// Equipped items
	Weap:  Beak,
	Armor: Empty,
	Trink: Empty,
	// Special Abilities
	Abl1: "Bite",
	Abl2: "Stomp",
	Abl3: "Rush",
	// Inventory
	Inv: []Item{},
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
