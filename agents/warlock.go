package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

// Old Man
var Warlock = Agent{
	// Name and Description
	Name:        "Warlock",
	Description: "feeble and insane. He has a knowing glint in his eye.",
	// Stats
	Str: Stat{"Strength", 2},
	Int: Stat{"Intelligence", 48},
	Dex: Stat{"Dexterity", 4},
	// Health and Wellness
	MxHp: Stat{"Max Health", 25},
	Hp:   Stat{"Current Health", 25},
	Dead: false,
	// Equiped items
	Weap:  Staff,
	Armor: "Tattered Robes",
	Trink: "Black Cat",
	// Special Abilities
	Abl1: "Fireball",
	Abl2: "Shocking Touch",
	Abl3: "Heal",
	// Inventory
	DropChance: 20,
	Inv:        []Item{},
	Art: `Y .  . .:.. .. . ...::~={i}{nXYoqdWQQQgy@BRWQQQQmZWWRQmmgggwmQWR#QQQQQQQQQQQQQQQ` + "\n" +
		`:   .  .    ...-......:==?1suqmSdS3QQQQQQQmwwmQQ##ZoowWVVV$BQWWQQWQQQQQQWQQBRV#Q` + "\n" +
		`.  ..::........ :.::=auqumBRo3YYDQWQQQQQWQQQQQQQmmQqasaawa}?TVQQQQQQQQWmmXySS   ` + "\n" +
		`..c+:..........;:=%l*(:--;:-..- -?9$QQQWQWQQQQQQQ$*$WWWWBQQQQQmQ@5XQQZ!?Y|*     ` + "\n" +
		`....|<c:----.:-:...------ ..    . ..   +9QQQS1UWE1$E!dWQBQQQQQQQQBBUUpdY+|++++  ` + "\n" +
		`..;j><}-:...:..-. .                    . "9QWmgcioC:xQQW(#QQQBWQQQQW#5lii=|:.:  ` + "\n" +
		`.::=jmwx}<=;:::..: .  .     . . .            -!"++)|'=3dB+<n2dawyyQQW@V1i=++=--=` + "\n" +
		`-:=i>"+::il|::... - . .   .  .                  --:...:*(.=|??9VTT?^- .  ~--....` + "\n" +
		`.==:;:;:.::-:...... .. ...         .   . .       .=;:=;- ...  .     .   . .._.. ` + "\n" +
		`:=::- ...........::::::;=....__. _..           ..)Xgp;- .                   ..  ` + "\n" +
		`:... .:. . ...:::<siaa>)+)l|;_:=;=""<:. ...  . .:jQQQf'            ... .. . . ..` + "\n" +
		`=;._.-: . . ...:;<nZmQgw>=<symmdU|xs=i_::.:.....=qQQQm.             .-:  . ....-` + "\n" +
		`..:....   .   ..:++*1TRAa%==++9T|+-=-::=:;:;:.:<QWQQQ=... .     .     . . ..... ` + "\n" +
		`... ...:-.    . ...:::=xoxnis__..._saaS>++|+==++|VQQQQ==<i,  -...... .  .....:>-` + "\n" +
		`...;+~;--..       . ..:-==<+i3mXWVG)+===+ii|+===idQQQQE:==<|=<_=_=;:_=;. .. .=: ` + "\n" +
		`:-'.. -.-_,..   .    . ..-=i1mmmWnwi+=:==||||>=<i#QQQQ!:<,.:--=;x?!""'.. . .-- -` + "\n" +
		`..-- :._;;:.  .          .:={qmW$mp>==:---~+3IniumQQQQ(.=n1. . .  ...=;-..  ... ` + "\n" +
		`..:..+_:..   .        . .  ::=iq2c<:-    ... -+3$QQQQ@|.:;:<>:..._,:::;.  ... . ` + "\n" +
		`-. ;:""-:<_a... . . .     ..:=iI*^'.  .:;,_..-:)dQQQQE+:.:===|n%l==~.:. .... .  ` + "\n" +
		`.::-"+aa%TY::-:_.-..  .  . ..----.    .::~+=...:)$WRD1;._;;+=;:::. ... .....    ` + "\n" +
		`--""Y!~<s:=:+:a%%'..... . ..... :_;.     . .  . =+<|*|=:-.: :......   :::;|:..  ` + "\n" +
		`::;:+"!|:+%QgQVC.-=_=;-.. . .  =<ac..            ..---'........ .  ..:=/|s=.   .` + "\n" +
		`|=:<aa={3$EWQaswgmpo(.. . . _,swscx|:_a_..            .   .. . . .:_<;|"*~-. . -` + "\n" +
		`::==;|=awmWQBQmVVT~=-_xpuATTa2mkn?$(>3oC....        . -. .... ...::.::~+=       ` + "\n" +
		`:::~<>=1X5pXmWQmapisgwWmmD(wVQmn75u~ij^:_+-.:_:.  .. .._,..:-.=.==>::-.<='.     ` + "\n",
}
