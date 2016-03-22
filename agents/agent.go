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

func (c Agent) ExpDrop() int {
	t := c.Str.Val + c.Int.Val + c.Dex.Val + c.MxHp.Val
	return t / 4
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

	var path string

	if c.File == "New" {
		path = ""
	} else {
		path = "save/"
	}

	if c.File != "" {
		//d, err := os.Open("save/" + c.File + ".json")
		d, err := os.Open(path + c.File + ".json")

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

	if c.Hp.Val < c.MxHp.Val {
		fmt.Printf("%s", Red(strconv.Itoa(c.Hp.Val)))
	} else {
		fmt.Printf("%s", Green(strconv.Itoa(c.Hp.Val)))
	}

	fmt.Printf("%s", Yellow(" XP:"))
	if StatCost(c) <= c.Exp {
		fmt.Printf("%s", Green(strconv.Itoa(c.Exp)))
	} else {
		fmt.Printf("%s", strconv.Itoa(c.Exp))
	}

	fmt.Printf("%s", Yellow(" W:"))
	fmt.Printf("%s", ItemC(c.Weap.Name))
	fmt.Printf("%s", Yellow(" A:"))
	fmt.Printf("%s", ItemC(c.Armor))
	fmt.Printf("%s", Yellow(" T:"))
	fmt.Printf("%s", ItemC(c.Trink))
	//	if c.Dead == false {
	//		fmt.Printf("%s", Yellow(" E?:"))
	//		fmt.Printf("%s", MagentaU("none"))
	//	} else {
	//		fmt.Printf("%s", Red(" Dead :("))
	//	}
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

// StatCost cost increases with the fibonacci number position.
// The average stat is compared with the position in the fibonacci sequence
// then the product of the postion and the multiplier var = Stat Cost
func StatCost(c Agent) int {
	m := 5 // multiplier
	tStat := c.Str.Val + c.Int.Val + c.Dex.Val + c.Hp.Val
	avStat := tStat / 4

	switch {
	case avStat <= 1:
		return 1 * m
	case avStat <= 2:
		return 2 * m
	case avStat <= 3:
		return 3 * m
	case avStat <= 5:
		return 5 * m
	case avStat <= 8:
		return 8 * m
	case avStat <= 13:
		return 13 * m
	case avStat <= 21:
		return 21 * m
	case avStat <= 34:
		return 34 * m
	case avStat <= 55:
		return 55 * m
	case avStat <= 89:
		return 89 * m
	case avStat <= 144:
		return 144 * m
	default:
		return 200 * m
	}
}
