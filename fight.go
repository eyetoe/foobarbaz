package main

import (
	"fmt"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/colors"
)

// Fight loop where c is character and f is foe
func Fight(c *Agent, f *Agent) {
	//ClearScreen()
	//fmt.Printf("\nYou have encountered a %s!\n", YellowU("Monster"))
	//Continue()
	ClearScreen()

	var charDamageOut, foeDamageOut, healmsg string

	display := func() {
		c.StatusBar()
		fmt.Println()
		f.FoeBar()
		fmt.Println(Green(f.Art))
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

				foeDamageOut = Damage(c, f)

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

			//fmt.Println("debug f.Name,  winner.Name, and done:", f.Name, winner.Name, done)
			//Continue()
			if f.Name == winner.Name {

				charDamageOut = Damage(f, c)

				if loser.Dead == true {
					//Continue()
					done = true
				}
			}
			ClearScreen()
			display()
			fmt.Printf(charAttackDetails)
			fmt.Printf(foeDamageOut)
			if done == false {
				fmt.Printf(foeAttackDetails)
				fmt.Printf(charDamageOut)
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
			fmt.Printf(Blue("\nYou consider the %s. %s\n"), f.Name, f.Description)
			f.FoeBar()
			fmt.Println()
			continue
			// Run from the fight
		case "r", "R":
			fmt.Println("Run!\n You have lost this battle, but may yet win the war!")
			break
			// Default back to loop
		default:
			//Fight(c, f)
			//break
			continue
		}
		return
	}
}
