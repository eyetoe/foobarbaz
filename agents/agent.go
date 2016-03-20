package agents

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/items"
	. "github.com/eyetoe/foobarbaz/locations"
)

type Stat struct {
	Name string
	Val  int
}

type Agent struct {
	Name        string
	Description string
	// 3 base attributes Str, Int, Dex
	Str Stat
	Int Stat
	Dex Stat
	// Max Health and Health
	MxHp Stat
	Hp   Stat
	Dead bool
	// 3 item slots: Weapon, Armor, Trinket
	Weap  Item
	Armor string
	Trink string
	// Special Abilities
	Abl1 string
	Abl2 string
	Abl3 string
	//
	Exp  int
	Inv  []Item
	Loc  Loc
	File string
}

func (c *Agent) Describe() {
	fmt.Printf("You consider the %s. ", c.Name)
	fmt.Printf("%s\n", c.Description)
}

// Adjust Hp "hit points"
func (c *Agent) AdjHp(a int) {
	c.Hp.Val = c.Hp.Val + a
	//	if a > 0 {
	//		fmt.Println(c.Name, "heals", a, "hit points")
	//} else {
	//fmt.Println(c.Name, "takes", a, "damage!")
	//}
	if c.Hp.Val > c.MxHp.Val {
		c.Hp.Val = c.MxHp.Val
	}
	if c.Hp.Val <= 0 {
		//fmt.Println(c.Name, "has died :( ")
		fmt.Printf("%s", BlueU(c.Name))
		fmt.Printf(Red(" has died:( \n"))
		c.Dead = true
		c.Exp = 0
	}
}

// Load Character from json file
func (c *Agent) Load() {
	if c.File != "" {
		d, err := os.Open("save/" + c.File + ".json")

		if err != nil {
			fmt.Println("Can't open file:", err.Error())
		}
		jp := json.NewDecoder(d)
		if err = jp.Decode(&c); err != nil {
			fmt.Println("Can't parse json:", err.Error())
		} else {
			fmt.Println(Black("./save/" + c.File + ".json loaded."))
		}
	}
	return
}

// Save Character to json file
func (c *Agent) Save() {
	if c.File != "" {
		//Oh sweet MarshalIndent, you make my json look pretty
		i, _ := json.MarshalIndent(c, "", "    ")

		if err := ioutil.WriteFile("save/"+c.File+".json", i, 0644); err != nil {
			fmt.Println("Can't write file:", err.Error())
		} else {
			fmt.Println(Black("./save/" + c.File + ".json saved."))
		}
	}
	return
}

// Render character status bar
func (c Agent) StatusBar() {
	// For sanity layout the StatusBar vertically while printing horizonal
	fmt.Println()
	fmt.Printf("%s", Fbb("FooBarBaz:"))
	fmt.Printf("%s", Yellow(" "))
	fmt.Printf("%s", BlueU(c.Name))
	fmt.Printf("%s", Yellow(" S:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.Str.Val)))
	fmt.Printf("%s", Yellow(" I:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.Int.Val)))
	fmt.Printf("%s", Yellow(" D:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.Dex.Val)))
	fmt.Printf("%s", Yellow(" HP:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.MxHp.Val)))
	fmt.Printf("%s", Yellow("|"))
	fmt.Printf("%s", Red(strconv.Itoa(c.Hp.Val)))

	fmt.Printf("%s", Yellow(" XP:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.Exp)))

	fmt.Printf("%s", Yellow(" W:"))
	fmt.Printf("%s", ItemC(c.Weap.Name))
	fmt.Printf("%s", Yellow(" A:"))
	fmt.Printf("%s", ItemC(c.Armor))
	fmt.Printf("%s", Yellow(" T:"))
	fmt.Printf("%s", ItemC(c.Trink))
	if c.Dead == false {
		fmt.Printf("%s", Yellow(" E?:"))
		fmt.Printf("%s", MagentaU("none"))
	} else {
		fmt.Printf("%s", Red(" Dead :("))
	}
	fmt.Println()
	if c.Dead == true {
		fmt.Printf(Red("\n%s collapsed in a sobbing frightned lump and expired.\n\n"), c.Name)
		os.Exit(0)
	}
}

// Render character status bar
func (c Agent) FoeBar() {
	// this may be useful it clears the screen
	//fmt.Print("[H[J")

	// For sanity layout the StatusBar vertically here although printing horizonal
	fmt.Printf("%s", Yellow(" "))
	fmt.Printf("%s", RedU(c.Name))
	fmt.Printf("%s", Yellow(" S:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.Str.Val)))
	fmt.Printf("%s", Yellow(" I:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.Int.Val)))
	fmt.Printf("%s", Yellow(" D:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.Dex.Val)))
	fmt.Printf("%s", Yellow(" HP:"))
	fmt.Printf("%s", Green(strconv.Itoa(c.MxHp.Val)))
	fmt.Printf("%s", Yellow("|"))
	fmt.Printf("%s", Red(strconv.Itoa(c.Hp.Val)))
	fmt.Printf("%s", Yellow(" W:"))
	fmt.Printf("%s", White(c.Weap.Name))
	fmt.Printf("%s", Yellow(" A:"))
	fmt.Printf("%s", White(c.Armor))
	fmt.Printf("%s", Yellow(" T:"))
	fmt.Printf("%s", White(c.Trink))
	fmt.Println()
}
