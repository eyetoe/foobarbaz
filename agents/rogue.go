package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

// Rogue
var Rogue = Agent{
	// Name and Description
	Name:        "Rogue",
	Description: "shifty and quick, it has larceny in it's repertoire",
	// Stats
	Str: Stat{"Strength", 8},
	Int: Stat{"Intelligence", 14},
	Dex: Stat{"Dexterity", 20},
	// Health and Wellness
	MaxHealth: Stat{"Max Health", 18},
	Health:   Stat{"Current Health", 18},
	Dead: false,
	// Equiped items
	Weap:  Dagger,
	Armor: Jacket,
	Trink: Empty,
	// Special Abilities
	Abl1: "Pickpocket",
	Abl2: "Hide in shadows",
	Abl3: "Backstab",
	// Inventory
	Inv: []Item{},
	Art: `                       __.------.                          ` + "\n" +
		`                      (__  ___   )                         ` + "\n" +
		`                        .)e  )\ /                          ` + "\n" +
		`                       /_.------                           ` + "\n" +
		`                       _/_    _/                           ` + "\n" +
		`                   __.'  / '   '-.__                       ` + "\n" +
		`                  / <.--'           '\                     ` + "\n" +
		//		`                 /   \   \c           |                    ` + "\n" +
		//		`                /    /    )  @o.  x    \                   ` + "\n" +
		`                |   /\    |c     / \.-  \                  ` + "\n" +
		`                \__/  )  /(     (   \   <>'\               ` + "\n" +
		`                     / _/ _\-    '-. \/_|_ /<>             ` + "\n" +
		`                    / /--/,-\     _ \     <>.'.            ` + "\n" +
		`                    \/'--\_._) - /   '-/\    '.\           ` + "\n" +
		`                     /        '.     /   )     '\          ` + "\n" +
		`                     \      \   \___/----'                 ` + "\n" +
		`                     |      /    '(                        ` + "\n" +
		//		` ___________         \    ./\_   _ \                       ` + "\n" +
		`   ______________    /     |  )    '|                      ` + "\n" +
		` __________________ |     /   \     \     ___________      ` + "\n" +
		`                   /     |     |____.)                     ` + "\n" +
		//		`                  /      \  a88a\___/88888a.               ` + "\n" +
		`                 \_      :)8888888888888888888a.           ` + "\n" +
		`                /' '-----'  'Y88888888888888888            ` + "\n" +
		`                \____|         '88888888888P'              ` + "\n",
}
