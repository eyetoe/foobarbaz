package main

import (
	"fmt"
	"strconv"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/colors"
	//	. "github.com/eyetoe/foobarbaz/items"
	. "github.com/eyetoe/foobarbaz/util"
)

func Inventory(c *Agent) {

	ClearScreen()
	c.StatusBar()

	//	fmt.Println(Yellow("\n" + `                            _.--.         ` + "\n" +
	//		`                        _.-'_:-'||        ` + "\n" +
	//		`                    _.-'_.-::::'||        ` + "\n" +
	//		`               _.-:'_.-::::::'  ||        ` + "\n" +
	//		`             .''-.-:::::::'     ||        ` + "\n" +
	//		`            /.'';|:::::::'      ||_       ` + "\n" +
	//		`           ||   ||::::::'     _.;._'-._   ` + "\n" +
	//		`           ||   ||:::::'  _.-!oo @.!-._'-.` + "\n" +
	//		`           \'.  ||:::::.-!()oo @!()@.-'_.|` + "\n" +
	//		`            '.'-;|:.-'.&$@.& ()$%-'o.'\U||` + "\n" +
	//		`              '>'-.!@%()@'@_%-'_.-o _.|'||` + "\n" +
	//		`               ||-._'-.@.-'_.-' _.-o  |'||` + "\n" +
	//		`               ||=[ '-._.-\U/.-'    o |'||` + "\n" +
	//		`               || '-.]=|| |'|      o  |'||` + "\n" +
	//		`               ||      || |'|        _| ';` + "\n" +
	//		`               ||      || |'|    _.-'_.-' ` + "\n" +
	//		`               |'-._   || |'|_.-'_.-'     ` + "\n" +
	//		`            jgs '-._'-.|| |' '_.-'        ` + "\n" +
	//		`                    '-.||_/.-'            ` + "\n"))

	//c.Inv = append(c.Inv, Hatchet)
	//c.Inv = append(c.Inv, Knife)
	//c.Inv = append(c.Inv, Dagger)
	//c.Inv = append(c.Inv, Staff)
	//c.Inv = append(c.Inv, Tack)

	fmt.Println(YellowU("Inventory\n"))

	for a, b := range c.Inv {
		//fmt.Println(GreenU(a), WhiteU(b.Name))
		fmt.Printf("%s:	%s %s Attack: +%s Damage: 1d%s\n", GreenU(strconv.Itoa(a)), WhiteU(b.Name), Blue(b.Slot), Red(strconv.Itoa(b.Attack)), Red(strconv.Itoa(b.Damage)))
		fmt.Println("	", Blue(b.Description), "\n")

	}
	Continue()

}
