package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/colors"
	"github.com/eyetoe/foobarbaz/locations"
)

func Prompt() {
	for {
		// setting File var in Agent struct, this allowss Char.Load to pick up
		// the save file without prompting.
		Char := Agent{File: "Izro"}
		Char.Load()
		Char.StatusBar()
		Char.Loc = locations.Start
		Char.Save()
		fmt.Printf("You are here: "+BlueU("%s\n"), Char.Loc.Name)

		// ask the firt question
		//fmt.Printf("Load, Save, New, Help\n")
		fmt.Printf(":> %sook, %so, %sharacter, %svent, %selp <: ", GreenU("L"), GreenU("G"), GreenU("C"), GreenU("E"), GreenU("H"))

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := string([]byte(input)[0])

		// choices here
		switch choice {
		case "l", "L":
			Look(Char)
			continue
		case "g", "G":
			fmt.Printf(Blue("Where are we going?.\n"))
			break
		case "c", "C":
			fmt.Printf(Blue("You examine your character sheet.\n"))
			CharacterSheet(&Char)
			continue
		case "e", "E":
			fmt.Printf(Blue("If there is an event, let's do it!\n"))
			var Foe Agent
			//Foe = Phantom
			Foe = Spawn()
			Monster(&Foe)
			continue
		case "h", "H":
			fmt.Println(Blue("All you help are belong to us."))
		}
	}
}
func Look(c Agent) {
	fmt.Printf(Blue("You are here: %s"), c.Loc)
}
func CharacterSheet(c *Agent) {
	fmt.Printf("%s\n", c.Name)
}
func Monster(f *Agent) {
	fmt.Printf("%s\n", f.Name)
}

func Banner() {
	fmt.Println("[H[J")
	fmt.Printf(White("Welcome to ...\n"))
	fmt.Printf(Red("FooBarBaz\n\n"))
}
