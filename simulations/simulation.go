package simulations

import (
	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/util"
)

func Odds(c *Agent, f *Agent) int {
	charWins := 0
	foeWins := 0
	for i := 0; i < 100; i++ {
		if SimFight(*c, *f) == true {
			//fmt.Println("Character won!")
			charWins++
		} else {
			//fmt.Println("Character lost!")
			foeWins++

		}
	}
	return charWins
}

// SimFight simulates combat between character and a proposed foe.
// Outcome predition is output as a percentage chance for the
// character to win. Experience point rewards should be based on
// this metric where fights with long odds give greater experience.
func SimFight(c Agent, f Agent) bool {

	// operate on a copy of structs for the simulation
	var x, y Agent
	x = c
	y = f

	//fmt.Println(x.Name, y.Name)
	for {

		// Character Attacks
		winner, loser := SimAttack(&x, &y)
		if winner.Name == x.Name {
			hp := Roll(2, winner.Weap.Damage)
			//fmt.Println(loser.Name, "takes", hp)
			loser.AdjHp(0 - hp)
			if loser.Dead == true {
				return true
			}
		}

		// Foe Attacks Second
		winner, loser = SimAttack(&y, &x)
		if winner.Name == y.Name {
			hp := Roll(2, winner.Weap.Damage)
			//fmt.Println(loser.Name, "takes", hp)
			loser.AdjHp(0 - hp)
			if loser.Dead == true {
				return false
			}
		}
	}

	//	return
}

func SimAttack(a *Agent, d *Agent) (*Agent, *Agent) {

	// roll for attacker and defender
	ar := Roll(2, 100)
	dr := Roll(2, 100)
	// bonuses
	arB := a.Str.Val + a.Weap.Attack
	drB := d.Dex.Val
	// totals
	aT := ar + arB
	dT := dr + drB

	// Attack wins if greater than Defence
	// But a tie goes to the Defence

	if aT > dT {
		return a, d
	} else {
		return d, a
	}

}
