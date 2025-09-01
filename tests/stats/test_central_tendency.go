package tests

import (
	"container/list"
	"fmt"

	"github.com/rrkas/GoLMAL/stats"
	"github.com/rrkas/GoLMAL/tests/data"
)

func TestArithmeticMean() {
	lst_int := list.New()
	slice_int := data.RandomStatsData()
	for _, e := range slice_int {
		lst_int.PushBack(e)
	}

	fmt.Println(lst_int)
	fmt.Println(slice_int)

	fmt.Println()

	fmt.Println(stats.ArithmeticMean[int](lst_int))
	fmt.Println(stats.ArithmeticMean[int](slice_int))
}

func TestMedian() {
	lst_int := list.New()
	slice_int := data.RandomStatsData()
	for _, e := range slice_int {
		lst_int.PushBack(e)
	}

	fmt.Println(lst_int)
	fmt.Println(slice_int)

	fmt.Println()

	fmt.Println(stats.Median[int](lst_int))
	fmt.Println(stats.Median[int](slice_int))
}

func TestMode() {
	lst_int := list.New()
	slice_int := data.RandomStatsData()
	for _, e := range slice_int {
		lst_int.PushBack(e)
	}

	fmt.Println(lst_int)
	fmt.Println(slice_int)

	fmt.Println()

	fmt.Println(stats.Mode[int](lst_int))
	fmt.Println(stats.Mode[int](slice_int))
}

func TestQuantile() {
	lst_int := list.New()
	// slice_int := data.RandomStatsData()
	slice_int := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, e := range slice_int {
		lst_int.PushBack(e)
	}

	fmt.Println(lst_int)
	fmt.Println(slice_int)

	fmt.Println()

	fmt.Println(stats.Quantile[int](lst_int, 0))
	fmt.Println(stats.Quantile[int](lst_int, 0.25))
	fmt.Println(stats.Quantile[int](lst_int, 0.5))
	fmt.Println(stats.Quantile[int](lst_int, 0.75))
	fmt.Println(stats.Quantile[int](lst_int, 1))

	fmt.Println()

	fmt.Println(stats.Quartile[int](slice_int))
}
