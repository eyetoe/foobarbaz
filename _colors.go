package main

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

func Color(c *Agent) {
	// this may be useful it clears the screen
	//fmt.Print("^[[H^[[J")
	Fbb := color.New(color.BgRed, color.FgYellow).SprintFunc()
	Red := color.New(color.BgBlack, color.FgRed).SprintFunc()
	Green := color.New(color.BgBlack, color.FgGreen).SprintFunc()
	BlueU := color.New(color.BgBlack, color.FgHiBlue, color.Bold, color.Underline).SprintFunc()
	ItemC := color.New(color.BgBlack, color.FgHiWhite, color.Bold).SprintFunc()
	attrC := color.New(color.BgBlack, color.FgMagenta, color.Underline).SprintFunc()
	Spc := color.New(color.BgBlack, color.FgYellow).SprintFunc()

	// For sanity layout the StatusBar vertically while printing horizonal
	fmt.Printf("%s", Fbb("FooBarBaz:"))
	fmt.Printf("%s", Spc(" "))
	fmt.Printf("%s", BlueU(c.Name))
	fmt.Printf("%s", Spc(" S:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.Str)))
	fmt.Printf("%s", Spc(" I:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.Int)))
	fmt.Printf("%s", Spc(" D:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.Dex)))
	fmt.Printf("%s", Spc(" HP:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.MxHp)))
	fmt.Printf("%s", Spc("|"))
	fmt.Printf("%s", Red(strconv.Itoa(c.Hp)))
	fmt.Printf("%s", Spc(" W:"))
	fmt.Printf("%s", ItemC(c.Weap))
	fmt.Printf("%s", Spc(" A:"))
	fmt.Printf("%s", ItemC(c.Armor))
	fmt.Printf("%s", Spc(" T:"))
	fmt.Printf("%s", ItemC(c.Trink))
	if c.Dead == false {
		fmt.Printf("%s", Spc(" E?:"))
		fmt.Printf("%s", attrC("none"))
	} else {
		fmt.Printf("%s", Red(" Dead :("))
	}
	fmt.Println()
}
