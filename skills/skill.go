package skills

import (
	"fmt"

	. "github.com/eyetoe/foobarbaz/agents"
	. "github.com/eyetoe/foobarbaz/colors"
	. "github.com/eyetoe/foobarbaz/util"
)

type Skill struct {
	Name        string
	Prob        int // probability 1-100 %
	Description string
}

var Door = Skill{
	Name:        "to force a door open.",
	Prob:        80,
	Description: "You may be able to force the door to open with sheer strength.",
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
