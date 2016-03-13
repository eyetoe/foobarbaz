package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/colors"
)

//func main() {

//	Prompt()

//	return
//}

func Prompt() {
	for {
		Char := Agent{}
		Char.Load("Izro")
		Foe := Minotaur
		Char.StatusBar()

		fmt.Printf(":> A %s does not see you. Engage? %s?\n<: ", Foe.Name, GrnU("y/n"))
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := string([]byte(input)[0])

		switch choice {
		case "y", "Y":
			fmt.Printf("You engage the %s.\n", Foe.Name)
			Fight(&Foe)
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

func Fight(f *Agent) {
	for {
		fmt.Printf(":> %sight, %svade, %sescribe, %sun\n<: ", GrnU("F"), GrnU("E"), GrnU("D"), GrnU("R"))
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := string([]byte(input)[0])

		switch choice {
		case "f", "F":
			fmt.Println("Fight!\n You enter the fray!")
			continue
		case "e", "E":
			fmt.Println("Evade!\n You stall for time, looking for an opening!")
			continue
		case "d", "D":
			fmt.Printf("You consider the %s. %s\n", f.Name, f.Description)
			continue
		case "r", "R":
			fmt.Println("Run!\n You have lost this battle, but may yet win the war!")
			break
		default:
			Fight(f)
			break
		}

		// fmt.Printf("You entered: %v", []byte(input))
		return
	}
}
