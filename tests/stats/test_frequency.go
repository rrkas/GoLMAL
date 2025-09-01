package test

import (
	"container/list"
	"fmt"

	"github.com/rrkas/GoLMAL/stats"
)

func TestFrequency() {
	lst_int := list.New()
	lst_int.PushBack(1)
	lst_int.PushBack(1)
	lst_int.PushBack(4)
	lst_int.PushBack(4)
	lst_int.PushBack(5)

	fmt.Println(lst_int)
	fmt.Println(stats.AbsoluteFrequency[int](lst_int))
	fmt.Println(stats.RelativeFrequency[int](lst_int))
	fmt.Println(stats.CumulativeAbsoluteFrequency[int](lst_int))
	fmt.Println(stats.CumulativeRelativeFrequency[int](lst_int))

	fmt.Println()

	slice_int := []int{1, 1, 4, 4, 5}
	fmt.Println(slice_int)
	fmt.Println(stats.AbsoluteFrequency[int](slice_int))
	fmt.Println(stats.RelativeFrequency[int](slice_int))
	fmt.Println(stats.CumulativeAbsoluteFrequency[int](slice_int))
	fmt.Println(stats.CumulativeRelativeFrequency[int](slice_int))
}
