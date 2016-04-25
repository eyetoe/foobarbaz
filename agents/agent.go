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
	Armor Item
	Trink Item
	// Special Abilities
	Abl1 string
	Abl2 string
	Abl3 string
	//
	Exp int
	//DropChance int
	Inv  []Item
	Loc  Loc
	File string
	Art  string
	// set this during combat, and clear after.  Used to decide when to show
	// red in health meter to indicate when you are within range of a killing blow.
	FoeMaxHit int
}

func (c Agent) ExpDrop() int {
	// add character stats
	s := c.Str.Val + c.Int.Val + c.Dex.Val + c.MxHp.Val
	// and weapon stats
	w := c.Weap.Attack + c.Weap.Damage + c.Weap.Defence + c.Weap.Dodge + c.Weap.Crit + c.Weap.DropChance
	// and armor stats
	a := c.Armor.Attack + c.Armor.Damage + c.Armor.Defence + c.Armor.Dodge + c.Armor.Crit + c.Armor.DropChance
	return s + w + a
}

func (c *Agent) Describe() {
	fmt.Printf(YellowU("\n%s:\n"), c.Name)
	fmt.Printf(Cyan("You consider the %s, "), c.Name)
	fmt.Printf(Cyan("%s\n"), c.Description)
	fmt.Printf("%s %s\n", White(" Total critical chance:"), Yellow(strconv.Itoa(c.Weap.Crit+(c.Int.Val/2)), "%\n"))
	c.Weap.Display()
	if c.Weap.DropChance > 0 {
		fmt.Printf("	%s%% %s\n\n", Yellow(strconv.Itoa(c.Weap.DropChance)), Cyan("drop chance"))
	}
	Continue()
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
	fmt.Printf("%s", ItemC(c.Armor.Name))
	fmt.Printf("%s", Yellow(" T:"))
	fmt.Printf("%s", ItemC(c.Trink.Name))
	fmt.Println()
	Meter(c.Hp.Val, c.MxHp.Val, c.FoeMaxHit, "Health", "█", "hero")
	// interesting palate of ascii for perusing
	//░▒█░   ░ ████▓▒░░ ████▓▒░░▓█  ▀█▓ ▓█   ▓██▒░██▓ ▒██▒░▓█  ▀█▓ ▓█   ▓██▒███████▒
}

// StatCost cost increases with the fibonacci number position.
// The average stat is compared with the position in the fibonacci sequence
// then the product of the postion and the multiplier var = Stat Cost
func StatCost(c Agent) int {
	m := 4 // multiplier
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
		fmt.Printf("[H[J")
		time.Sleep(1000 * time.Millisecond)

		for _, c := range string(fmt.Sprintf("\n A mystical light shines down on %s's lifeless corpse.\n\n", c.Name)) {
			fmt.Printf(Blue(string(c)))
			time.Sleep(20 * time.Millisecond)
		}
		time.Sleep(1000 * time.Millisecond)
		for _, c := range string(fmt.Sprintf(" A sulfurous effluvium expands from the body.\n\n")) {
			fmt.Printf(Blue(string(c)))
			time.Sleep(20 * time.Millisecond)
		}
		time.Sleep(1000 * time.Millisecond)
		for _, c := range string(fmt.Sprintf(" %s takes a gasping breath, and lives", c.Name)) {
			fmt.Printf(Blue(string(c)))
			time.Sleep(20 * time.Millisecond)
		}
		for t := 0; t < 5; t++ {
			fmt.Printf(Blue("."))
			time.Sleep(150 * time.Millisecond)
		}
		fmt.Println("\n")

		c.Hp.Val = c.MxHp.Val
		c.Dead = false
		c.Save()
		Continue()
	} else {
		return
	}
}
func MakeMonster(c *Agent) Agent {

	divisor := 2

	power := (c.Str.Val + c.Int.Val + c.Dex.Val + c.MxHp.Val) / divisor
	health := Roll(2, power)

	var Monster = Agent{
		// Name and Description
		Name:        "Shoggoth",
		Description: "is a rancid writhing ball of organic components.",
		// Stats
		Str: Stat{"Strength", Roll(3, power)},
		Int: Stat{"Intelligence", Roll(3, power/2)},
		Dex: Stat{"Dexterity", Roll(3, power)},
		// Health and Wellness
		MxHp: Stat{"Max Health", health},
		Hp:   Stat{"Current Health", health},
		Dead: false,
		// Equiped items
		Weap:  MakeWeapon(c),
		Armor: Empty,
		Trink: Empty,
		// Inventory
		Inv: []Item{},
		Art: "\n" +
			`           .ooO8O80^^^^o_                        ` + "\n" +
			`           '^""\0... 0 .^^\                      ` + "\n" +
			`          .'    "\0. .0  0.0\                    ` + "\n" +
			`           .       \0..0' '"Y8\.                 ` + "\n" +
			`                   .\0..0' '0'88\.    ,          ` + "\n" +
			`                   ..\0..0'  0'88\.   :8b        ` + "\n" +
			`                   ..-\0..0' ,0'88\.   88b       ` + "\n" +
			`                 ..--no\0..0' '0'"8\.  Y8"8.     ` + "\n" +
			`               ...-noo.ob0..0. b ''8\  :8 Y8.    ` + "\n" +
			`           ,oooooooooobb.bYo :.'b'b'8;  8  8b    ` + "\n" +
			`    ,ood8P0"0"0"0"88888ooi 8b Y.:b:b:8;,8 ,:8.   ` + "\n" +
			`,od88888bo 0'0,  0 0"0"888o'8b'8 Y.8.88d8 : 8;   ` + "\n" +
			`"""""""""""8oo',0   0.0  ""888b8b:8db8888 d :8 :;` + "\n" +
			`          d8888boP 0 "Y88o. ""Y8888888888 8 d8.88` + "\n" +
			`        o""""888888o''o'"88bood8888888888:8,;8888` + "\n",
	}
	return Monster
}

func MakeWeapon(c *Agent) Item {
	// divisor: larger divisor makes easier weapon
	divisor := 8
	power := (c.Str.Val + c.Int.Val + c.Dex.Val + c.MxHp.Val) / divisor

	var Tentacle = Item{
		Name:        "Tentacle",
		Description: "a writhing limb of seeking flesh!",
		Slot:        "Weapon",
		//Affects:     []Affect{},
		Attack: Roll(1, power),
		Damage: Roll(2, power),
		Crit:   1,
	}
	return Tentacle
}
