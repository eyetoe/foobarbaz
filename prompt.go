package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/colors"
)

func Prompt() {
	for {
		Char := Agent{File: "Izro"}
		Char.Load()
		Foe := Minotaur
		Char.StatusBar()

		fmt.Printf(":> A %s does not see you. Engage? %s?\n<: ", Foe.Name, GreenU("y/n"))
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := string([]byte(input)[0])

		switch choice {
		case "y", "Y":
			fmt.Printf("You engage the %s.\n", Foe.Name)
			Fight(&Char, &Foe)
			continue
		case "n", "N":
			fmt.Println("You slip into the shadows unseen.")
			break
		default:
			Prompt()
			break
		}

		return
	}
}

func Fight(c *Agent, f *Agent) {
	for {
		c.StatusBar()
		fmt.Printf(":> %sight, %svade, %sescribe, %sun\n<: ", GreenU("F"), GreenU("E"), GreenU("D"), GreenU("R"))
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := string([]byte(input)[0])

		switch choice {
		case "f", "F":
			fmt.Printf("\n%s attacks %s with %s.\n", c.Name, f.Name, c.Weap.Name)
			winner, loser := Attack(c, f)
			if c.Name == winner.Name {
				Damage(*winner, loser)
				if loser.Dead == true {
					break
				}
			}
			fmt.Printf("\n%s attacks %s with %s.\n", f.Name, c.Name, f.Weap.Name)
			winner, loser = Attack(f, c)
			if f.Name == winner.Name {
				Damage(*winner, loser)
				if loser.Dead == true {
					break
				}
			}

			continue
		case "e", "E":
			fmt.Println("Evade!\n You stall for time, looking for an opening!")
			continue
		case "d", "D":
			f.FoeBar()
			fmt.Printf("You consider the %s. %s\n", f.Name, f.Description)
			continue
		case "r", "R":
			fmt.Println("Run!\n You have lost this battle, but may yet win the war!")
			break
		default:
			Fight(c, f)
			break
		}

		// fmt.Printf("You entered: %v", []byte(input))
		return
	}

}
