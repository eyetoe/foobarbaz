package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
	var charDamageOut,
		charAttackDetails,
		foeDamageOut,
		foeAttackDetails,
		healmsg,
		dropmsg,
		potionOut string

	// declare attack winner
	var winner *Agent

	// display
	display := func() {
		c.StatusBar()
		fmt.Println()
		//FoeBar(*c, *f)

		if f.Dead == false && c.Dead == false {
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
		if f.Dead == true {

			victory := ` 		 _    ___      __                   __` + "\n" +
				`		| |  / (_)____/ /_____  _______  __/ /` + "\n" +
				`		| | / / / ___/ __/ __ \/ ___/ / / / / ` + "\n" +
				`		| |/ / / /__/ /_/ /_/ / /  / /_/ /_/  ` + "\n" +
				`		|___/_/\___/\__/\____/_/   \__, (_)   ` + "\n" +
				`		                          /____/      ` + "\n\n"

			fmt.Printf("%s", Green(victory))
		}
		if c.Dead == true {

			defeat := `		    ____       ____           __  __ ` + "\n" +
				`		   / __ \___  / __/__  ____ _/ /_/ / ` + "\n" +
				`		  / / / / _ \/ /_/ _ \/ __ '/ __/ /  ` + "\n" +
				`		 / /_/ /  __/ __/  __/ /_/ / /_/_/   ` + "\n" +
				`		/_____/\___/_/  \___/\__,_/\__(_)    ` + "\n\n"

			fmt.Printf("%s", Red(defeat))
		}

		FoeBar(*c, *f)
		Meter(f.Health.Val, f.MaxHealth.Val, c.Weap.Damage, "Health", "█", "foe")
		fmt.Println()
		//░▒█░   ░ ████▓▒░░ ████▓▒░░▓█  ▀█▓ ▓█   ▓██▒░██▓ ▒██▒░▓█  ▀█▓ ▓█   ▓██▒███████▒

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
		fmt.Printf(potionOut)
		fmt.Printf(healmsg)
		fmt.Printf(dropmsg)

		// clear vars between runs
		charAttackDetails = ""
		foeDamageOut = ""
		foeAttackDetails = ""
		charDamageOut = ""
		potionOut = ""
		healmsg = ""
		dropmsg = ""

		var hasWeapon, hasArmor, hasTrinket = false, false, false
		// check pulse
		if f.Dead == true {
			// check and only offer weapon is the character doesn't already have it
			if f.Weap.Name != c.Weap.Name && f.Weap.DropChance >= Roll(1, 100) {
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
			if f.Armor.Name != c.Armor.Name && f.Armor.DropChance >= Roll(1, 100) {
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
			if f.Trink.Name != c.Trink.Name && f.Trink.DropChance >= Roll(1, 100) {
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
		fmt.Printf("%sealth Potion(%s), ", GreenU("H"), Yellow(strconv.Itoa(numPotions)))
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
				foeDamageOut = DoDamage(c, f, odds)
				if f.Dead == true {
					healmsg = WinHeal(c)
					dropmsg = DropPotion(c)
					continue
				}
			}
			// Foe Attacks Second
			winner, _, foeAttackDetails = Attack(f, c)
			if f.Name == winner.Name {
				charDamageOut = DoDamage(f, c, odds)
			}
			continue
		// Evade
		case "e", "E":
			for _, each := range c.Equipped() {
				fmt.Println(each)
			}
			Continue()
			fmt.Println("Evade!\n You stall for time, looking for an opening!")
			fmt.Println("Note: should have a separate menu here where you can use items.  So you have to take a break from the fight, to be able to pull out an item or potion.  including switching weapons, drinking potions, using magic items.")
			continue
		// Use Item
		case "u", "U":
			continue
		// Consider Foe
		case "d", "D":
			f.Consider()
			FoeBar(*c, *f)
			fmt.Println()
			continue
		// Use Potion
		case "h", "H":
			potionOut = Use(c, Potion)
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
	//arB := a.Str.Val + a.Weap.Attack
	arB := a.TotalAttack()
	// subtract the dodge percentage hit from armor
	//drB := int(float64(d.Dex.Val) - float64(d.Dex.Val)*(float64(d.Armor.Dodge)*.01))
	drB := d.TotalDodge()

	// totals
	aT := ar + arB
	dT := dr + drB

	// uncomment here to 'show your work'
	//outText = fmt.Sprintf(Black("Attack roll: %d plus Bonus: %d for Total: %d\n"), ar, arB, aT)
	//outText = outText + fmt.Sprintf(Black("Resist roll: %d plus Bonus: %d for Total: %d\n"), dr, drB, dT)

	// Attack wins if greater than Resist
	// But a tie goes to the Resist
	if aT > dT {
		outText = outText + fmt.Sprintf(Green("%s hits with %s!\n "), a.Name, a.Weap.Name)
		return a, d, outText
	} else {
		outText = outText + fmt.Sprintf(Red("%s misses!\n"), a.Name)
		return d, a, outText
	}

}

// Calculate and apply damage to Agent
func DoDamage(a *Agent, d *Agent, odds int) string {

	var damage int
	var textOut string

	// Damage and max damage if it's a critical
	if Roll(1, 100) > a.TotalCritical() {
		damage = Roll(2, a.TotalDamage())
	} else {
		damage = a.TotalDamage()
		textOut = fmt.Sprintf(CyanU("Critical") + " ")
	}

	// If damage is greater than the damage resist then subtract
	if damage > d.TotalResist() {
		damage = damage - d.TotalResist()
		//if unlocked, damage = damage - TotalResist()
		d.AdjHealth(0 - damage)
		textOut = textOut + fmt.Sprintf("for %s damage. ", Red(strconv.Itoa(damage)))
		textOut = textOut + fmt.Sprintf("%s's health = %s.\n", d.Name, Red(strconv.Itoa(d.Health.Val)))
		//else don't adjust
	} else {
		damage = 0
		textOut = textOut + fmt.Sprintf("%s! for %s damage. ", YellowU("Resist"), Red(strconv.Itoa(damage)))
		textOut = textOut + fmt.Sprintf("%s's health = %s.\n", d.Name, Red(strconv.Itoa(d.Health.Val)))
	}

	// Experience Reward
	// if this is a dead monster
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
	if c.MaxHealth.Val != c.Health.Val {
		h := Roll(1, (c.MaxHealth.Val - c.Health.Val))
		if c.MaxHealth.Val > c.Health.Val && c.MaxHealth.Val+50 >= Roll(2, 100) {
			c.AdjHealth(h)
			textOut = textOut + fmt.Sprintf(Green("\nIn victory heal %d hit points!\n\n"), h)
			c.Save()
		}
	}
	return textOut
}

func DropPotion(c *Agent) string {
	var textOut string
	potionDropChance := 100
	if c.MaxHealth.Val > c.Health.Val && c.MaxHealth.Val+potionDropChance >= Roll(2, 100) {
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
	fmt.Printf("%s has dropped it's %s.\n\n", Yellow(f.Name), Yellow(i.Name))

	var oldodds, newodds int
	var x Agent
	x = *c

	switch i.Slot {
	case "Weapon":
		x.Weap = i
		oldodds = Odds(c, &x)
		newodds = Odds(&x, c)
		fmt.Printf("Replace?\n")
		fmt.Printf("%d%% vs.", oldodds)
		c.Weap.Display()
		fmt.Printf("with..\n")
		fmt.Printf("%d%% vs.", newodds)
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
		x.Armor = i
		oldodds = Odds(c, &x)
		newodds = Odds(&x, c)
		fmt.Printf("Replace?\n")
		fmt.Printf("%d%% vs.", oldodds)
		c.Armor.Display()
		fmt.Printf("with..\n")
		fmt.Printf("%d%% vs.", newodds)
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
		x.Trink = i
		oldodds = Odds(c, &x)
		newodds = Odds(&x, c)
		fmt.Printf("Replace?\n")
		fmt.Printf("%d%% vs.", oldodds)
		c.Trink.Display()
		fmt.Printf("with..\n")
		fmt.Printf("%d%% vs.", newodds)
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

func SpawnChooser(c *Agent) *Agent {
	//func SpawnChooser(c *Agent) {
	monsters := []Agent{
		//Add monsters here to be included in random spawn
		Kobold,
		Spider,
		Phantom,
		Coyote,
		Pixie,
		Warlock,
		Rogue,
		Blackshuck,
		Minotaur,
		Griffon,
		Lacrimosa,
		Drake,
		FlyingPig,
		Goat,
		FreshZombie,
		Swordsman,
		MakeMonster(c),
	}
	fmt.Println()
	for {

		for number, each := range monsters {
			fmt.Println(number, each.Name)
		}
		reader := bufio.NewReader(os.Stdin)
		// Prompt and read
		fmt.Print("Choose number of monster to spawn: ")
		choice, _ := reader.ReadString('\n')

		choice = strings.Trim(choice, "\n")

		if intchoice, err := strconv.Atoi(choice); err == nil {
			if intchoice <= len(monsters) {

				return &monsters[intchoice]
				break
			}
		} else {
			continue
		}

	}

	return c
}

func Spawn(c Agent) Agent {
	rand.Seed(time.Now().UTC().UnixNano())
	monsters := []Agent{
		//Add monsters here to be included in random spawn
		Kobold,
		Spider,
		Phantom,
		Coyote,
		Pixie,
		Warlock,
		Rogue,
		Blackshuck,
		Minotaur,
		Griffon,
		Lacrimosa,
		Drake,
		FlyingPig,
		Goat,
		FreshZombie,
		Swordsman,
		//Blob, // blob dosn't work, dynamic stats are persistent until game is reloaded
		//  so the MakeMonster generator is much better.
		//		MakeMonster(&c),
		//		MakeMonster(&c),
		MakeMonster(&c),
		MakeMonster(&c),
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

		if candidateOdds < 95 && candidateOdds > 30 {
			return candidate
		}
	}
	// we didn't find a candidate after a 100 tries, so just return one
	return candidate
}
