package main

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/eyetoe/foobarbaz/agents"
)

func Roll() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(100)
	//fmt.Println(num)
	return num + 1

}

func Skill(a Agent, s Stat, d Stat) bool {
	fmt.Println("Skill Check!")
	// need = difficulty - skill
	n := d.Val - s.Val
	r := Roll()
	fmt.Printf("%s attempts to use %s skill.\n", a.Name, s.Name)
	fmt.Printf("Difficulty(%d) minus Skill(%d) := Roll %d or higher to succeed.\n", d.Val, s.Val, n)
	fmt.Printf("Roll...... %d !\n", r)

	if r >= n {
		fmt.Printf("%s Succeeds!\n", a.Name)
		return true
	} else {
		fmt.Printf("%s Fails!\n", a.Name)
		return false
	}

}

// Attack takes two structs and returns two structs.
// First input struct is the 'attacker'
// Second input struct is the 'defender'
// First output struct is the 'winner'
// Second output struct is the 'loser'
func Attack(a *Agent, d *Agent) (*Agent, *Agent) {
	// roll for attacker and defender
	ar := Roll()
	dr := Roll()
	// bonuses
	arB := a.Str.Val
	drB := d.Str.Val
	// totals
	aT := ar + arB
	dT := dr + drB

	fmt.Printf("Attack roll: %d plus Bonus: %d for Total: %d\n", ar, arB, aT)
	fmt.Printf("Defence roll: %d plus Bonus: %d for Total: %d\n", dr, drB, dT)
	// Attack wins if greater than Defence
	// But a tie goes to the Defence
	if ar > dr {
		return a, d
	} else {
		return d, a
	}

}

// Contest takes two sets of structs and stats, and returns two structs
// First input struct and stat represent the first contestant
// Second input struct and stat are the second contestant
func Contest(a *Agent, as Stat, d *Agent, ds Stat) (*Agent, *Agent) {
	// roll for a and d in contest
	ar := Roll()
	dr := Roll()
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
