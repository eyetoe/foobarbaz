package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
	. "github.com/eyetoe/foobarbaz/util"
)

var randBase = Roll(1, 100)

//var randHP = Roll(1, randBase)

// A mighty minotaur
var Blob = Agent{
	// Name and Description
	Name:        "Blob",
	Description: "undulating mass of living tissue, it senses you, and gurgles in your direction.",
	// Stats
	//Str: Stat{"Strength", Roll(1, randBase)},
	Str: Stat{"Strength", randBase},
	//Int: Stat{"Intelligence", Roll(1, randBase)},
	Int: Stat{"Intelligence", randBase},
	//Dex: Stat{"Dexterity", Roll(1, randBase)},
	Dex: Stat{"Dexterity", randBase},
	// Health and Wellness
	//MaxHealth: Stat{"Max Health", randHP},
	MaxHealth: Stat{"Max Health", randBase},
	//Health:   Stat{"Current Health", randHP},
	Health:   Stat{"Current Health", randBase},
	Dead: false,
	// Equiped items
	Weap:  Blobarm,
	Armor: Empty,
	Trink: Empty,
	// Special Abilities
	Abl1: "Psychic Blast",
	Abl2: "Fear",
	Abl3: "Regenerate",
	// Inventory
	Inv: []Item{},
	Art: `                   o                                     ` + "\n" +
		`                  o$"oo                                  ` + "\n" +
		`                    "$ $ooo$$$$"""o$$oooo                ` + "\n" +
		`            ooo       $$"o"o o        $"""oo             ` + "\n" +
		`            $o"$""""""""o$"$"$$" "$"o$$"$o" "o           ` + "\n" +
		`             "$o        o""$o   """ " " "o$  o$          ` + "\n" +
		`          ooo"""$$o$ "      "$"" $   o   "   $ $"o       ` + "\n" +
		`       o""o  ""oo$$$""$"o    "    o""   $"  " o$$ ""oo   ` + "\n" +
		`     o"o"   oo"$$$$     o   " "o        " o$ $"$ "o " "o ` + "\n" +
		`    $     o$ o" o"$$  "" "  ""$  oo"$o $o$$$$$oo o$   "o$` + "\n" +
		`    "$o    "$"o""o$$$ $$ $oo$$$$$$o$$$$$ o  oo        "o"` + "\n" +
		`      ""oo  ""oo$$"$   $$" o$" $$"""$oo$  oo$     o  o"  ` + "\n" +
		`         """ooooo   """""o$o$""""  "    """"ooooooo""    ` + "\n" +
		`                 """"$$$$ooooooooooo$$$$ooo""            ` + "\n",
}
