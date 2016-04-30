package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

var Swordsman = Agent{
	// Name and Description
	Name:        "Swordsman",
	Description: "trained in the art of the sword, you are his target.",
	// Stats
	Str: Stat{"Strength", 45},
	Int: Stat{"Intelligence", 25},
	Dex: Stat{"Dexterity", 25},
	// Health and Wellness
	MaxHealth: Stat{"Max Health", 40},
	Health:    Stat{"Current Health", 40},
	Dead:      false,
	// Equipped items
	Weap:  Broadsword,
	Armor: Platemail,
	Trink: Shield,
	// Special Abilities
	Abl1: "",
	Abl2: "",
	Abl3: "",
	// Inventory
	Inv: []Item{},
	Art: `                         |  ||         __.--._                       ` + "\n" +
		`                         |  ||      /~~   __.-~\ _                   ` + "\n" +
		`                         |  ||  _.-~ / _---._ ~-\/~\                 ` + "\n" +
		`                         |  || // /  /~/  .-  \  /~-\                ` + "\n" +
		`                         |  ||((( /(/_(.-(-~~~~~-)_/ |               ` + "\n" +
		`                         |  || ) (( |_.----~~~~~-._\ /               ` + "\n" +
		`                         |  ||    ) |              \_|               ` + "\n" +
		`                         |  ||     (| =-_   _.-=-  |~)        ,      ` + "\n" +
		`                         |  ||      | '~~ |   ~~'  |/~-._-'/'/_,     ` + "\n" +
		`                         |  ||       \    |        /~-.__---~ , ,    ` + "\n" +
		`                         |  ||       |   ~-''     || '\_~~~----~     ` + "\n" +
		`                         |  ||_.ssSS$$\ -====-   / )\_  ~~--~        ` + "\n" +
		`                 ___.----|~~~|%$$$$$$/ \_    _.-~ /' )$s._           ` + "\n" +
		`        __---~-~~        |   |%%$$$$/ /  ~~~~   /'  /$$$$$$$s__      ` + "\n" +
		`      /~       ~\    ============$$/ /        /'  /$$$$$$$$$$$SS-.   ` + "\n" +
		`    /'      ./\\\\\\_( ~---._(_))$/ /       /'  /$$$$%$$$$$~      \  ` + "\n" +
		`    (      //////////(~-(..___)/$/ /      /'  /$$%$$%$$$$'         \ ` + "\n" +
		`     \    |||||||||||(~-(..___)$/ /  /  /'  /$$$%$$$%$$$            |` + "\n" +
		`      '-__ \\\\\\\\\\\(-.(_____) /  / /'  /$$$$%$$$$$%$             |` + "\n" +
		`          ~~""""""""""-\.(____) /   /'  /$$$$$%%$$$$$$\_            /` + "\n" +
		`                        $|===|||  /'  /$$$$$$$%%%$$$$$( ~         ,'|` + "\n",
}
