package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eyetoe/foobarbaz/agents"
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

func Attack(a *agents.Agent, d *agents.Agent) (*agents.Agent, *agents.Agent) {
	ar := Roll() + a.Str
	dr := Roll() + d.Dex
	fmt.Println("Attack Roll  + Bonus =", ar)
	fmt.Println("Defence Roll + Bonus =", dr)
	// Attack wins if greater than Defence
	// But a tie goes to the Defence
	if ar > dr {
		return a, d
	} else {
		return d, a
	}

}
