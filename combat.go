package main

import (
	"fmt"
	"math/rand"
)

func Roll() int {
	return rand.Intn(100)
	//fmt.Print(rand.Intn(100), ",")
}

func Contest(a int, b int) int {
	c := a + b
	//if c := a + b; c < 100 {
	if c < 100 {
		fmt.Println("c is:", c)
	}
	return c
}

//func main() {
//
//	fmt.Println("combat.go")
//	winner := Contest(8, 9)
//	fmt.Printf("contest: %d\n", winner)
//}
