package main

import (
	"fmt"
	"os"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/util"
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
		fmt.Printf("%svent, ", GreenU("E"))
		fmt.Printf("%sest, ", Green("R"))
		fmt.Printf("%srain ", Green("T"))
		fmt.Printf("%sile, ", GreenU("F"))
		fmt.Printf("%so <: ", GreenU("G"))

		choice, _, _ := GetChar()

		// choices here
		switch choice {
		case "e", "E":
			var Foe Agent
			Foe = Spawn()
			Fight(&Char, &Foe)
			continue
		case "r", "R":
			fmt.Printf(Blue("\nYou rest and recuperate.\n"))
			if Char.Hp.Val < Char.MxHp.Val/2 {
				Char.Hp.Val = Char.MxHp.Val / 2
				fmt.Printf(Green("You heal.\n"))
				Char.Save()
			}
			Continue()
			continue
		case "t", "T":
			ExpMgr(&Char)
			continue
		case "f", "F":
			Banner()
			PickChar()
			Char.Load()
			Resurrect(&Char)
			continue
		case "g", "G":
			Go(&Char)
			break
		}
	}
}

func Go(c *Agent) {
	ClearScreen()
	c.StatusBar()
	fmt.Println()
	fmt.Printf(":> %sxit, %sack <:", Green("E"), Green("B"))
	choice, _, _ := GetChar()

	switch choice {
	case "e", "E":
		os.Exit(0)
	case "b", "B":
		return
	}
}
