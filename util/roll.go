package util

import (
	"math/rand"
	"time"
)

// Roll is used to simulate die roll.
// roll n dice of d size, add the results and divide by n
// Roll(1, 100) will give even distribution between 1 - 100
// Roll(2, 100) will push up the bell curve in the middle of the distribution.
// Roll(3, 100) even more so
func Roll(n int, d int) int {
	num := 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		num = num + r.Intn(d)
	}
	num = num / n
	return num + 1
}
