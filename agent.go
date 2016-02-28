package main

type Agent struct {
	Name string
	// 3 base attributes Str, Int, Dex
	Str int
	Int int
	Dex int
	// Max Health and Health
	MxHp int
	Hp   int
	// 3 item slots: Weapon, Armor, Trinket
	Weap  string
	Armor string
	Trink string
	Atk1  string
	Atk2  string
	Atk3  string
	Dead  bool
}

func main() {
}
