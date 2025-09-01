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
		numbers[i] = rand.Intn(20) // random value between 0–19
	}

	return numbers
}
