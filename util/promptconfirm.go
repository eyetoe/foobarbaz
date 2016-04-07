package util

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/eyetoe/foobarbaz/colors"
)

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
