package main

import (
	"fmt"
	"strconv"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/simulations"
)

// Render character status bar
func FoeBar(c Agent, f Agent) {
	// this may be useful it clears the screen
	//fmt.Print("[H[J")

	// Color Foe's name and show victory chance percentage
	odds := Odds(&c, &f)
	var conName string
	switch {
	case odds >= 80:
		conName = fmt.Sprintf(GreenU("%s:%d%%"), f.Name, odds)
	case odds >= 60:
		conName = fmt.Sprintf(CyanU("%s:%d%%"), f.Name, odds)
	case odds >= 40:
		conName = fmt.Sprintf(BlueU("%s:%d%%"), f.Name, odds)
	case odds >= 20:
		conName = fmt.Sprintf(YellowU("%s:%d%%"), f.Name, odds)
	case odds >= 0:
		conName = fmt.Sprintf(RedU("%s:%d%%"), f.Name, odds)
	}

	// For sanity layout the StatusBar vertically here although printing horizonal
	fmt.Printf("%s", Yellow(" "))
	//fmt.Printf("%s", RedU(f.Name))
	fmt.Printf("%s", RedU(conName))
	fmt.Printf("%s", Yellow(" S:"))
	fmt.Printf("%s", Green(strconv.Itoa(f.Str.Val)))
	fmt.Printf("%s", Yellow(" I:"))
	fmt.Printf("%s", Green(strconv.Itoa(f.Int.Val)))
	fmt.Printf("%s", Yellow(" D:"))
	fmt.Printf("%s", Green(strconv.Itoa(f.Dex.Val)))
	fmt.Printf("%s", Yellow(" HP:"))
	fmt.Printf("%s", Green(strconv.Itoa(f.MxHp.Val)))
	fmt.Printf("%s", Yellow("|"))
	fmt.Printf("%s", Red(strconv.Itoa(f.Hp.Val)))
	fmt.Printf("%s", Yellow(" W:"))
	fmt.Printf("%s", White(f.Weap.Name))
	fmt.Printf("%s", Yellow(" A:"))
	fmt.Printf("%s", White(f.Armor))
	fmt.Printf("%s", Yellow(" T:"))
	fmt.Printf("%s", White(f.Trink))
	fmt.Println()
}
