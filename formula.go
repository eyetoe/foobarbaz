package main

import (
	"math/rand"
	"time"

	. "github.com/eyetoe/foobarbaz/agents"
)

func Spawn() Agent {
	rand.Seed(time.Now().UTC().UnixNano())
	monsters := []Agent{
		// Add monsters here to be included in random spawn
		Minotaur,
		Coyote,
		Phantom,
	}
	return monsters[rand.Intn(len(monsters))]
	//monster := monsters[rand.Intn(len(monsters))]
	//fmt.Println(monster.Name, monster.Description)

}
