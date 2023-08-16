package main

import (
	"math"
)

func TwoCrystalBalls(breaks []bool) int {

	zeroIndexLength := float64(len(breaks) - 1)
	jmp := int(math.Sqrt(zeroIndexLength))
	i := jmp

	// intentionally breaking the crystal ball >:)
	for ; i < len(breaks); i += jmp {
		if breaks[i] {
			break
		}
	}

	// breakpoint
	i -= jmp

	for j := 0; j < jmp && i < len(breaks); j++ {
		i++
		if breaks[i] {
			return i
		}

	}

	return -1
}
