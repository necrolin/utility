// Utility functions for generating random numbers

package random

import (
	"math/rand"
	"time"
)

// Init a random seed
func Init() {
	rand.Seed(time.Now().UnixNano())
}

// Range return a []int slice containing numbers 'from'-'to'
func Range(from, to int) []int {
	nums := make([]int, 0)
	for i := from; i <= to; i++ {
		nums = append(nums, i)
	}
	return nums
}

// Rand returns an int between min and max. [min, max]
// Both min and max are included
func Rand(min, max int) int {
	return min + rand.Int()%(max-min+1)
}

// Roll imitates the roll of a die
func Roll() int {
	return Rand(1, 6)
}
