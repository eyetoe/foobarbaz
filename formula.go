package main

import . "github.com/eyetoe/foobarbaz/agents"

// StatCost cost increases with the fibonacci number position.
// The average stat is compared with the position in the fibonacci sequence
// then the product of the postion and the multiplier var = Stat Cost
func StatCost(c *Agent) int {
	m := 5 // multiplier
	tStat := c.Str.Val + c.Int.Val + c.Dex.Val + c.Hp.Val
	avStat := tStat / 4

	switch {
	case avStat <= 1:
		return 1 * m
	case avStat <= 2:
		return 2 * m
	case avStat <= 3:
		return 3 * m
	case avStat <= 5:
		return 5 * m
	case avStat <= 8:
		return 8 * m
	case avStat <= 13:
		return 13 * m
	case avStat <= 21:
		return 21 * m
	case avStat <= 34:
		return 34 * m
	case avStat <= 55:
		return 55 * m
	case avStat <= 89:
		return 89 * m
	case avStat <= 144:
		return 144 * m
	default:
		return 200 * m
	}
}
