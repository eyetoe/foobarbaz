package skills

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
