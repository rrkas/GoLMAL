package data

import (
	"math/rand"
)

func RandomStatsData() []int {
	// Random length between 5 and 100
	length := rand.Intn(96) + 5

	// Create slice with random values
	numbers := make([]int, length)
	for i := range numbers {
		numbers[i] = -20 + rand.Intn(40) // random value between [-20, 19]
	}

	return numbers
}
