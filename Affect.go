package main

type Affect struct {
	Name string
	// The long description seen when using 'examine' keyword
	Description string
	// Chance to Proc affect
	Proc int
	// Flavor message when affect occurs
	Message string
	// modify 'To Hit' chance
	Attack int
	// passive defence
	Damage int
	// acitve defence
	Dodge int
	// damage over time
	DoT int
	// modify critical chance
	Crit int
}

func main() {

}
