package items

import . "github.com/eyetoe/foobarbaz/affects"

var Potion = Item{
	Name:        "Potion",
	Description: "A mysterious milky blue elixir!",
	Slot:        "Inventory",
	Affects:     []Affect{},
	Attack:      0,
	Damage:      0,
	Crit:        0,
}

//func (p Potion) Use(a *Agent) {
//	if c.MxHp.Val > c.Hp.Val {

//		//c.Inv = append(c.Inv, Potion)
//		// remove the potion from inventory
//		c.Inv[i] = c.Inv[len(c.Inv)-1]
//		c.Inv = c.Inv[:len(c.Inv)-1]
//
//		// randome health up to MxHp, noralized by 2 rolls / 2
//		h := Roll(2, c.MxHp.Val)
//
//		textOut = textOut + fmt.Sprintf(Yellow("\nYou heal %d health!\n\n"), h)
//		c.AdjHp(h)
//		c.Save()
//	}
//
//}

//&& c.MxHp.Val+30 >= Roll(1, 100)
