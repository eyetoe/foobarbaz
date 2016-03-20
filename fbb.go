package main

import (
	"fmt"
	"os"
	"strconv"

	. "github.com/eyetoe/foobarbaz/affects"
	// dot preceding the import means, use this namespace for the import
	// this means functions in the imported package don't need to have the
	// directory prefixed, soo you can use Agent() rather than agents.Agent()
	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/items"
	. "github.com/eyetoe/foobarbaz/skills"
)

func main() {
	Banner()
	// for testing, Resurrect if character is dead
	Char := Agent{File: "Izro"}
	Char.Load()
	Resurrect(&Char)

	Continue()
	Prompt()
	Testies()
	return
}

// Start the game.
func Play() {
	Char := Agent{File: "Izro"}
	Char.Load()
	Char.StatusBar()
}

func Resurrect(c *Agent) {
	if c.Dead == true {
		fmt.Printf(Blue("\n A mystical light shines down on %s's lifeless corpse.\n\n A sulfurous effluvium expands from the body.\n\n %s takes a gasping breath, and lives!\n\n"), c.Name, c.Name)
		c.Hp.Val = c.MxHp.Val
		c.Dead = false
		c.Save()
	} else {
		return
	}
}

func Testies() {
	Char := Agent{File: "Izro"}
	Char.Load()
	Char.StatusBar()

	Foe := Minotaur

	if len(os.Args) >= 2 {

		//arg2, err := os.Args[2:][0]
		arg1 := os.Args[1:][0]
		fmt.Println("Running test:", arg1)

		switch arg1 {
		// display affect
		case "describe":
			fmt.Println("======================= Testing Character Description")
			Char.Describe()
		case "affect":
			fmt.Println("======================= Testing Affect Struct")
			fmt.Printf("%s has %s%% chance to cause affect.\nDescription: %s is %s\n", OnFire.Name, strconv.Itoa(OnFire.Proc), OnFire.Name, OnFire.Description)
		// display item
		case "item":
			fmt.Println("======================= Testing Item Struct")
			//fmt.Printf("%s is equipped in the %s slot.\nIt grants + %s to attack rolls, and + %s to damage rolls\nDescription: %s is %s\n", Staff.Name, Staff.Slot, strconv.Itoa(Staff.Attack), strconv.Itoa(Staff.Damage), Staff.Name, Staff.Description)
			Char.Weap = Staff
			Char.Save()
		// display foe
		case "foe":
			fmt.Println("======================= Testing Foe Description")
			Foe.StatusBar()
			Foe.Save()
			Foe.Describe()
		// display combat
		case "combat":
			combat := func() {
				// working on flow for a wrapper function to display combat dialog
				fmt.Println("======================= Testing combat dialog")
				atk := &Char
				def := &Foe

				fmt.Printf("%s attacks %s with %s.\n", atk.Name, def.Name, atk.Weap.Name)
				winner, loser := Attack(&Char, &Foe)
				fmt.Printf("%s has prevailed!\n", winner.Name)
				fmt.Printf("Alas, %s has fallen short!\n", loser.Name)

				return
			}
			combat()
		// display contest
		case "contest":
			// working on flow for a wrapper function to display comba
			contest := func() {
				fmt.Println("======================= Testing contest dialog")
				atk := &Char
				def := &Foe
				cwinner, closer := Contest(&Char, Char.Str, &Foe, Foe.Str)
				fmt.Printf("%s vs. %s in a battle of skill.\n", atk.Name, def.Name)
				fmt.Printf("%s has prevailed!\n", cwinner.Name)
				fmt.Printf("Alas, %s has fallen short!\n", closer.Name)
				return
			}

			for i := 0; i < 5; i++ {
				contest()
			}
		// display contest
		case "skill":
			for i := 0; i < 5; i++ {
				fmt.Println("======================= Testing skill check dialog")
				//if Skill(Char, Char.Dex, Stat{"Roll", Roll()}) {
				if SkillCheck(Char, Char.Str, Door) {
					fmt.Println("Skill is true")
				}
			}
		// display dice rolls
		case "dice":
			test_dice(20)
		// adjust hp
		case "adjust":
			if len(os.Args) < 2 {
				arg2 := 0
				fmt.Printf("arg2 is: %d\n", arg2)
			} else {
				arg2, err := strconv.Atoi(os.Args[2:][0])
				if err != nil {
					usage()
					fmt.Printf("%q\n", err)
				} else {
					Char.AdjHp(arg2)
					if arg2 > 0 {
						fmt.Printf(Green("%s heals %d!\n"), Char.Name, arg2)
					} else {
						fmt.Printf(Red("%s takes %d damage!\n"), Char.Name, arg2)
					}

					Char.Save()
					Char.StatusBar()
				}
			}
		case "color", "colors":
			fmt.Println("Say what?", Fbb("I said Bam!"))
			fmt.Println(Fbb("Fbb"), Red("Red"), Green("Green"), BlueU("BlueU"), ItemC("ItemC"), AttrC("attrC"), Spc("Spc"))

		default:
			fmt.Printf("Run specific tests by passing one of the following arguments:\n dice\n skill\n combat\n contest\n item\n affect\n foe\n color\n adjust <int>\n")
			fmt.Printf("e.g. :\n go run *.go dice\n\n")
		} // switch
	} // if len(os.Args) >= 2

	return
} // main

// usage message for cli
func usage() {
	fmt.Println("Usage: fbb (int hp adjust)")
}

// test die rolls
func test_dice(n int) {
	fmt.Printf("======================= Testing %d d100 die rolls.\n", n)

	// roll some dice n times
	for i := 0; i < n; i++ {
		fmt.Printf("You roll the dice: %d\n", Roll(100))
	}
}
