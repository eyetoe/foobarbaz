package util

import (
	"fmt"

	. "github.com/eyetoe/foobarbaz/colors"
)

func Continue() {
	fmt.Printf(Blue("\nPress "))
	fmt.Printf(BlueU("Return"))
	fmt.Printf(Blue(" to continue...\n"))

	choice := GetReturn()
	switch choice {
	default:
		return
	}
}
