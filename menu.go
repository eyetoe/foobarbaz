package main

import (
	"fmt"
	"os"

	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/util"
)

type Menu struct {
	list []Choices
}

type Choices struct {
	name   string
	letter string
	action func()
}

func (m Menu) Render() {
	for {
		fmt.Printf(":> %sxit, %sack %sick Monster <:", Green("E"), Green("B"), Green("P"))
		choice, _, _ := GetChar()

		switch choice {
		case "e", "E":
			os.Exit(0)
		case "b", "B":
			return
		}
	}
}
