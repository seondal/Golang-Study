package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf(
			"제곱근은 양수여야 합니다. f:%g", f
		)
	}
	return math.Sqrt(f),nil
}