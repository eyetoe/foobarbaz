package agents

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/items"
	. "github.com/eyetoe/foobarbaz/locations"
	. "github.com/eyetoe/foobarbaz/util"
)

// HOME is the game working dir
// SAVE is the game save file under HOME
var HOME, SAVES = FbbDirs()

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
	Exp        int
	DropChance int
	Inv        []Item
	Loc        Loc
	File       string
	Art        string
}

func (c Agent) ExpDrop() int {
	t := c.Str.Val + c.Int.Val + c.Dex.Val + c.MxHp.Val
	return t
}

func (c *Agent) Describe() {
	fmt.Printf(Blue("\nYou consider the %s, "), c.Name)
	fmt.Printf(Blue("%s\n\n"), c.Description)
}

// Adjust Hp "hit points"
func (c *Agent) AdjHp(a int) {
	c.Hp.Val = c.Hp.Val + a
	if c.Hp.Val > c.MxHp.Val {
		c.Hp.Val = c.MxHp.Val
	}
	if c.Hp.Val <= 0 {
		c.Dead = true
		c.Exp = 0
	}
}

// Load Character from json file
func (c *Agent) Load() {
	if c.File != "" {
		d, err := os.Open(SAVES + c.File + ".json")

		if err != nil {
			fmt.Println("Can't open file:", err.Error())
		}
		jp := json.NewDecoder(d)
		if err = jp.Decode(&c); err != nil {
			fmt.Println("Can't parse json:", err.Error())
		} else {
			fmt.Println(Black(SAVES + c.File + ".json loaded."))
		}
	}
	return
}

// Save Character to json file
func (c *Agent) Save() {
	if c.File != "" {
		//Oh sweet MarshalIndent, you make my json look pretty
		i, _ := json.MarshalIndent(c, "", "    ")

		//if err := ioutil.WriteFile("save/"+c.File+".json", i, 0644); err != nil {
		if err := ioutil.WriteFile(SAVES+c.File+".json", i, 0644); err != nil {
			fmt.Println("Can't write file:", err.Error())
		} else {
			fmt.Println(Black(SAVES + c.File + ".json saved."))
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

		for t := 0; t < 5; t++ {
			fmt.Printf(Blue("."))
			time.Sleep(250 * time.Millisecond)
		}
		fmt.Println()
		Resurrect(&c)
		os.Exit(0)
	}
	Meter(c.Hp.Val, c.MxHp.Val, "Health", "=")
}

// StatCost cost increases with the fibonacci number position.
// The average stat is compared with the position in the fibonacci sequence
// then the product of the postion and the multiplier var = Stat Cost
func StatCost(c Agent) int {
	m := 1 // multiplier
	tStat := c.Str.Val + c.Int.Val + c.Dex.Val + c.MxHp.Val
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

// FbbDirs returns the HOME and SAVE vars values
func FbbDirs() (string, string) {
	// Is the ENV var $HOME set?
	home, ok := os.LookupEnv("HOME")
	if ok != true {
		fmt.Println("Please set your $HOME environment variable.")
		os.Exit(1)
	}
	return home + "/.foobarbaz/", home + "/.foobarbaz/saves/"
}

// EnvSetup prepares the game HOME directory
func EnvSetup() {
	// Is the client dir already there?
	if _, err := os.Stat(HOME); err != nil {
		if os.IsNotExist(err) {
			// Make the dir
			if err := os.Mkdir(HOME, 0744); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	// Is the saves dir already there?
	if _, err := os.Stat(SAVES); err != nil {
		if os.IsNotExist(err) {
			// Make the dir
			if err := os.Mkdir(SAVES, 0744); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
	// set working dir to $HOME/.foobarbaz
	if err := os.Chdir(HOME); err != nil {
		fmt.Println("Can't move to working directory", err)
	}

	// TESTING: list files in working dir
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		fmt.Println("list files:", f.Name())
	}
	return
}

func Resurrect(c *Agent) {
	if c.Dead == true {
		//Banner()
		fmt.Printf("[H[J")
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf(Blue("\n A mystical light shines down on %s's lifeless corpse.\n\n A sulfurous effluvium expands from the body.\n\n %s takes a gasping breath, and lives!\n\n"), c.Name, c.Name)
		c.Hp.Val = c.MxHp.Val
		c.Dead = false
		c.Save()
		//Continue()
		for t := 0; t < 5; t++ {
			fmt.Printf(Blue("."))
			time.Sleep(250 * time.Millisecond)
		}
	} else {
		return
	}
}
