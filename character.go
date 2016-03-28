package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/colors"
)

func Character(c *Agent) {
	for {
		ClearScreen()
		c.StatusBar()

		fmt.Printf("\n%s \n:> ", BlueU("Character"))
		fmt.Printf("%sest, ", Green("R"))
		fmt.Printf("%sp, ", Green("X"))
		fmt.Printf("%sbility, ", Green("A"))
		fmt.Printf("%snventory, ", Green("I"))
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
		case "b", "B":
			return
		}
	}
}

// PickChar is the initial character load/create prompt
func PickChar() {
	var files []os.FileInfo
	for {
		Banner()
		num := 1
		fmt.Printf(YellowU("\nChoose a character:\n"))
		fmt.Printf(":> %selete <:\n\n", GreenU("D"))
		fmt.Printf(Blue("  %s  %s\n"), GreenU("0"), GreenU("New"))
		//files, _ = ioutil.ReadDir("./save/")
		files, _ = ioutil.ReadDir(SAVES)
		for _, f := range files {
			fmt.Printf("  %s  %s\n", GreenU(strconv.Itoa(num)), Blue(strings.Replace(f.Name(), ".json", "", -1)))
			num++
		}
		fmt.Printf("\n<: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := string([]byte(input)[0])

		cnum := 1
		for _, f := range files {
			if choice == strconv.Itoa(cnum) {
				fmt.Printf("Loading --> %s\n", Blue(f.Name()))
				SaveFile = strings.Replace(f.Name(), ".json", "", -1)
				return
			} else {
				cnum++
			}
		}
		switch choice {
		case "0", "o", "O":
			if cnum >= 10 {
				fmt.Println(Red("\nYou already have the maximum number of characters.\nDelete one first.\n"))
				Continue()
				continue
			}
			NewCharacter()
			return
		//case "b", "B":
		//return
		case "d", "D":
			DeleteCharacter()
			break
		}
	}
}

func DeleteCharacter() {
	var files []os.FileInfo
	for {
		Banner()
		num := 1
		fmt.Printf(RedU("\nDelete which character?\n"))
		fmt.Printf(":> %sack <:\n\n", GreenU("B"))
		//fmt.Printf(Blue("  %s  %s\n"), GreenU("0"), Blue("New"))
		// List Characters for selection
		files, _ = ioutil.ReadDir(SAVES)
		for _, f := range files {
			fmt.Printf("  %s  %s\n", GreenU(strconv.Itoa(num)), Red(strings.Replace(f.Name(), ".json", "", -1)))
			num++
		}
		fmt.Printf("\n<: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := string([]byte(input)[0])

		cnum := 1
		// Delete the choice
		for _, f := range files {
			if choice == strconv.Itoa(cnum) {
				fmt.Printf("Deleting --> %s\n", Red(f.Name()))

				if Confirm("Are you sure you want to delete " + f.Name()) {
					err := os.Remove(SAVES + f.Name())
					if err != nil {
						fmt.Println(Red(err))
					}
					Continue()
					continue
				} else {
					continue
				}

			} else {
				cnum++
			}
		}
		switch choice {
		case "b", "B":
			return
		}
	}
}

func NewCharacter() {
	SaveFile = PromptConfirm("What shall your name be? ")

	//make sure we don't use a name already taken
	files, _ := ioutil.ReadDir(SAVES)
	for _, f := range files {
		chkFile := strings.Replace(f.Name(), ".json", "", -1)
		if SaveFile == chkFile || SaveFile == "New" {
			fmt.Printf("\nThe name %s is already taken\n\n", Red(SaveFile))
			Continue()
			NewCharacter()
			return
		}
	}

	//Char := Agent{File: "New"}
	//Char.Load()
	Char := New
	Char.Name = SaveFile
	Char.File = SaveFile
	Char.Save()

	fmt.Printf("\nThen we shall call you, %s\n\n", BlueU(SaveFile))
	Continue()
	return
}

func ExpMgr(c *Agent) {
	for {
		ClearScreen()
		c.StatusBar()
		cost := StatCost(*c)
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