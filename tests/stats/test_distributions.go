package tests

import (
	"fmt"

	"github.com/rrkas/GoLMAL/stats/distributions/discrete"
)

func testBernoulli() {
	b, _ := discrete_dists.Bernoulli(0.3)
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

	fmt.Println()

	fmt.Println("Mean:", b.Mean())
	fmt.Println("Median:", b.Median())
	fmt.Println("Mode:", b.Mode())
	fmt.Println("Variance:", b.Variance())
	fmt.Println("MedianAbsoluteDeviation:", b.MedianAbsoluteDeviation())
	fmt.Println("Skewness:", b.Skewness())
	fmt.Println("ExcessKurtosis:", b.ExcessKurtosis())
	fmt.Println("Entropy:", b.Entropy())
	fmt.Println("MomentGeneratingFunction:", b.MomentGeneratingFunction(2.3))
	fmt.Println("FisherInformation:", b.FisherInformation())
}

func TestDistributions() {
	testBernoulli()
}
