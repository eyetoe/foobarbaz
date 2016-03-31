package meters

import (
	"fmt"
	"strconv"

	. "github.com/eyetoe/foobarbaz/colors"
)

func Meter(num, max int, l, t string) {
	p := float32(num) / float32(max)

	// prepare as percentage of 40 col
	percent40 := p * 40

	// display meter
	fmt.Printf("%s	%s	:", Yellow(l), White(strconv.Itoa(num)))
	for c := 1; c <= 40; c++ {

		if c <= int(percent40) {
			fmt.Printf(Green(t))
		} else {

			fmt.Printf(Red(t))
		}
		if c == 40 {
			//fmt.Println(">")
			fmt.Println(":")
		}
	}

}
