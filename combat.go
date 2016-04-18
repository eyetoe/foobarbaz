package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/art"
	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/items"
	. "github.com/eyetoe/foobarbaz/simulations"
	. "github.com/eyetoe/foobarbaz/util"
)

// Fight loop where c is character and f is foe
func Fight(c *Agent, f *Agent) {
	c.FoeMaxHit = f.Weap.Damage

	// calculate odds for the ascii art display only once
	odds := Odds(c, f)

	// declare strings
	var charDamageOut, charAttackDetails, foeDamageOut, foeAttackDetails, healmsg, dropmsg string

	// declare attack winner
	var winner *Agent

	// display
	display := func() {
		c.StatusBar()
		fmt.Println()
		FoeBar(*c, *f)

		// Color Foe's name and show victory chance percentage
		switch {
		case odds >= 80:
			fmt.Printf(Green("%s\n"), f.Art)
		case odds >= 60:
			fmt.Printf(Cyan("%s\n"), f.Art)
		case odds >= 40:
			fmt.Printf(Blue("%s\n"), f.Art)
		case odds >= 20:
			fmt.Printf(Yellow("%s\n"), f.Art)
		case odds >= 0:
			fmt.Printf(Red("%s\n"), f.Art)
		}

	}

	// Fight loop start here
	for {

		// redraw and display stats
		ClearScreen()
		display()
		fmt.Printf(charAttackDetails)
		fmt.Printf(foeDamageOut)
		fmt.Printf(foeAttackDetails)
		fmt.Printf(charDamageOut)
		fmt.Printf(healmsg)
		fmt.Printf(dropmsg)

		// clear vars between runs
		charAttackDetails = ""
		foeDamageOut = ""
		foeAttackDetails = ""
		charDamageOut = ""
		healmsg = ""

		var hasWeapon, hasArmor, hasTrinket = false, false, false
		// check pulse
		if f.Dead == true {
			// check and only offer weapon is the character doesn't already have it
			if f.Weap.Name != c.Weap.Name && f.Weap.DropChance != Roll(1, 100) {
				for n, _ := range c.Inv {
					if f.Weap.Name == c.Inv[n].Name {
						hasWeapon = true
					}
				}
				if hasWeapon == false {
					OfferItem(c, f, f.Weap)
				}
			}
			// check and only offer armor is the character doesn't already have it
			if f.Armor.Name != c.Armor.Name && f.Armor.DropChance != Roll(1, 100) {
				for n, _ := range c.Inv {
					if f.Armor.Name == c.Inv[n].Name {
						hasArmor = true
					}
				}
				if hasArmor == false {
					OfferItem(c, f, f.Armor)
				}
			}
			// check and only offer trinket is the character doesn't already have it
			if f.Trink.Name != c.Trink.Name && f.Trink.DropChance != Roll(1, 100) {
				for n, _ := range c.Inv {
					if f.Trink.Name == c.Inv[n].Name {
						hasTrinket = true
					}
				}
				if hasTrinket == false {
					OfferItem(c, f, f.Trink)
				}
			}
			fmt.Printf(RedU("%s has died!\n"), f.Name)
			c.FoeMaxHit = 0
			c.Save()
			Continue()
			break
		}
		if c.Dead == true {
			fmt.Printf(RedU("%s has died!\n"), c.Name)
			c.FoeMaxHit = 0
			c.Save()
			Continue()
			break
		}

		numPotions := 0
		for _, c := range c.Inv {
			if c.Name == "Potion" {
				numPotions++
			}

		}

		// Combat Prompt
		fmt.Printf("\n:> %sight, ", GreenU("F"))
		fmt.Printf("%sse, ", GreenU("U"))
		fmt.Printf("%svade, ", GreenU("E"))
		fmt.Printf("%sescribe, ", GreenU("D"))
		fmt.Printf("%sotion(%s), ", GreenU("P"), Yellow(strconv.Itoa(numPotions)))
		fmt.Printf("%sun\n<: ", GreenU("R"))
		choice, _, _ := GetChar()

		switch choice {
		// Fight
		case "f", "F":
			ClearScreen()
			display()
			// Player Attacks First
			winner, _, charAttackDetails = Attack(c, f)
			if c.Name == winner.Name {
				foeDamageOut = Damage(c, f, odds)
				if f.Dead == true {
					healmsg = WinHeal(c)
					dropmsg = DropPotion(c)
					continue
				}
			}
			// Foe Attacks Second
			winner, _, foeAttackDetails = Attack(f, c)
			if f.Name == winner.Name {
				charDamageOut = Damage(f, c, odds)
			}
			// Evade
			continue
		// Event
		case "e", "E":
			fmt.Println("Evade!\n You stall for time, looking for an opening!")
			fmt.Println("Note: should have a separate menu here where you can use items.  So you have to take a break from the fight, to be able to pull out an item or potion.  including switching weapons, drinking potions, using magic items.")
			continue
		// Use Item
		case "u", "U":
			//fmt.Printf(Use(c, Potion))
			//Continue()
			continue
		// Describe Foe
		case "d", "D":
			//fmt.Printf(Blue("\nYou consider the %s. %s\n"), f.Name, f.Description)
			fmt.Println("Bacon bits")
			//c.Weap.Display()
			f.Describe()
			FoeBar(*c, *f)
			fmt.Println()
			continue
		// Use Potion
		case "p", "P":
			fmt.Printf(Use(c, Potion))
			Continue()
			continue
		// Run from the fight
		case "r", "R":
			fmt.Println("Run!\n You have lost this battle, but may yet win the war!")
			c.FoeMaxHit = 0
			c.Save()
			//Continue()
			break
		// Default back to loop
		default:
			continue
		}
		return
	}

}

// Attack takes two structs and returns two structs.
// First input struct is the 'attacker'
// Second input struct is the 'defender'
// First output struct is the 'winner'
// Second output struct is the 'loser'
func Attack(a *Agent, d *Agent) (*Agent, *Agent, string) {

	var outText string

	// roll for attacker and defender
	ar := Roll(2, 100)
	dr := Roll(2, 100)
	// bonuses
	arB := a.Str.Val + a.Weap.Attack
	// subtract the dodge percentage hit from armor
	drB := int(float64(d.Dex.Val) - float64(d.Dex.Val)*(float64(d.Armor.Dodge)*.01))

	// totals
	aT := ar + arB
	dT := dr + drB

	// uncomment here to 'show your work'
	//outText = fmt.Sprintf(Black("Attack roll: %d plus Bonus: %d for Total: %d\n"), ar, arB, aT)
	//outText = outText + fmt.Sprintf(Black("Defence roll: %d plus Bonus: %d for Total: %d\n"), dr, drB, dT)

	// Attack wins if greater than Defence
	// But a tie goes to the Defence
	if aT > dT {
		outText = outText + fmt.Sprintf(Green("%s hits with %s!\n "), a.Name, a.Weap.Name)
		return a, d, outText
	} else {
		outText = outText + fmt.Sprintf(Red("%s misses!\n"), a.Name)
		return d, a, outText
	}

}

// Calculate and apply damage to Agent
func Damage(a *Agent, d *Agent, odds int) string {

	var textOut string

	hp := Roll(2, a.Weap.Damage)

	d.AdjHp(0 - hp)

	//if d.Dead == false {
	textOut = textOut + fmt.Sprintf("for %s damage. ", Red(strconv.Itoa(hp)))
	textOut = textOut + fmt.Sprintf("%s's health = %s.\n", d.Name, Red(strconv.Itoa(d.Hp.Val)))
	//}

	// Monster agents don't have a save file set
	if d.File == "" && d.Dead == true {

		// reverse the percentage
		mods := 100 - odds

		percentage := float32(mods) * .01

		// reduce the drop by the reverse percentage
		exp := float32(d.ExpDrop()) * percentage

		// int to float conversion rounds down towards zero by dropping
		// everything after the decimal point. So I add 1 to the exp here
		// so the player never gets 0 exp reward
		exp++

		// exp is a float32 so do math with exp as an int
		a.Exp = a.Exp + int(exp)

		textOut = textOut + fmt.Sprintf(Green("\nYou gain %d experience.\n"), int(exp))
		a.Save()
	}
	d.Save()
	return textOut
}

func WinHeal(c *Agent) string {
	var textOut string
	h := Roll(2, c.MxHp.Val)
	if c.MxHp.Val > c.Hp.Val && c.MxHp.Val+30 >= Roll(2, 100) {
		c.AdjHp(h)
		textOut = textOut + fmt.Sprintf(Green("\nIn victory heal %d hit points!\n\n"), h)
		c.Save()
	}
	return textOut
}

func DropPotion(c *Agent) string {
	var textOut string
	//if c.MxHp.Val > c.Hp.Val && c.MxHp.Val+30 >= Roll(1, 100) {
	if c.MxHp.Val > c.Hp.Val && c.MxHp.Val+100 >= Roll(2, 100) {
		c.Inv = append(c.Inv, Potion)
		textOut = textOut + fmt.Sprintf("%s %s!\n\n", Green("You find a"), YellowU("potion"))
		c.Save()
	}
	return textOut
}

func OfferItem(c, f *Agent, i Item) {
	ClearScreen()
	fmt.Println(Yellow(Sword1()))
	fmt.Println(Blue("!! ITEM DROP !!\n"))
	//fmt.Printf("%s has dropped it's %s.\n\n", Yellow(f.Name), Yellow(f.Weap.Name))
	fmt.Printf("%s has dropped it's %s.\n\n", Yellow(f.Name), Yellow(i.Name))

	switch i.Slot {
	case "Weapon":
		fmt.Printf("Replace?\n")
		c.Weap.Display()
		fmt.Printf("with..\n")
		f.Weap.Display()
		if Confirm("\nWould you like to make this swap?") == true {
			c.Inv = append(c.Inv, c.Weap)
			c.Weap = f.Weap
			c.Save()
		} else {
			c.Inv = append(c.Inv, f.Weap)
			c.Save()
		}
	case "Armor":
		fmt.Printf("Replace?\n")
		c.Armor.Display()
		fmt.Printf("with..\n")
		f.Armor.Display()
		if Confirm("\nWould you like to make this swap?") == true {
			c.Inv = append(c.Inv, c.Armor)
			c.Armor = f.Armor
			c.Save()
		} else {
			c.Inv = append(c.Inv, f.Armor)
			c.Save()
		}
	case "Trinket":
		fmt.Printf("Replace?\n")
		c.Trink.Display()
		fmt.Printf("with..\n")
		f.Trink.Display()
		if Confirm("\nWould you like to make this swap?") == true {
			c.Inv = append(c.Inv, c.Trink)
			c.Trink = f.Trink
			c.Save()
		} else {
			c.Inv = append(c.Inv, f.Trink)
			c.Save()
		}
	}

}

func Spawn(c Agent) Agent {
	rand.Seed(time.Now().UTC().UnixNano())
	monsters := []Agent{
		// Add monsters here to be included in random spawn
		//		Spider,
		//		Phantom,
		Pixie,
		Kobold,
		Warlock,
		//		Coyote,
		Rogue,
		Minotaur,
		Lacrimosa,
		//		Griffon,
		//		Blackshuck,
		//		Goat,
		Drake,
		//		FlyingPig,
	}

	// candidate is a proposed foe.  The candidate is tested with the Odds()
	// function to see if it's got >5% and <95% chance to win combat.  If it
	// does we return that candidate.   We also cap the number of candidate
	// loops at 100, so that we don't keep looping forever looking for a foe
	// we will never find.
	var candidate Agent
	fmt.Println(Cyan("\n\n\nFinding opponent"))
	for i := 0; i < 100; i++ {

		fmt.Printf(".")
		candidate = monsters[rand.Intn(len(monsters))]
		candidateOdds := Odds(&c, &candidate)

		if candidateOdds < 95 && candidateOdds > 5 {
			return candidate
		}
	}
	// we didn't find a candidate after a 100 tries, so just return one
	return candidate
}
