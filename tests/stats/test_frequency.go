package tests

import (
	"container/list"
	"fmt"

	"github.com/rrkas/GoLMAL/stats"
	"github.com/rrkas/GoLMAL/tests/data"
)

func TestFrequency() {
	lst_int := list.New()
	slice_int := data.RandomStatsData()
	for _, e := range slice_int {
		lst_int.PushBack(e)
	}

	fmt.Println(lst_int)
	fmt.Println(slice_int)

	fmt.Println()

	fmt.Println(stats.AbsoluteFrequency[int](lst_int))
	fmt.Println(stats.AbsoluteFrequency[int](slice_int))

	fmt.Println()

	fmt.Println(stats.RelativeFrequency[int](lst_int))
	fmt.Println(stats.RelativeFrequency[int](slice_int))

	fmt.Println()

	fmt.Println(stats.CumulativeAbsoluteFrequency[int](lst_int))
	fmt.Println(stats.CumulativeAbsoluteFrequency[int](slice_int))

	fmt.Println()

	fmt.Println(stats.CumulativeRelativeFrequency[int](lst_int))
	fmt.Println(stats.CumulativeRelativeFrequency[int](slice_int))
}
