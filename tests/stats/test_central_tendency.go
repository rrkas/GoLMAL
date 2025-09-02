package tests

import (
	"container/list"
	"fmt"

	"github.com/rrkas/GoLMAL/stats"
	"github.com/rrkas/GoLMAL/tests/data"
)

func TestCentralTendencies() {
	lst_int := list.New()
	slice_int := data.RandomStatsData()
	for _, e := range slice_int {
		lst_int.PushBack(e)
	}

	fmt.Println(lst_int)
	fmt.Println(slice_int)

	// fmt.Println()
	// fmt.Println(stats.ArithmeticMean[int](lst_int))
	// fmt.Println(stats.ArithmeticMean[int](slice_int))

	// fmt.Println()
	// fmt.Println(stats.Median[int](lst_int))
	// fmt.Println(stats.Median[int](slice_int))

	// fmt.Println()
	// fmt.Println(stats.Mode[int](lst_int))
	// fmt.Println(stats.Mode[int](slice_int))

	// fmt.Println()
	// fmt.Println(stats.Quantile[int](slice_int, 0))
	// fmt.Println(stats.Quantile[int](slice_int, 0.25))
	// fmt.Println(stats.Quantile[int](slice_int, 0.5))
	// fmt.Println(stats.Quantile[int](slice_int, 0.75))
	// fmt.Println(stats.Quantile[int](slice_int, 1))

	// fmt.Println()
	// fmt.Println(stats.Quartile[int](slice_int))

	// fmt.Println()
	// fmt.Println(stats.Decile[int](slice_int))

	// fmt.Println()
	// fmt.Println(stats.Percentile[int](slice_int))

	// fmt.Println()
	// fmt.Println(stats.Range[int](slice_int))

	// fmt.Println()
	// fmt.Println(stats.InterquartileRange[int](slice_int))

	// fmt.Println()
	// fmt.Println(stats.MeanAbsoluteDeviation[int](slice_int))

	// fmt.Println()
	// fmt.Println(stats.MeanSquaredDeviation[int](slice_int))

	// fmt.Println()
	// fmt.Println(stats.StandardDeviation[int](slice_int))

	// fmt.Println()
	// fmt.Println(stats.Skewness[int](slice_int))

	// fmt.Println()
	// fmt.Println(stats.Kurtosis[int](slice_int))
}
