package choices

import (
	"fmt"

	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/util"
)

type Menu struct {
	list []Choice
}

func (m Menu) Render() {

Menu:
	for {
		fmt.Printf(":> ")
		for _, each := range m.list {
			fmt.Printf("%s%s, ", GreenU(each.letter), Blue(each.name))
		}
		fmt.Printf("\nChoose <: ")
		c, _, _ := GetChar()

		for _, each := range m.list {
			if each.letter == c {
				fmt.Println(c)
				each.action()
				break Menu
			}
		}
		continue
	}
}
