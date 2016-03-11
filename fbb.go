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
	. "github.com/eyetoe/foobarbaz/inv"
)

func main() {
	Char := Agent{}
	Char.Load("Izro")
	Char.StatusBar()
	fmt.Println("======================= Testing Character Description")
	Char.Describe()

	//Foe := agents.Minotaur
	Foe := Minotaur

	// Take first arg as hit point adjust
	//[1:] is the slice from 2nd argument (skipping prog name)
	//[0] is the first element in the array that is returned
	arg1, err := strconv.Atoi(os.Args[1:][0])
	if err != nil {
		Usage()
		fmt.Printf("%q\n", err)
	} else {
		Char.AdjHp(arg1)
	}

	// create Agent struct
	Char.Armor = "Tshirt"

	Char.Save("Izro")
	Char.StatusBar()

	fmt.Println("======================= Testing Affect Struct")
	fmt.Println(OnFire)
	fmt.Println("======================= Testing Item Struct")
	fmt.Println(Staff)

	Foe.StatusBar()
	Foe.Save("Minotaur")
	fmt.Println("======================= Testing Foe Description")
	Foe.Describe()

	// combat
	combat := func() {
		// working on flow for a wrapper function to display combat dialog
		fmt.Println("======================= Testing combat dialog")
		atk := &Char
		def := &Foe

		fmt.Printf("%s attacks %s with %s.\n", atk.Name, def.Name, atk.Weap)
		winner, loser := Attack(&Char, &Foe)
		fmt.Printf("%s has prevailed!\n", winner.Name)
		fmt.Printf("Alas, %s has fallen short!\n", loser.Name)

		return
	}
	combat()

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

	for i := 0; i < 5000; i++ {
		contest()
	}

	// roll some dice!
	//	test_dice(20)

	return
}

// usage message for cli
func Usage() {
	fmt.Println("Usage: fbb (int hp adjust)")
}

// test die rolls
func test_dice(n int) {
	fmt.Printf("======================= Testing %d d100 die rolls.\n", n)

	// roll some dice n times
	for i := 0; i < n; i++ {
		fmt.Printf("You roll the dice: %d\n", Roll())
	}
}
