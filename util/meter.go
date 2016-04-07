package util

import (
	"fmt"
	"strconv"

	. "github.com/eyetoe/foobarbaz/colors"
)

func Meter(num, max int, l, t string) {
	p := float32(num) / float32(max)

	barWidth := 50
	// prepare as percentage of barWidth columns
	percentBar := p * float32(barWidth)

	// display meter
	fmt.Printf("%s	%s	:", Yellow(l), White(strconv.Itoa(num)))
	for c := 1; c <= barWidth; c++ {

		if c <= int(percentBar) {
			fmt.Printf(Green(t))
		} else {

			fmt.Printf(Red(t))
		}
		if c == barWidth {
			//fmt.Println(">")
			fmt.Println(":")
		}
	}

}
