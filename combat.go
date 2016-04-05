package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/skills"
)

func Roll(n int, d int) int {
	num := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		num = num + r.Intn(d)
	}
	num = num / n
	return num + 1
}

// Skill determines the winner in a contest of skill
// pass in an Agent, a Stat to base the skill check on, and a Difficulty
// e.g. a player would like to try to open a heavy door by force. Success
// should be based on players's strength but also on a general level of difficulty
// representing how well built the door is.
//  Skill(Char, Char.Str, 50)
func SkillCheck(a Agent, s Stat, d Skill) bool {
	fmt.Println("Skill Check!")
	// need = difficulty - skill
	n := d.Prob - s.Val
	r := Roll(1, 100)
	fmt.Printf("%s.\n %s attempts to use %s skill.\n", d.Description, a.Name, s.Name)
	fmt.Printf("Difficulty(%d) minus Skill(%d) := Roll %d or higher to succeed.\n", d.Prob, s.Val, n)
	fmt.Printf("Roll...... %d !\n", r)

	if r >= n {
		fmt.Printf("%s %s!\n", a.Name, Green("Succeeds"))
		return true
	} else {
		fmt.Printf("%s %s!\n", a.Name, Red("Fails"))
		return false
	}

}

// Calculate and apply damage to Agent
func Damage(a *Agent, d *Agent, odds int) string {

	var textOut string

	hp := Roll(1, a.Weap.Damage)

	d.AdjHp(0 - hp)

	if d.Dead == false {
		textOut = textOut + fmt.Sprintf("for %s damage. ", Red(strconv.Itoa(hp)))
		textOut = textOut + fmt.Sprintf("%s's health = %s.\n", d.Name, Red(strconv.Itoa(d.Hp.Val)))
	}

	// Monster agents don't have a save file set
	if d.File == "" && d.Dead == true {

		// reverse the percentage
		mods := 100 - odds

		percentage := float32(mods) * .01

		// reduce the drop by the reverse percentage
		exp := float32(d.ExpDrop()) * percentage

		// exp is a float32 so do math back as in
		a.Exp = a.Exp + int(exp)

		textOut = textOut + fmt.Sprintf(Green("You gain %d experience.\n"), int(exp))
		a.Save()
	}
	d.Save()
	return textOut
}

// Attack takes two structs and returns two structs.
// First input struct is the 'attacker'
// Second input struct is the 'defender'
// First output struct is the 'winner'
// Second output struct is the 'loser'
func Attack(a *Agent, d *Agent) (*Agent, *Agent, string) {

	var outText string

	// roll for attacker and defender
	ar := Roll(1, 100)
	dr := Roll(1, 100)
	// bonuses
	arB := a.Str.Val + a.Weap.Attack
	drB := d.Dex.Val
	// totals
	aT := ar + arB
	dT := dr + drB

	outText = fmt.Sprintf(Black("Attack roll: %d plus Bonus: %d for Total: %d\n"), ar, arB, aT)
	outText = outText + fmt.Sprintf(Black("Defence roll: %d plus Bonus: %d for Total: %d\n"), dr, drB, dT)

	// Attack wins if greater than Defence
	// But a tie goes to the Defence

	if aT > dT {
		outText = outText + fmt.Sprintf(Green("%s hits!\n "), a.Name)
		return a, d, outText
	} else {
		outText = outText + fmt.Sprintf(Red("%s misses!\n"), a.Name)
		return d, a, outText
	}

}

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

func Spawn() Agent {
	rand.Seed(time.Now().UTC().UnixNano())
	monsters := []Agent{
		// Add monsters here to be included in random spawn
		Spider,
		Phantom,
		Pixie,
		Kobold,
		Warlock,
		Coyote,
		Rogue,
		Minotaur,
		Lacrimosa,
		Griffon,
	}
	return monsters[rand.Intn(len(monsters))]

}
