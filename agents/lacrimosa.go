package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

// A mighty minotaur
var Lacrimosa = Agent{
	// Name and Description
	Name:        "Lacrimosa",
	Description: "psychic, flying, eye-snake looks at you with malicious intelligence",
	// Stats
	Str: Stat{"Strength", 20},
	Int: Stat{"Intelligence", 40},
	Dex: Stat{"Dexterity", 30},
	// Health and Wellness
	MxHp: Stat{"Max Health", 25},
	Hp:   Stat{"Current Health", 25},
	Dead: false,
	// Equiped items
	Weap:  Malaise,
	Armor: "none",
	Trink: "none",
	// Special Abilities
	Abl1: "Psychic Blast",
	Abl2: "Fear",
	Abl3: "Regenerate",
	// Inventory
	DropChance: 0,
	Inv:        []Item{},
	Art: `                                                    '-. ` + "\n" +
		`                                                      .'` + "\n" +
		`                                                   _.'.'` + "\n" +
		`                                               _.-' .'  ` + "\n" +
		`                                       ___.---' _.-'    ` + "\n" +
		`                               __..---'___..---'        ` + "\n" +
		`                      _...--.-'   _.--'                 ` + "\n" +
		`                  _.-'.-'.-'  _.-'                      ` + "\n" +
		`               .-' .'  .'   .'                          ` + "\n" +
		`    .         /   /   /    /                    .       ` + "\n" +
		`    \'-.._    |  |    \    '.              _..-'/       ` + "\n" +
		`   .'-.._ ''--.._\     '. -- '.      _..-''  _..-'.     ` + "\n" +
		`   '_    _       '-. .'        '. .-'      _    _'      ` + "\n" +
		`     '.''           .            \          ''.'        ` + "\n" +
		`      '-.-'    _   .              :   _   '-.-'         ` + "\n" +
		`        '..-..'    ;       .' '.  '    '..-..'          ` + "\n" +
		`            /      .      : .-. : :        \            ` + "\n" +
		`            '._     \     ;( O ) /      _.'             ` + "\n" +
		`               '-._.''.    .'-'.' '._.-'                ` + "\n" +
		`                       '-....-'                         ` + "\n",
}
