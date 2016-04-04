package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/art"
	. "github.com/eyetoe/foobarbaz/colors"
	"github.com/pkg/term"
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

		choice, _, _ := GetChar()

		// choices here
		switch choice {
		case "l", "L":
			Look(Char)
			continue
		case "g", "G":
			Go(&Char)
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
			Char.Load()
			Resurrect(&Char)
			continue
		case "h", "H":
			fmt.Println(Blue("\nAll you help are belong to us."))
			Continue()
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
func Look(c Agent) {
	fmt.Printf(Blue("You are here: %s\n"), c.Loc.Name)
	fmt.Printf(Blue("%s\n"), c.Loc.Description)
	Continue()
}

func WinHeal(c *Agent) string {
	var textOut string
	h := Roll(2, c.MxHp.Val)
	if c.MxHp.Val > c.Hp.Val && c.MxHp.Val+30 >= Roll(1, 100) {
		textOut = textOut + c.AdjHp(h)
		textOut = textOut + fmt.Sprintf(Green("\nIn victory heal %d hit points!\n\n"), h)
		c.Save()
	}
	return textOut
}
func OfferItem(c *Agent, f *Agent) {
	ClearScreen()
	fmt.Println(Yellow(Sword1()))
	fmt.Println(Blue("!! ITEM DROP !!\n"))
	fmt.Printf("%s has dropped it's %s.\n\n", Yellow(f.Name), Yellow(f.Weap.Name))
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
	fmt.Printf(Blue("\nPress "))
	fmt.Printf(BlueU("Return"))
	fmt.Printf(Blue(" to continue...\n"))

	//choice, _, _ := GetChar()
	choice := GetReturn()
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
		// don't use this commented func..
		//response = GetReturn()

		// Remove all non alpha characters, including spaces
		response = stripchars(response, " 1234567890,>?<|/{}[]=+-_*&^%$#@!/(/)\\")
		// don't allow empty
		if response == "" {
			continue
		}

	Confirm:
		for {
			fmt.Printf("You choose: %s. confirm y/n? > ", BlueU(response))

			choice := GetReturn()

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

		choice := GetReturn()

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

// GetReturn returns a single string character after user enters newline
func GetReturn() (ascii string) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return string([]byte(input)[0])
}

// GetChar returns a single string character, without waiting for newline
func GetChar() (ascii string, keyCode int, err error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)

	var numRead int
	numRead, err = t.Read(bytes)
	if err != nil {
		return
	}
	if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
		// Three-character control sequence, beginning with "ESC-[".

		// Since there are no ASCII codes for arrow keys, we use
		// Javascript key codes.
		if bytes[2] == 65 {
			// Up
			keyCode = 38
		} else if bytes[2] == 66 {
			// Down
			keyCode = 40
		} else if bytes[2] == 67 {
			// Right
			keyCode = 39
		} else if bytes[2] == 68 {
			// Left
			keyCode = 37
		}
	} else if numRead == 1 {
		ascii = fmt.Sprintf("%c", int(bytes[0]))
	} else {
		// Two characters read??
	}
	t.Restore()
	t.Close()
	return
}
