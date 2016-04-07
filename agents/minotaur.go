package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

// A mighty minotaur
var Minotaur = Agent{
	// Name and Description
	Name:        "Minotaur",
	Description: "half man half bull, it has hatred in it's eyes.",
	// Stats
	Str: Stat{"Strength", 30},
	Int: Stat{"Intelligence", 5},
	Dex: Stat{"Dexterity", 5},
	// Health and Wellness
	MxHp: Stat{"Max Health", 40},
	Hp:   Stat{"Current Health", 40},
	Dead: false,
	// Equiped items
	Weap:  Battleaxe,
	Armor: "Hide",
	Trink: "Amulet",
	// Special Abilities
	Abl1: "Roar",
	Abl2: "Stomp",
	Abl3: "Snort",
	// Inventory
	DropChance: 20,
	Inv:        []Item{},
	Art: `                                                                _    ` + "\n" +
		`                                                              _( (~\ ` + "\n" +
		`       _ _                        /                          ( \> > \` + "\n" +
		`   -/~/ / ~\                     :;                \       _  > /(~\/` + "\n" +
		`  || | | /\ ;\                   |l      _____     |;     ( \/    > >` + "\n" +
		`  _\\)\)\)/ ;;;                  '8o __-~     ~\   d|      \      // ` + "\n" +
		` ///(())(__/~;;\                  "88p;.  -. _\_;.oP        (_._/ /  ` + "\n" +
		`(((__   __ \\   \                  '>,% (\  (\./)8"         ;:'  i   ` + "\n" +
		`)))--'.'-- (( ;,8 \               ,;%%%:  ./V^^^V'          ;.   ;.  ` + "\n" +
		`((\   |   /)) .,88  ': ..,,;;;;,-::::::'_::\   ||\         ;[8:   ;  ` + "\n" +
		` )|  ~-~  |(|(888; ..'''::::8888oooooo.  :\'^^^/,,~--._    |88::  |  ` + "\n" +
		` |\ -===- /|  \8;; '':.      oo.8888888888:'((( o.ooo8888Oo;:;:'  |  ` + "\n" +
		` |_~-___-~_|   '-\.   '        'o'88888888b' )) 888b88888P""'     ;  ` + "\n" +
		` ; ~~~~;~~         "'--_'.       b'888888888;(.,"888b888"  ..::;-'   ` + "\n" +
		`   ;      ;              ~"-....  b'8888888:::::.'8888. .:;;;''      ` + "\n" +
		`      ;    ;                 ':::. ':::OOO:::::::.'OO' ;;;''         ` + "\n" +
		` :       ;                     '.      "''::::::''    .'             ` + "\n" +
		`    ;                           '.   \_              /               ` + "\n" +
		`  ;       ;                       +:   ~~--  ':'  -';                ` + "\n" +
		//		`      ;                            ':         : .::/                 ` + "\n" +
		`      ;                            ;;+_  :::. :..;;;                 ` + "\n" +
		`                                   ;;;;;;,;;;;;;;;,;                 ` + "\n",
}

//	Art: `    ,           ,   ` + "\n" +
//		`   /             \  ` + "\n" +
//		`  ((__-^^-,-^^-__)) ` + "\n" +
//		`   '-_---' '---_-'  ` + "\n" +
//		`    <__|o' 'o|__>   ` + "\n" +
//		`       \  '  /      ` + "\n" +
//		`        ): :(       ` + "\n" +
//		`        :o_o:       ` + "\n" +
//		`         "-"        ` + "\n",
