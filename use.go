package main

import (
	"fmt"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/items"
	. "github.com/eyetoe/foobarbaz/util"
)

func Use(c *Agent, p Item) string {
	var textOut string
	if c.MaxHealth.Val > c.Health.Val {

		// remove the potion from inventory
		for i, _ := range c.Inv {

			if c.Inv[i].Name == p.Name {
				c.Inv[i] = c.Inv[len(c.Inv)-1]
				c.Inv = c.Inv[:len(c.Inv)-1]

				// random health up to MaxHealth, noralized by 2 rolls / 2
				h := Roll(2, c.MaxHealth.Val)

				textOut = textOut + fmt.Sprintf(Cyan("You drink the potion.\n"))
				textOut = textOut + fmt.Sprintf(Yellow("You heal %d health!\n"), h)
				c.AdjHealth(h)
				c.Save()
				return textOut
			}
		}
		textOut = textOut + fmt.Sprintf(Red("\nYou don't have any potions.\n\n"))
		return textOut
	}
	textOut = textOut + fmt.Sprintf(Red("\nYour health is already full.\n\n"))
	return textOut

}
