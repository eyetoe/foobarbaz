package main

import (
	"fmt"
	//	. "github.com/eyetoe/foobarbaz/colors"
	//	. "github.com/eyetoe/foobarbaz/util"
	. "github.com/eyetoe/foobarbaz/choices"
)

var help = Choice{
	name:   "help",
	letter: "?",
	action: Showhelp,
}
var thingy = Choice{
	name:   "hingy",
	letter: "t",
	action: Dothis,
}
var mabob = Choice{
	name:   "abob",
	letter: "m",
	action: More.Render,
}
var feblep = Choice{
	name:   "eblep",
	letter: "f",
	action: Feep.Render,
}

var More = Menu{
	[]Choice{
		help,
		thingy,
	},
}
var Todo = Menu{
	[]Choice{
		thingy,
		mabob,
		help,
		feblep,
	},
}
var Feep = Menu{
	list: []Choice{
		{name: "urlap",
			letter: "b",
			action: Burlap,
		},
		{name: "atchel",
			letter: "s",
			action: Satchel,
		},
		help,
	},
}

func main() {
	Todo.Render()

}

func Dothis() {

	fmt.Println(".......................this")
}
func Dothat() {

	fmt.Println(".......................that")
}
func Showhelp() {

	fmt.Println(".......................help")
}
func Burlap() {

	fmt.Println(".......................burlap")
}
func Satchel() {

	fmt.Println(".......................satchel")
}
