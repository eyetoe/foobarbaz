package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

// an annoying pixie
var Pixie = Agent{
	Name:        "Pixie",
	Description: "small and mischevious, it looks at you with murderous eyes.",
	// Stats
	Str: Stat{"Strength", 2},
	Int: Stat{"Intelligence", 50},
	Dex: Stat{"Dexterity", 20},
	// Health and Wellness
	MxHp: Stat{"Max Health", 5},
	Hp:   Stat{"Current Health", 5},
	Dead: false,
	// Equiped items
	Weap:  Tack,
	Armor: Vest,
	Trink: Empty,
	// Special Abilities
	Abl1: "Bite",
	Abl2: "kick",
	Abl3: "howl",
	// Inventory
	Inv: []Item{},
	Art: `                     ,_  .--.      ` + "\n" +
		`               , ,   _)\/    ;--.  ` + "\n" +
		`       . ' .    \_\-'   |  .'    \ ` + "\n" +
		`      -= * =-   (.-,   /  /       |` + "\n" +
		`       ' .\'    ).  ))/ .'   _/\ / ` + "\n" +
		`           \_   \_  /( /     \ /(  ` + "\n" +
		`           /_\ .--'   \-.    //  \ ` + "\n" +
		`           ||\/        , '._//    |` + "\n" +
		`           ||/ /^(_ (_,;^-._/     /` + "\n" +
		`           \_.'   )   / \       .' ` + "\n" +
		`                .' .  |  ;.   /    ` + "\n" +
		`               /      |\(  '.(     ` + "\n" +
		`              |   |/  | '    '     ` + "\n" +
		`              |   |  /             ` + "\n" +
		`              |   |.'              ` + "\n" +
		`           __/'  /                 ` + "\n" +
		`       _ .'  _.-'                  ` + "\n" +
		`    _.' '.-;'/                     ` + "\n" +
		`   /_.-'' ( /                      ` + "\n" +
		`          \/                       ` + "\n",
	//		`        ( /                        ` + "\n" +
	//		`       /_/                         ` + "\n",
}
