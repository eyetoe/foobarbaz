package util

import (
	"fmt"
	"strconv"

	. "github.com/eyetoe/foobarbaz/colors"
)

// Meter provides a health meter to represent health scaled to
// barWidth.
func Meter(num, max, fmax int, l, t string) {
	// calculate the actual percentage
	p := float32(num) / float32(max)

	barWidth := 50

	// prepare as percentage of barWidth columns
	percentBar := p * float32(barWidth)

	// display meter
	fmt.Printf("%s	%s	:", Yellow(l), White(strconv.Itoa(num)))
	// drawing from left to right
	for c := 1; c <= barWidth; c++ {

		if c <= int(percentBar) {
			fmt.Printf(Green(t))
		} else {
			if num <= fmax {
				fmt.Printf(Magenta(t))
			} else {
				fmt.Printf(Red(t))
			}
		}
		if c == barWidth {
			//fmt.Println(">")
			fmt.Println(":")
		}
	}

}
