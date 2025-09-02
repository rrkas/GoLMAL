package tests

import (
	"fmt"

	dist "github.com/rrkas/GoLMAL/stats/distributions/discrete"
)

func testBernoulli() {
	b, _ := dist.Bernoulli(0.3)
	fmt.Println("P:", b.P)
	for k := -1; k <= 1; k++ {
		fmt.Println()
		func(k int) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("panic occurred:", err)
				}
			}()

			fmt.Println("k:", k)
			fmt.Println("PMF:", b.ProbabilityMassFunction(k))
			fmt.Println("CDF:", b.CumulativeDensityFunction(k))
		}(k)
	}
}

func TestDistributions() {
	testBernoulli()
}
