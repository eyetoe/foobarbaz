package main

import (
	"fmt"
	"os"
	"time"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/util"
)

func Prompt() {
	for {
		// setting File var in Agent struct, this allows Char.Load to pick up
		// the save file without prompting.
		//Char := Agent{File: SaveFile}
		//Char.Load()
		ClearScreen()
		if Char.Dead == true {
			fmt.Printf(Red("\n%s collapsed in a sobbing frightned lump and expired.\n\n"), Char.Name)

			for t := 0; t < 5; t++ {
				fmt.Printf(Blue("."))
				time.Sleep(100 * time.Millisecond)
			}
			Continue()
			fmt.Println(Dying())
			if Confirm(Magenta("Rage against the dying of the light?\n\n")) == true {
				Resurrect(&Char)
				continue
			} else {
				os.Exit(0)
			}
		}
		Char.StatusBar()
		Char.Save()
		fmt.Printf("\n%s <- You are here.\n", BlueU(Char.Loc.Name))

		// ask the first question
		fmt.Printf("%svent, ", GreenU("E"))
		fmt.Printf("%sest, ", GreenU("R"))
		fmt.Printf("%srain ", GreenU("T"))
		fmt.Printf("%so, ", GreenU("G"))
		fmt.Printf("%soad, ", GreenU("L"))
		fmt.Printf("%spawn, ", GreenU("S"))
		fmt.Printf("%snventory <: ", GreenU("I"))

		choice, _, _ := GetChar()

		// choices here
		switch choice {
		case "e", "E":
			//var Foe Agent
			Foe = Spawn(Char)
			Fight(&Char, &Foe)
			continue
		case "r", "R":
			fmt.Printf(Blue("\nYou rest and recuperate.\n"))
			if Char.Health.Val < Char.MaxHealth.Val/2 {
				Char.Health.Val = Char.MaxHealth.Val / 2
				fmt.Printf(Green("You heal.\n"))
				Char.Save()
			}
			Continue()
			continue
		case "t", "T":
			ExpMgr(&Char)
			continue
		case "g", "G":
			Go(&Char)
			break
		case "l", "L":
			Banner()
			PickChar()
			Char.Load()
			Resurrect(&Char)
			continue
		case "s", "S":
			var Foe Agent
			Foe = *SpawnChooser(&Char)
			Fight(&Char, &Foe)
			continue
		case "i", "I":
			Inventory(&Char)
			break
		}
	}
}

func Go(c *Agent) {
	for {
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
}
