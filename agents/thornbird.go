package agents

import (
	. "github.com/eyetoe/foobarbaz/items"
)

var Thornbird = Agent{
	// Name and Description
	Name:        "Thornbird",
	Description: "ravenous giant bird of prey.",
	// Stats
	Str: Stat{"Strength", 30},
	Int: Stat{"Intelligence", 10},
	Dex: Stat{"Dexterity", 80},
	// Health and Wellness
	MaxHealth: Stat{"Max Health", 88},
	Health:    Stat{"Current Health", 88},
	Dead:      false,
	// Equipped items
	Weap:  RavenousBeak,
	Armor: Empty,
	Trink: Empty,
	// Special Abilities
	Abl1: "",
	Abl2: "",
	Abl3: "",
	// Inventory
	Inv: []Item{},
	Art: `                            \                  /             ` + "\n" +
		`                   _________))                ((__________    ` + "\n" +
		`                  /.-------./\\    \    /    //\.--------.\   ` + "\n" +
		`                 //#######//##\\   ))  ((   //##\\########\\  ` + "\n" +
		`                //#######//###((  ((    ))  ))###\\########\\ ` + "\n" +
		`               ((#######((#####\\  \\  //  //#####))########))` + "\n" +
		`                \##' '###\######\\  \)(/  //######/####' '##/ ` + "\n" +
		`                 )'    ''#)'  '##\'->oo<-'/##'  '(#''     '(  ` + "\n" +
		`                         (       ''\'..'/''       )           ` + "\n" +
		`                                    \""(                      ` + "\n" +
		`                                     '- )                     ` + "\n" +
		`                                     / /                      ` + "\n" +
		`                                    ( /\                      ` + "\n" +
		`                                    /\| \                     ` + "\n" +
		`                                   (  \                       ` + "\n" +
		`                                       )                      ` + "\n" +
		`                                      /                       ` + "\n" +
		`                                     (                        ` + "\n" +
		`                                     '   Ed Zahurak           ` + "\n",
}
