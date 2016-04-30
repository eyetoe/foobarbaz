package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

var FreshZombie = Agent{
	// Name and Description
	Name:        "Fresh Zombie",
	Description: "... Oh no!  Oh dear God no! The smell! Oh God the smell.",
	// Stats
	Str: Stat{"Strength", 70},
	Int: Stat{"Intelligence", 16},
	Dex: Stat{"Dexterity", 14},
	// Health and Wellness
	MaxHealth: Stat{"Max Health", 100},
	Health:    Stat{"Current Health", 100},
	Dead:      false,
	// Equipped items
	Weap:  ZombieBite,
	Armor: Empty,
	Trink: Empty,
	// Special Abilities
	Abl1: "",
	Abl2: "",
	Abl3: "",
	// Inventory
	Inv: []Item{},
	Art: `                      :::!~!!!!!:.          ` + "\n" +
		`                  .xUHWH!! !!?M88WHX:.      ` + "\n" +
		`                .X*#M@$!!  !X!M$$$$$$WWx:.  ` + "\n" +
		`               :!!!!!!?H! :!$!$$$$$$$$$$8X: ` + "\n" +
		`              !!~  ~:~!! :~!$!#$$$$$$$$$$8X:` + "\n" +
		`             :!~::!H!<   ~.U$X!?R$$$$$$$$MM!` + "\n" +
		`             ~!~!!!!~~ .:XW$$$U!!?$$$$$$RMM!` + "\n" +
		`               !:~~~ .:!M"T#$$$$WX??#MRRMMM!` + "\n" +
		`               ~?WuxiW*'   '"#$$$$8!!!!??!!!` + "\n" +
		`             :X- M$$$$       '"T#$T~!8$WUXU~` + "\n" +
		`            :%'  ~#$$$m:        ~!~ ?$$$$$$ ` + "\n" +
		`          :!'.-   ~T$$$$8xx.  .xWW- ~""##*" ` + "\n" +
		`.....   -~~:<' !    ~?T#$$@@W@*?$$      /'  ` + "\n" +
		`W$@@M!!! .!~~ !!     .:XUW$W!~ '"~:    :    ` + "\n" +
		`#"~~'.:x%'!!  !H:   !WM$$$$Ti.: .!WUn+!'    ` + "\n" +
		`:::~:!!':X~ .: ?H.!u "$$$B$$$!W:U!T$$M~     ` + "\n" +
		`.~~   :X@!.-~   ?@WTWo("*$$$W$TH$! '        ` + "\n" +
		`Wi.~!X$?!-~    : ?$$$B$Wu("**$RM!           ` + "\n" +
		`$R@i.~~ !     :   ~$$$$$B$$en:''            ` + "\n" +
		`?MXT@Wx.~    :     ~"##*$$$$M~              ` + "\n",
}
