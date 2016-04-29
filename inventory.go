package main

import (
	"fmt"
	"strconv"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/items"
	//	. "github.com/eyetoe/foobarbaz/items"
	. "github.com/eyetoe/foobarbaz/util"
)

func Inventory(c *Agent) {

	ClearScreen()
	c.StatusBar()

	//	fmt.Println(Yellow("\n" + `                            _.--.         ` + "\n" +
	//		`                        _.-'_:-'||        ` + "\n" +
	//		`                    _.-'_.-::::'||        ` + "\n" +
	//		`               _.-:'_.-::::::'  ||        ` + "\n" +
	//		`             .''-.-:::::::'     ||        ` + "\n" +
	//		`            /.'';|:::::::'      ||_       ` + "\n" +
	//		`           ||   ||::::::'     _.;._'-._   ` + "\n" +
	//		`           ||   ||:::::'  _.-!oo @.!-._'-.` + "\n" +
	//		`           \'.  ||:::::.-!()oo @!()@.-'_.|` + "\n" +
	//		`            '.'-;|:.-'.&$@.& ()$%-'o.'\U||` + "\n" +
	//		`              '>'-.!@%()@'@_%-'_.-o _.|'||` + "\n" +
	//		`               ||-._'-.@.-'_.-' _.-o  |'||` + "\n" +
	//		`               ||=[ '-._.-\U/.-'    o |'||` + "\n" +
	//		`               || '-.]=|| |'|      o  |'||` + "\n" +
	//		`               ||      || |'|        _| ';` + "\n" +
	//		`               ||      || |'|    _.-'_.-' ` + "\n" +
	//		`               |'-._   || |'|_.-'_.-'     ` + "\n" +
	//		`            jgs '-._'-.|| |' '_.-'        ` + "\n" +
	//		`                    '-.||_/.-'            ` + "\n"))

	//c.Inv = append(c.Inv, Hatchet)
	//c.Inv = append(c.Inv, Knife)
	//c.Inv = append(c.Inv, Dagger)
	//c.Inv = append(c.Inv, Staff)
	//c.Inv = append(c.Inv, Tack)

	fmt.Println(YellowU("\n\nInventory:\n"))

	//	for a, b := range c.Inv {
	//		//fmt.Println(GreenU(a), WhiteU(b.Name))
	//		fmt.Printf("%s:	%s %s Attack: +%s Damage: 1d%s\n", GreenU(strconv.Itoa(a)), WhiteU(b.Name), Blue(b.Slot), Red(strconv.Itoa(b.Attack)), Red(strconv.Itoa(b.Damage)))
	//		fmt.Println("	", Blue(b.Description), "\n")
	//
	//	}

	numPotions := 0
	var allWeapons []Item
	var allArmor []Item
	var allTrinkets []Item
	for _, b := range c.Inv {
		if b.Name == "Potion" {
			numPotions++
		}
		if b.Slot == "Weapon" {
			allWeapons = append(allWeapons, b)
		}
		if b.Slot == "Armor" {
			allArmor = append(allArmor, b)
		}
		if b.Slot == "Trinket" {
			allTrinkets = append(allTrinkets, b)
		}

	}
	fmt.Println(YellowU("Equipped:\n"))
	fmt.Printf(":	%s %s Attack: +%s Damage: 1d%s\n", WhiteU(c.Weap.Name), Blue(c.Weap.Slot), Red(strconv.Itoa(c.Weap.Attack)), Red(strconv.Itoa(c.Weap.Damage)))
	fmt.Println("	", Blue(c.Weap.Description))
	fmt.Printf(":	%s %s Attack: +%s Damage: 1d%s\n", WhiteU(c.Armor.Name), Blue(c.Armor.Slot), Red(strconv.Itoa(c.Armor.Attack)), Red(strconv.Itoa(c.Armor.Damage)))
	fmt.Println("	", Blue(c.Armor.Description))
	fmt.Printf(":	%s %s Attack: +%s Damage: 1d%s\n", WhiteU(c.Trink.Name), Blue(c.Trink.Slot), Red(strconv.Itoa(c.Trink.Attack)), Red(strconv.Itoa(c.Trink.Damage)))
	fmt.Println("	", Blue(c.Trink.Description))

	fmt.Printf("(%s) %s:\n  %s\n\n", Cyan(strconv.Itoa(numPotions)), Yellow(Potion.Name), Cyan(Potion.Description))

	fmt.Println(YellowU("Weapons:\n"))
	for n, t := range allWeapons {
		fmt.Printf("%s:	%s %s Attack: +%s Damage: 1d%s\n", GreenU(strconv.Itoa(n)), WhiteU(t.Name), Blue(t.Slot), Red(strconv.Itoa(t.Attack)), Red(strconv.Itoa(t.Damage)))
		fmt.Println("	", Blue(t.Description))

	}

	fmt.Println(YellowU("Armor:\n"))
	for n, t := range allArmor {
		fmt.Printf("%s:	%s %s Attack: +%s Damage: 1d%s\n", GreenU(strconv.Itoa(n)), WhiteU(t.Name), Blue(t.Slot), Red(strconv.Itoa(t.Attack)), Red(strconv.Itoa(t.Damage)))
		fmt.Println("	", Blue(t.Description))

	}

	fmt.Println(YellowU("Trinkets:\n"))
	for n, t := range allTrinkets {
		fmt.Printf("%s:	%s %s Attack: +%s Damage: 1d%s\n", GreenU(strconv.Itoa(n)), WhiteU(t.Name), Blue(t.Slot), Red(strconv.Itoa(t.Attack)), Red(strconv.Itoa(t.Damage)))
		fmt.Println("	", Blue(t.Description))

	}

	Continue()

}
