package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

var Umadon = Agent{
	// Name and Description
	Name:        "Umadon",
	Description: "mystical ancient demigoddess.",
	// Stats
	Str: Stat{"Strength", 50},
	Int: Stat{"Intelligence", 40},
	Dex: Stat{"Dexterity", 40},
	// Health and Wellness
	MaxHealth: Stat{"Max Health", 30},
	Health:    Stat{"Current Health", 30},
	Dead:      false,
	// Equipped items
	Weap:  SoulCrush,
	Armor: Empty,
	Trink: Empty,
	// Special Abilities
	Abl1: "",
	Abl2: "",
	Abl3: "",
	// Inventory
	Inv: []Item{},
	Art: `                                        .--.---.                          ` + "\n" +
		`                                      _ |||||||| _                        ` + "\n" +
		`                                      \\\  |   |//                        ` + "\n" +
		`                                       \_   \  ./                         ` + "\n" +
		`                            .--.         \   \/           .--.            ` + "\n" +
		`                            ||||_        /\.  '\         _||||            ` + "\n" +
		`                            |  ||      ./  /\   \        ||  |            ` + "\n" +
		`                            |  /     ./   /  \   '\      \   |            ` + "\n" +
		`                            |  |    /    /    '\   \      \  |            ` + "\n" +
		`                .--.        |  |   /   ./  ___  \   \     |  |       .--. ` + "\n" +
		`               //| \        |  |   |   |.-'''\''\\   |    |  |       / |\\` + "\n" +
		`               \\\| \       |  |   |   /    __|__|   |    |  |      /  ///` + "\n" +
		`                ''\  \      |  |   |  '.   /     \   |    |  |     /  /'' ` + "\n" +
		`                   \  \     '   \  |   |(\/ o  o |   |   /   '    /  .'   ` + "\n" +
		`                   '   '     \   \ |   |'\    u  |   | ./   /    /   /    ` + "\n" +
		`                    \   '_    \   '\    \ \  -- /    |/    /    /   /     ` + "\n" +
		`                     \    '---.\    \    \/'-._/\   //    /   _/   /      ` + "\n" +
		`                      \_        '  _-    /       ' .-  .-----'    /       ` + "\n" +
		`                        '---.___ /'                     \        /        ` + "\n" +
		`                               ./                        \------'         ` + "\n" +
		`                              /    .-'|            |/\    \               ` + "\n" +
		`                           _./Kali/'  /             \ '\   '-.            ` + "\n",
}
