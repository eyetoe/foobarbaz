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

	return
}

func Usage() {
	fmt.Println("Usage: fbb (int hp adjust)")
}
