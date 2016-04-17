package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

// A mighty minotaur
var Drake = Agent{
	// Name and Description
	Name:        "Fire Drake",
	Description: "you feel it's warmth from across the room.  And see it's eyes afire with rage!",
	// Stats
	Str: Stat{"Strength", 45},
	Int: Stat{"Intelligence", 22},
	Dex: Stat{"Dexterity", 45},
	// Health and Wellness
	MxHp: Stat{"Max Health", 45},
	Hp:   Stat{"Current Health", 45},
	Dead: false,
	// Equiped items
	Weap:  Drakestooth,
	Armor: Empty,
	Trink: Empty,
	// Special Abilities
	Abl1: "Fire Breath",
	Abl2: "Tail Slap",
	Abl3: "Horrible Bite",
	// Inventory
	DropChance: 20,
	Inv:        []Item{},
	Art: `              /|                                           |\              ` + "\n" +
		`             /||             ^               ^             ||\             ` + "\n" +
		`            / \\__          //               \\          __// \            ` + "\n" +
		`           /  |_  \         | \   /     \   / |         /  _|  \           ` + "\n" +
		`          /  /  \  \         \  \/ \---/ \/  /         /  /     \          ` + "\n" +
		`         /  /    |  \         \  \/\   /\/  /         /  |       \         ` + "\n" +
		`        /  /     \   \__       \ ( 0\ /0 ) /       __/   /        \        ` + "\n" +
		`       /  /       \     \___    \ \_/|\_/ /    ___/     /\         \       ` + "\n" +
		`      /  /         \_)      \___ \/-\|/-\/ ___/      (_/\ \      '  \      ` + "\n" +
		`     /  /           /\__)       \/  oVo  \/       (__/   ' \      '  \     ` + "\n" +
		`    /  /           /,   \__)    (_/\ _ /\_)    (__/         '      \  \    ` + "\n" +
		`   /  '           //       \__)  (__V_V__)  (__/                    \  \   ` + "\n" +
		`  /  '  '        /'           \   |{___}|   /         .              \  \  ` + "\n" +
		` /  '  /        /              \/ |{___}| \/\          '              \  \ ` + "\n" +
		`/     /        '       .        \/{_____}\/  \          \    '         \  \` + "\n" +
		`     /                ,         /{_______}\   \          \    \         \  ` + "\n" +
		`                     /         /{___/_\___}\   '          \    '           ` + "\n" +
		`                               Adrian Elhart                               ` + "\n",
}
