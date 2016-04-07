package util

import (
	"fmt"

	. "github.com/eyetoe/foobarbaz/colors"
)

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
