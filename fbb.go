package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/fatih/color"
)

type Character struct {
	Name  string
	Str   int
	Int   int
	Dex   int
	MxHp  int
	Hp    int
	Weap  string
	Armor string
	Trink string
	Dead  bool
}

// Adjust Hp "hit points"
func (c *Character) adjhp(a int) {
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

func (c *Character) Load(f string) {
	d, err := os.Open(f + ".json")
	if err != nil {
		fmt.Println("Can't open file:", err.Error())
	}
	jp := json.NewDecoder(d)
	if err = jp.Decode(&c); err != nil {
		fmt.Println("Can't parse json:", err.Error())
	}
	return
}

func (c *Character) Save(f string) {
	i, _ := json.Marshal(c)
	//if err := ioutil.WriteFile("playersave.json", i, 0644); err != nil {
	if err := ioutil.WriteFile(f+".json", i, 0644); err != nil {
		fmt.Println("Can't write file:", err.Error())
	}
	return
}

func (c Character) StatusBar() {

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
		fmt.Printf("%s", Red("Dead :("))
	}
	fmt.Println()
}

func main() {

	// create Character struct
	Player := Character{}
	Player.Load("Izro")
	Player.StatusBar()

	Player.Armor = "Tshirt"

	Player.adjhp(-15)
	Player.adjhp(5)

	Player.Save("Izro")
	Player.StatusBar()

	return
}
