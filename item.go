package main

type Item struct {
	Name        string
	Description string
	Slot        string
	Affects     []Affect
	// modify 'To Hit' chance
	Attack int
	// modify damage dealt if hit
	Damage int
	// passive defence. heavier armor, Str = better Defence
	Defence int
	// active defence. lighter armor, Dex = better Dodge
	Dodge int

	// dammage over time adjustment. Bleed, Poison, Fire, etc
	DoT int
	// modify critical chance.
	Crit int
}

func main() {
}
