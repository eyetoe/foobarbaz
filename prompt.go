package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/art"
	. "github.com/eyetoe/foobarbaz/colors"
)

func Prompt() {
	for {
		// setting File var in Agent struct, this allows Char.Load to pick up
		// the save file without prompting.
		Char := Agent{File: SaveFile}
		Char.Load()
		ClearScreen()
		Char.StatusBar()
		Char.Save()
		fmt.Printf("\n%s <- You are here.\n", BlueU(Char.Loc.Name))

		// ask the first question
		fmt.Printf(":> %sook, ", GreenU("L"))
		fmt.Printf("%so, ", GreenU("G"))
		fmt.Printf("%svent, ", GreenU("E"))
		fmt.Printf("%sharacter, ", GreenU("C"))
		fmt.Printf("%srefs, ", GreenU("P"))
		fmt.Printf("%sile, ", GreenU("F"))
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
		case "e", "E":
			var Foe Agent
			Foe = Spawn()
			Fight(&Char, &Foe)
			continue
		case "c", "C":
			Character(&Char)
			continue
		case "p", "P":
			Preferences(&Char)
			continue
		case "f", "F":
			Banner()
			PickChar()
			continue
		case "h", "H":
			fmt.Println(Blue("\nAll you help are belong to us."))
			Continue()
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

	var cout, fout string

	display := func() {
		c.StatusBar()
		fmt.Println()
		f.FoeBar()
		fmt.Println(Green(SpiderImage()))
	}
	display()

	for {
		fmt.Printf(":> %sight, %svade, %sescribe, %sun\n<: ", GreenU("F"), GreenU("E"), GreenU("D"), GreenU("R"))
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := string([]byte(input)[0])

		switch choice {
		case "f", "F":
			ClearScreen()
			display()
			fmt.Printf("\n%s attacks %s with %s.\n", c.Name, f.Name, c.Weap.Name)
			winner, loser, charAttackDetails := Attack(c, f)
			if c.Name == winner.Name {
				fout = Damage(c, f)
				if loser.Dead == true {
					if f.Weap.Name != c.Weap.Name && loser.DropChance >= Roll(1, 100) {
						OfferItem(c, f)
					}
					WinHeal(c)
					Continue()
					break
				}
			}
			fmt.Printf("\n%s attacks %s with %s.\n", f.Name, c.Name, f.Weap.Name)
			winner, loser, foeAttackDetails := Attack(f, c)
			if f.Name == winner.Name {
				cout = Damage(f, c)
				if loser.Dead == true {
					Continue()
					break
				}
			}
			ClearScreen()
			display()
			fmt.Println(charAttackDetails)
			fmt.Println(fout)
			fmt.Println(foeAttackDetails)
			fmt.Println(cout)
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

func WinHeal(c *Agent) {
	h := Roll(2, c.MxHp.Val)
	if c.MxHp.Val > c.Hp.Val && c.MxHp.Val+30 >= Roll(1, 100) {
		c.AdjHp(h)
		fmt.Printf(Green("\nIn victory heal %d hit points!\n\n"), h)
		c.Save()
	}
}
func OfferItem(c *Agent, f *Agent) {
	ClearScreen()
	fmt.Println(Yellow(Sword1()))
	fmt.Println(Blue("!! ITEM DROP !!"))
	fmt.Printf("%s has dropped it's %s.\n", Yellow(f.Name), Yellow(f.Weap.Name))
	fmt.Printf("Replace?\n	")
	c.Weap.Display()
	fmt.Printf("with..\n	")
	f.Weap.Display()
	if Confirm("\nWould you like to make this swap?") == true {
		c.Weap = f.Weap
		c.Save()
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

//PromptConfirm returns a user entered string with confirmation from
// the user.  e.g.
//answer := PromptConfirm("What's your name? ")
//fmt.Println("Your name is:", answer)
func PromptConfirm(question string) string {
	// assign var input outside of loop so that it's in scope for return
	var response string
Ask:
	for {
		fmt.Printf("%s", Yellow(question))

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		response = scanner.Text()

		// Remove all non alpha characters, including spaces
		response = stripchars(response, " 1234567890,>?<|/{}[]=+-_*&^%$#@!/(/)\\")
		// don't allow empty
		if response == "" {
			continue
		}

	Confirm:
		for {
			fmt.Printf("You choose: %s. confirm y/n? > ", BlueU(response))

			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			choice := string([]byte(input)[0])

			switch choice {
			case "y", "Y":
				break Ask
			case "n", "N":
				break Confirm
			}
		}
	}
	return response
}

func Confirm(response string) bool {
	for {
		fmt.Printf("%s. confirm y/n? > ", Yellow(response))

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := string([]byte(input)[0])

		switch choice {
		case "y", "Y":
			return true
		case "n", "N":
			return false
		}
	}

}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}
