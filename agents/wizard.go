package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

var Wizard = Agent{
	// Name and Description
	Name:        "Wizard",
	Description: "ancient arcane master.",
	// Stats
	Str: Stat{"Strength", 30},
	Int: Stat{"Intelligence", 60},
	Dex: Stat{"Dexterity", 60},
	// Health and Wellness
	MaxHealth: Stat{"Max Health", 30},
	Health:    Stat{"Current Health", 30},
	Dead:      false,
	// Equipped items
	Weap:  Fireball,
	Armor: Empty,
	Trink: Empty,
	// Special Abilities
	Abl1: "",
	Abl2: "",
	Abl3: "",
	// Inventory
	Inv: []Item{},
	Art: `                         .    .                             ` + "\n" +
		`                             .       =O===                  ` + "\n" +
		`                          . _       %- - %%%                ` + "\n" +
		`                           (_)     % <    D%%               ` + "\n" +
		`                            |      ' ,  ,$$$                ` + "\n" +
		`                           -/-     %%%%  |                  ` + "\n" +
		`                            |//; ,---' _ |----.             ` + "\n" +
		`                             \ )(           /  )            ` + "\n" +
		`                             | \/ \.   '  _.|,  \           ` + "\n" +
		`                             |  \ /(   /    / \_ \          ` + "\n" +
		`                              \ /  (        | /  )          ` + "\n" +
		`                                   (  ,  ___|/ ,'           ` + "\n" +
		`                                   (   _/ (  |  \_          ` + "\n" +
		`                                 _/\--//     ,\\\           ` + "\n" +
		`                       _______ _/\\   / |      |_________   ` + "\n" +
		`                      |_______/   \\ / _/   ,   |________|  ` + "\n" +
		`                      __|___#/     \  _  '      |##____|___ ` + "\n" +
		`                     /______/ _    |  / \   |   /__________\` + "\n" +
		`                           |   /           _/\_/  b'ger     ` + "\n" +
		`                           /_     '_,____\/ )(              ` + "\n" +
		`                          |/ \_   /        (  \             ` + "\n" +
		`                           | /-\_/          Oooo            ` + "\n" +
		`                           )(.                              ` + "\n" +
		`                          /  )                              ` + "\n" +
		`                         oooO                               ` + "\n",
}
