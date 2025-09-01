package test

import (
	"container/list"
	"fmt"

	"github.com/rrkas/GoLMAL/stats"
)


func TestFrequency()  {
	lst_int := list.New()
	lst_int.PushBack(1)
	lst_int.PushBack(1)
	lst_int.PushBack(2)
	
	fmt.Println(lst_int)
	fmt.Println(stats.Frequency[int](lst_int))


	slice_int:= []int{1,1,4,4,5}
	fmt.Println(slice_int)
	fmt.Println(stats.Frequency[int](slice_int))
}







