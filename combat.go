package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/art"
	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/simulations"
	. "github.com/eyetoe/foobarbaz/util"
)

// Fight loop where c is character and f is foe
func Fight(c *Agent, f *Agent) {
	ClearScreen()
	// calculate odd for the ascii art display only once
	odds := Odds(c, f)

	var charDamageOut, foeDamageOut, healmsg string

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

		//fmt.Println(Green(f.Art))
	}
	display()

	for {
		fmt.Printf(":> %sight, %svade, %sescribe, %sun\n<: ", GreenU("F"), GreenU("E"), GreenU("D"), GreenU("R"))
		choice, _, _ := GetChar()
		var done bool = false

		switch choice {
		// Fight
		case "f", "F":
			ClearScreen()
			display()

			fmt.Printf("\n%s attacks %s with %s.\n", c.Name, f.Name, c.Weap.Name)

			// Player Attacks First
			winner, loser, charAttackDetails := Attack(c, f)

			if c.Name == winner.Name {

				foeDamageOut = Damage(c, f, odds)

				if loser.Dead == true {
					if f.Weap.Name != c.Weap.Name && loser.DropChance >= Roll(1, 100) {
						OfferItem(c, f)
					}
					healmsg = WinHeal(c)
					done = true
				}
			}
			fmt.Printf("\n%s attacks %s with %s.\n", f.Name, c.Name, f.Weap.Name)

			// Foe Attacks Second
			winner, loser, foeAttackDetails := Attack(f, c)

			if f.Name == winner.Name {

				charDamageOut = Damage(f, c, odds)

				if loser.Dead == true {
					fmt.Printf("\n\n%s died.\n\n", c.Name)
					Continue()
					done = true
				}
			}
			ClearScreen()
			display()
			fmt.Printf(charAttackDetails)
			fmt.Println(foeDamageOut)
			foeDamageOut = ""
			if done == false {
				fmt.Printf(foeAttackDetails)
				fmt.Println(charDamageOut)
				charDamageOut = ""
			}
			fmt.Printf(healmsg)
			//Continue()
			if done == true {
				Continue()
				break
			} else {
				continue
			}

			// Evade
		case "e", "E":
			fmt.Println("Evade!\n You stall for time, looking for an opening!")
			fmt.Println("Note: should have a separate menu here where you can use items.  So you have to take a break from the fight, to be able to pull out an item or potion.  including switching weapons, drinking potions, using magic items.")
			continue
			// Describe the Foe
		case "d", "D":
			//fmt.Printf(Blue("\nYou consider the %s. %s\n"), f.Name, f.Description)
			f.Describe()
			FoeBar(*c, *f)
			fmt.Println()
			continue
			// Run from the fight
		case "r", "R":
			fmt.Println("Run!\n You have lost this battle, but may yet win the war!")
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
	ar := Roll(1, 100)
	dr := Roll(1, 100)
	// bonuses
	arB := a.Str.Val + a.Weap.Attack
	drB := d.Dex.Val
	// totals
	aT := ar + arB
	dT := dr + drB

	outText = fmt.Sprintf(Black("Attack roll: %d plus Bonus: %d for Total: %d\n"), ar, arB, aT)
	outText = outText + fmt.Sprintf(Black("Defence roll: %d plus Bonus: %d for Total: %d\n"), dr, drB, dT)

	// Attack wins if greater than Defence
	// But a tie goes to the Defence

	if aT > dT {
		outText = outText + fmt.Sprintf(Green("%s hits!\n "), a.Name)
		return a, d, outText
	} else {
		outText = outText + fmt.Sprintf(Red("%s misses!\n"), a.Name)
		return d, a, outText
	}

}

// Calculate and apply damage to Agent
func Damage(a *Agent, d *Agent, odds int) string {

	var textOut string

	hp := Roll(1, a.Weap.Damage)

	d.AdjHp(0 - hp)

	if d.Dead == false {
		textOut = textOut + fmt.Sprintf("for %s damage. ", Red(strconv.Itoa(hp)))
		textOut = textOut + fmt.Sprintf("%s's health = %s.\n", d.Name, Red(strconv.Itoa(d.Hp.Val)))
	}

	// Monster agents don't have a save file set
	if d.File == "" && d.Dead == true {

		// reverse the percentage
		mods := 100 - odds

		percentage := float32(mods) * .01

		// reduce the drop by the reverse percentage
		exp := float32(d.ExpDrop()) * percentage

		// exp is a float32 so do math back as in
		a.Exp = a.Exp + int(exp)

		textOut = textOut + fmt.Sprintf(Green("You gain %d experience.\n"), int(exp))
		a.Save()
	}
	d.Save()
	return textOut
}

func WinHeal(c *Agent) string {
	var textOut string
	h := Roll(2, c.MxHp.Val)
	if c.MxHp.Val > c.Hp.Val && c.MxHp.Val+30 >= Roll(1, 100) {
		c.AdjHp(h)
		textOut = textOut + fmt.Sprintf(Green("\nIn victory heal %d hit points!\n\n"), h)
		c.Save()
	}
	return textOut
}

func OfferItem(c *Agent, f *Agent) {
	ClearScreen()
	fmt.Println(Yellow(Sword1()))
	fmt.Println(Blue("!! ITEM DROP !!\n"))
	fmt.Printf("%s has dropped it's %s.\n\n", Yellow(f.Name), Yellow(f.Weap.Name))
	fmt.Printf("Replace?\n	")
	c.Weap.Display()
	fmt.Printf("with..\n	")
	f.Weap.Display()
	if Confirm("\nWould you like to make this swap?") == true {
		c.Weap = f.Weap
		c.Save()
	}
}

func Spawn() Agent {
	rand.Seed(time.Now().UTC().UnixNano())
	monsters := []Agent{
		// Add monsters here to be included in random spawn
		Spider,
		Phantom,
		Pixie,
		Kobold,
		Warlock,
		Coyote,
		Rogue,
		Minotaur,
		Lacrimosa,
		Griffon,
	}
	return monsters[rand.Intn(len(monsters))]

}
