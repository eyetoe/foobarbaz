package main

import (
	"fmt"

	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/items"
	. "github.com/eyetoe/foobarbaz/util"
)

func Sandbox() {

	fmt.Println(Switcheroo("This is a string of text with random colors."))
	Teletype("This is a long string of text that will print out slowly........  yep slooooowly")
	Continue()

	Pancho.Display()
	Vest.Display()
	Waistcoat.Display()
	Jacket.Display()
	Leathers.Display()
	Chainmail.Display()
	Platemail.Display()
	Continue()
}
