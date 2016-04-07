package main

import (
	"fmt"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/util"
)

// Contest takes two sets of structs and stats, and returns two structs
// First input struct and stat represent the first contestant
// Second input struct and stat are the second contestant
func Contest(a *Agent, as Stat, d *Agent, ds Stat) (*Agent, *Agent) {
	// roll for a and d in contest
	ar := Roll(1, 100)
	dr := Roll(1, 100)
	// bonuses
	arB := as.Val
	drB := ds.Val
	// totals
	aT := ar + arB
	dT := dr + drB

	fmt.Printf("%s rolls: %d plus Bonus: %d for Total: %d\n", a.Name, ar, arB, aT)
	fmt.Printf("%s rolls: %d plus Bonus: %d for Total: %d\n", d.Name, dr, drB, dT)
	// Attack wins if greater than Defence
	// But a tie goes to the Defence
	if aT > dT {
		return a, d
	} else if dT > aT {
		return d, a
		//todo in contest tie should not go to defender, since there is no defender
	} else {
		fmt.Println("Tie! re-rolling")
		return Contest(a, as, d, ds)

	}

}
