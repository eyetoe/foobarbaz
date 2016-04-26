package items

import (
	"fmt"
	"strconv"

	"github.com/eyetoe/foobarbaz/affects"
	. "github.com/eyetoe/foobarbaz/colors"
)

type Item struct {
	Name        string
	Description string
	Slot        string
	Affects     []affects.Affect
	// modify 'To Hit' chance
	Attack int
	// modify damage dealt if hit
	Damage int
	// Passive Resist is damage resistance
	Resist int
	// A measure of how impared agent is from armor.  This int is the subtracted
	// percentage adjusted from Dex
	Dodge int
	// dammage over time adjustment. Bleed, Poison, Fire, etc
	DoT int
	// modify critical chance.
	Crit int
	// chance to drop from foe
	DropChance int
}

func (i *Item) Display() {
	//fmt.Printf("%s\n", Blue(i.Name))
	//fmt.Printf("	Slot: %s, Attack: %s, Damage: %s\n", Yellow(i.Slot), Red(strconv.Itoa(i.Attack)), Red(strconv.Itoa(i.Damage)))

	// Prety print var strings if defined
	if i.Name != "" {
		fmt.Printf("  %s: ", YellowU(i.Name))
	}
	if i.Slot != "" {
		fmt.Printf("%s\n", i.Slot)
	}
	if i.Description != "" {
		fmt.Printf("  %s %s\n", Cyan(i.Name), Cyan(i.Description))
	}

	// Prety print var ints if defined
	if i.Attack != 0 {
		fmt.Printf("	Attack:	+%s, \n", Yellow(strconv.Itoa(i.Attack)))
	}
	if i.Damage != 0 {
		fmt.Printf("	Damage:	1d%s, \n", Yellow(strconv.Itoa(i.Damage)))
	}
	if i.Resist != 0 {
		fmt.Printf("	Resist:	+%s, \n", Yellow(strconv.Itoa(i.Resist)))
	}
	// display dodge if it's an armor piece
	if i.Slot == "Armor" {
		fmt.Printf("	Dodge:  	-%s%%, \n", Yellow(strconv.Itoa(i.Dodge)))
	}
	if i.DoT != 0 {
		fmt.Printf("	DoT:	%s, \n", Yellow(strconv.Itoa(i.DoT)))
	}
	if i.Crit != 0 {
		fmt.Printf("	Crit:	+%s%%, \n", Yellow(strconv.Itoa(i.Crit)))
	}

	// Prety print var Affects if defined
	if &i.Affects != nil {
		fmt.Printf("	Affects: ")
		for _, e := range i.Affects {
			fmt.Printf("%s, ", Yellow(e.Name))

		}
		fmt.Printf("\n")
	}
}
