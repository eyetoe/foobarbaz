package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 0; i < 10000; i++ {
		fmt.Println(Roll())
	}
}
func Roll() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(100)
	//fmt.Println(num)
	return num + 1

}
