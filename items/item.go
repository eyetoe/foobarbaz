package items

import (
	"fmt"
	"strconv"

	"github.com/eyetoe/foobarbaz/affects"
	. "github.com/eyetoe/foobarbaz/colors"
)

type Item struct {
	Name        string
	Description string
	Slot        string
	Affects     []affects.Affect
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

func (i *Item) Display() {
	fmt.Printf("%s\n", Blue(i.Name))
	fmt.Printf("	Slot: %s, Attack: %s, Damage: %s\n", Yellow(i.Slot), Red(strconv.Itoa(i.Attack)), Red(strconv.Itoa(i.Damage)))
	fmt.Printf("	%s\n", Yellow(i.Description))

}
