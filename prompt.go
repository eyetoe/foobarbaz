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
		// setting File var in Agent struct, this allowss Char.Load to pick up
		// the save file without prompting.
		Char := Agent{File: "Izro"}
		Char.Load()
		Resurrect(&Char)
		Char.StatusBar()
		//Char.Loc = locations.Start
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
			Go()
			break
		case "c", "C":
			Character(&Char)
			continue
		case "e", "E":
			var Foe Agent
			//Foe = Phantom
			Foe = Spawn()
			//Monster(&Foe)
			Fight(&Char, &Foe)
			continue
		case "h", "H":
			fmt.Println(Blue("All you help are belong to us."))
		}
	}
}
func Character(c *Agent) {
	for {
		c.StatusBar()
		//fmt.Printf(Blue("You examine your character sheet.\n"))
		fmt.Printf("%s:> %sest, %sp, %sbility, %snventory, %sack <:", Blue("Character"), Green("R"), Green("X"), Green("A"), Green("I"), Green("B"))
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := string([]byte(input)[0])
		switch choice {
		case "r", "R":
			fmt.Printf(Blue("You rest and recuperate.\n"))
			if c.Hp.Val < c.MxHp.Val/2 {
				c.Hp.Val = c.MxHp.Val / 2
				fmt.Printf(Green("You heal.\n"))
			}
			continue
		case "x", "X":
			fmt.Printf(Blue("\nYou have %s %s"), Green(strconv.Itoa(c.Exp)), Blue("experience.\n"))
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
}

// Fight loop where c is character and f is foe
func Fight(c *Agent, f *Agent) {
	for {
		c.StatusBar()
		fmt.Printf("You have encountered a %s.\n", WhiteU("Monster!"))
		fmt.Printf("%s:> %sight, %svade, %sescribe, %sun\n<: ", Red(f.Name), GreenU("F"), GreenU("E"), GreenU("D"), GreenU("R"))
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
					break
				}
			}
			fmt.Printf("\n%s attacks %s with %s.\n", f.Name, c.Name, f.Weap.Name)
			winner, loser = Attack(f, c)
			if f.Name == winner.Name {
				Damage(f, c)
				if loser.Dead == true {
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

func Banner() {

	fmt.Println("[H[J")
	fmt.Printf(White("Welcome to ...\n"))
	fmt.Println(Red(
		`   ______         ______           ______              ` + "\n" +
			`   |  ___|        | ___ \          | ___ \             ` + "\n" +
			`   | |_ ___   ___ | |_/ / __ _ _ __| |_/ / __ _ ____   ` + "\n" +
			`   |  _/ _ \ / _ \| ___ \/ _' | '__| ___ \/ _' |_  /   ` + "\n" +
			`   | || (_) | (_) | |_/ / (_| | |  | |_/ / (_| |/ /    ` + "\n" +
			`   \_| \___/ \___/\____/ \__,_|_|  \____/ \__,_/___|   ` + "\n" +
			`                                                       `))
}
