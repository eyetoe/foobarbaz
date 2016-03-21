package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/colors"
)

func Prompt() {
	for {
		// setting File var in Agent struct, this allows Char.Load to pick up
		// the save file without prompting.
		Char := Agent{File: "Izro"}
		Char.Load()
		ClearScreen()
		Char.StatusBar()
		Char.Save()
		fmt.Printf("\n%s <- You are here.\n", BlueU(Char.Loc.Name))

		// ask the first question
		//fmt.Printf(":> %sook, %so, %sharacter, %svent, %selp <: ", GreenU("L"), GreenU("G"), GreenU("C"), GreenU("E"), GreenU("H"))
		fmt.Printf(":> %sook, ", GreenU("L"))
		fmt.Printf("%so, ", GreenU("G"))
		fmt.Printf("%sharacter, ", GreenU("C"))
		fmt.Printf("%svent, ", GreenU("E"))
		fmt.Printf("%sreferences, ", GreenU("P"))
		fmt.Printf("%selp <: ", GreenU("H"))

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := string([]byte(input)[0])

		// choices here
		switch choice {
		case "l", "L":
			Look(Char)
			continue
		case "g", "G":
			Go()
			break
		case "c", "C":
			Character(&Char)
			continue
		case "e", "E":
			var Foe Agent
			//Foe = Phantom
			Foe = Spawn()
			Fight(&Char, &Foe)
			continue
		case "p", "P":
			Preferences(&Char)
			continue
		case "h", "H":
			fmt.Println(Blue("All you help are belong to us."))
		}
	}
}

func Character(c *Agent) {
	for {
		ClearScreen()
		c.StatusBar()

		fmt.Printf("\n%s \n:> ", BlueU("Character"))
		fmt.Printf("%sest, ", Green("R"))
		fmt.Printf("%sp, ", Green("X"))
		fmt.Printf("%sbility, ", Green("A"))
		fmt.Printf("%snventory, ", Green("I"))
		fmt.Printf("%saves, ", Green("S"))
		fmt.Printf("%sack <:", Green("B"))

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := string([]byte(input)[0])
		switch choice {
		case "r", "R":
			fmt.Printf(Blue("\nYou rest and recuperate.\n"))
			if c.Hp.Val < c.MxHp.Val/2 {
				c.Hp.Val = c.MxHp.Val / 2
				fmt.Printf(Green("You heal.\n"))
				c.Save()
			}
			Continue()
			continue
		case "x", "X":
			ExpMgr(c)
			continue
		case "a", "A":
			fmt.Printf(Blue("\nAbility 1: %s.\n"), c.Abl1)
			fmt.Printf(Blue("Ability 2: %s.\n"), c.Abl2)
			fmt.Printf(Blue("Ability 3: %s.\n"), c.Abl3)
			continue
		case "i", "I":
			fmt.Printf(Blue(c.Inv))
			continue
		case "s", "S":
			SaveMgr(c)
			continue
		case "b", "B":
			return
		}
	}
}

func SaveMgr(c *Agent) {
	for {
		fmt.Printf("\n%s \n:> ", BlueU("Save File Management"))
		fmt.Printf("%sew, ", Green("N"))
		fmt.Printf("%save, ", Green("S"))
		fmt.Printf("%soad, ", Green("L"))
		fmt.Printf("%sack <:", Green("B"))

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := string([]byte(input)[0])

		switch choice {
		case "n", "N":
			fmt.Printf("\ncomming soon, New Character\n")
			continue
		case "s", "S":
			c.Save()
			continue
		case "l", "L":
			c.Load()
			continue
		case "b", "B":
			return
		}
	}

}

func ExpMgr(c *Agent) {
	for {
		ClearScreen()
		c.StatusBar()
		cost := StatCost(c)
		fmt.Printf("\n%s \n", YellowU("Experience Point Store!\n"))
		fmt.Printf(Blue("You have %s %s\n"), Green(strconv.Itoa(c.Exp)), Blue("experience.\n"))
		fmt.Printf(Blue("Choose a Stat to increase 1 point\n"))
		fmt.Printf(":>	%strength	 (%s xp),\n", Green("S"), Yellow(strconv.Itoa(cost)))
		fmt.Printf("	%sntelligence	 (%s xp),\n", Green("I"), Yellow(strconv.Itoa(cost)))
		fmt.Printf("	%sexterity	 (%s xp),\n", Green("D"), Yellow(strconv.Itoa(cost)))
		fmt.Printf("	%sP		 (%s xp),\n", Green("H"), Yellow(strconv.Itoa(cost)))
		fmt.Printf("	%sack\n", Green("B"))
		fmt.Printf("<:")

		consolation := func() {
			fmt.Printf(Yellow("\nAlas you have not the required experience.\n"))
			fmt.Printf(Red("Without risk there can be no reward.\n"))
			Continue()
		}

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := string([]byte(input)[0])

		switch choice {
		case "s", "S":
			if c.Exp >= cost {
				c.Exp = c.Exp - cost
				c.Str.Val = c.Str.Val + 1
				fmt.Println(Blue("\nYour muscles quiver and flex."))
				fmt.Printf(Yellow("\n Strength is now %s.\n"), Green(strconv.Itoa(c.Str.Val)))
				c.Save()
				Continue()
			} else {
				consolation()
			}

			continue
		case "i", "I":
			if c.Exp >= cost {
				c.Exp = c.Exp - cost
				c.Int.Val = c.Int.Val + 1
				fmt.Println(Blue("\nYou have an epiphany of significant import."))
				fmt.Printf(Yellow("\n Intelligence is now %s.\n"), Green(strconv.Itoa(c.Int.Val)))
				c.Save()
				Continue()
			} else {
				consolation()
			}

			continue
		case "d", "D":
			if c.Exp >= cost {
				c.Exp = c.Exp - cost
				c.Dex.Val = c.Dex.Val + 1
				fmt.Println(Blue("\nYour nerves tingle and radiate with anticipation."))
				fmt.Printf(Yellow("\n Dexterity is now %s.\n"), Green(strconv.Itoa(c.Dex.Val)))
				c.Save()
				Continue()
			} else {
				consolation()
			}

			continue
		case "h", "H":
			if c.Exp >= cost {
				c.Exp = c.Exp - cost
				c.MxHp.Val = c.MxHp.Val + 1
				fmt.Println(Blue("\nAttention to physical health improves your metabolism."))
				fmt.Printf(Yellow("\n Max Hp is now %s.\n"), Green(strconv.Itoa(c.MxHp.Val)))
				c.Save()
				Continue()
			} else {
				consolation()
			}

			continue
		case "b", "B":
			return
		}
	}

}

func Go() {
	fmt.Printf(":> %sxit, %sack <:", Green("E"), Green("B"))
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	choice := string([]byte(input)[0])
	switch choice {
	case "e", "E":
		os.Exit(0)
	case "b", "B":
		return
	}
}
func Look(c Agent) {
	fmt.Printf(Blue("You are here: %s\n"), c.Loc.Name)
	fmt.Printf(Blue("%s\n"), c.Loc.Description)
	Continue()
}

// Fight loop where c is character and f is foe
func Fight(c *Agent, f *Agent) {
	ClearScreen()
	fmt.Printf("\nYou have encountered a %s!\n", YellowU("Monster"))
	Continue()
	ClearScreen()
	for {
		c.StatusBar()
		fmt.Printf("\n%s\n:> %sight, %svade, %sescribe, %sun\n<: ", RedU(f.Name), GreenU("F"), GreenU("E"), GreenU("D"), GreenU("R"))
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := string([]byte(input)[0])

		switch choice {
		case "f", "F":
			fmt.Printf("\n%s attacks %s with %s.\n", c.Name, f.Name, c.Weap.Name)
			winner, loser := Attack(c, f)
			if c.Name == winner.Name {
				Damage(c, f)
				if loser.Dead == true {
					Continue()
					break
				}
			}
			fmt.Printf("\n%s attacks %s with %s.\n", f.Name, c.Name, f.Weap.Name)
			winner, loser = Attack(f, c)
			if f.Name == winner.Name {
				Damage(f, c)
				if loser.Dead == true {
					Continue()
					break
				}
			}
			continue
		case "e", "E":
			fmt.Println("Evade!\n You stall for time, looking for an opening!")
			continue
		case "d", "D":
			fmt.Printf(Blue("\nYou consider the %s. %s\n"), f.Name, f.Description)
			f.FoeBar()
			fmt.Println()
			continue
		case "r", "R":
			fmt.Println("Run!\n You have lost this battle, but may yet win the war!")
			break
		default:
			Fight(c, f)
			break
		}
		return
	}
}

func Preferences(c *Agent) {
	fmt.Println(Yellow("\nArrr, here be ye preferences..."), BlueU(c.Name), "\n")
	Continue()
}

func Continue() {
	fmt.Printf(BlueU("\nEnter"))
	fmt.Printf(Blue(" to continue...\n"))
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	choice := string([]byte(input)[0])
	switch choice {
	default:
		return
	}
}

func ClearScreen() {
	fmt.Printf("[H[J")
}

func Banner() {

	ClearScreen()
	fmt.Printf(White("\nWelcome to ...\n"))
	fmt.Println(Red(
		`   ______         ______           ______              ` + "\n" +
			`   |  ___|        | ___ \          | ___ \             ` + "\n" +
			`   | |_ ___   ___ | |_/ / __ _ _ __| |_/ / __ _ ____   ` + "\n" +
			`   |  _/ _ \ / _ \| ___ \/ _' | '__| ___ \/ _' |_  /   ` + "\n" +
			`   | || (_) | (_) | |_/ / (_| | |  | |_/ / (_| |/ /    ` + "\n" +
			`   \_| \___/ \___/\____/ \__,_|_|  \____/ \__,_/___|   ` + "\n" +
			`                                                       `))
}
