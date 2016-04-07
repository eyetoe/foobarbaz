package util

import (
	"fmt"
)

func ClearScreen() {
	fmt.Printf("[H[J")
}
