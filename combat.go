package main

import (
	"fmt"
	"math/rand"
	"time"

	. "github.com/eyetoe/foobarbaz/agents"
)

//func main() {

//	for i := 0; i < 10000; i++ {
//		fmt.Println(Roll())
//	}
//}
func Roll() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(100)
	//fmt.Println(num)
	return num + 1

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
	arB := a.Str
	drB := d.Str
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
func Contest(a *Agent, as int, d *Agent, ds int) (*Agent, *Agent) {
	ar := Roll() + as
	dr := Roll() + ds
	fmt.Println(a.Name, "Roll  + Bonus =", ar)
	fmt.Println(d.Name, "Roll + Bonus =", dr)
	// Attack wins if greater than Defence
	// But a tie goes to the Defence
	if ar > dr {
		return a, d
	} else {
		return d, a
	}

}
