package agents

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/eyetoe/foobarbaz/inv"
	"github.com/fatih/color"
)

type Stat struct {
	Name string
	Val  int
}

type Agent struct {
	Name        string
	Description string
	// 3 base attributes Str, Int, Dex Str int
	Str int
	Int int
	Dex int
	// Max Health and Health
	MxHp int
	Hp   int
	Dead bool
	// 3 item slots: Weapon, Armor, Trinket
	Weap  string
	Armor string
	Trink string
	// Special Abilities
	Abl1 string
	Abl2 string
	Abl3 string
	//
	Inv []inv.Item
}

func (c *Agent) Describe() {
	fmt.Printf("You consider the %s. ", c.Name)
	fmt.Printf("%s\n", c.Description)

}

// Adjust Hp "hit points"
func (c *Agent) AdjHp(a int) {
	c.Hp = c.Hp + a
	if a > 0 {
		fmt.Println(c.Name, "heals", a, "hit points")
	} else {
		fmt.Println(c.Name, "takes", a, "damage!")
	}
	if c.Hp > c.MxHp {
		c.Hp = c.MxHp
	}
	if c.Hp < 0 {
		fmt.Println(c.Name, "has died :( ")
		c.Dead = true
	}
}

// Load Character from json file
func (c *Agent) Load(f string) {
	d, err := os.Open("save/" + f + ".json")
	fmt.Println("./save/"+f+".json", "loaded.")

	if err != nil {
		fmt.Println("Can't open file:", err.Error())
	}
	jp := json.NewDecoder(d)
	if err = jp.Decode(&c); err != nil {
		fmt.Println("Can't parse json:", err.Error())
	}
	return
}

// Save Character to json file
func (c *Agent) Save(f string) {
	fmt.Println("./save/"+f+".json", "saved.")
	//Oh sweet MarshalIndent, you make my json look pretty
	i, _ := json.MarshalIndent(c, "", "    ")

	if err := ioutil.WriteFile("save/"+f+".json", i, 0644); err != nil {
		fmt.Println("Can't write file:", err.Error())
	}
	return
}

// Render character status bar
func (c Agent) StatusBar() {
	// this may be useful it clears the screen
	//fmt.Print("[H[J")
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
