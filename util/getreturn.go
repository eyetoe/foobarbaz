package util

import (
	"bufio"
	"os"
)

// GetReturn returns a single string character after user enters newline
func GetReturn() (ascii string) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return string([]byte(input)[0])
}
