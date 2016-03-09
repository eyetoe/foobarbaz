package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/eyetoe/foobarbaz/affects"
	"github.com/eyetoe/foobarbaz/agents"
	"github.com/eyetoe/foobarbaz/inv"
)

func main() {
	Char := agents.Agent{}
	Char.Load("Izro")
	Char.StatusBar()
	Char.Describe()

	Foe := agents.Minotaur

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

	//Char.adjhp(5)

	Char.Save("Izro")
	Char.StatusBar()

	fmt.Println(affects.OnFire)
	fmt.Println(inv.Staff)
	fmt.Println("Foe is:\n", Foe)

	Foe.StatusBar()
	Foe.Save("Minotaur")
	Foe.Describe()

	fmt.Printf("You roll the dice: %d\n", Roll())
	fmt.Printf("You roll the dice: %d\n", Roll())
	fmt.Printf("You roll the dice: %d\n", Roll())
	fmt.Printf("You roll the dice: %d\n", Roll())
	fmt.Printf("You roll the dice: %d\n", Roll())
	fmt.Printf("You roll the dice: %d\n", Roll())
	fmt.Printf("You roll the dice: %d\n", Roll())
	fmt.Printf("You roll the dice: %d\n", Roll())

	// working on flow for a wrapper function to display combat dialog
	atk := &Char
	def := &Foe

	fmt.Printf("%s attacks %s with %s.\n", atk.Name, def.Name, atk.Weap)
	winner, loser := Attack(&Char, &Foe)
	fmt.Printf("%s has prevailed!\n", winner.Name)
	fmt.Printf("Alas %s has fallen short!\n", loser.Name)
	//

	return
}

func Usage() {
	fmt.Println("Usage: fbb (int hp adjust)")
}
