package stats

import (
	"container/list"
	"sort"
)

// ArithmeticMean calculates the mean for []T or *list.List
func ArithmeticMean[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) float64 {
	var sum float64
	var count int

	switch v := input.(type) {
	case []T: // Handle slice
		for _, item := range v {
			sum += float64(item)
			count++
		}

	case *list.List: // Handle container/list.List
		for e := v.Front(); e != nil; e = e.Next() {
			item := e.Value.(T) // type assert to T
			sum += float64(item)
			count++
		}

	default:
		panic("unsupported input type: must be []T or *list.List")
	}

	if count == 0 {
		return 0
	}

	return sum / float64(count)
}

// Median calculates the median for []T or *list.List
func Median[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) float64 {
	var values []float64

	switch v := input.(type) {
	case []T: // Handle slice
		for _, item := range v {
			values = append(values, float64(item))
		}

	case *list.List: // Handle container/list.List
		for e := v.Front(); e != nil; e = e.Next() {
			item := e.Value.(T) // type assert to T
			values = append(values, float64(item))
		}

	default:
		panic("unsupported input type: must be []T or *list.List")
	}

	if len(values) == 0 {
		return 0
	}

	sort.Float64s(values)
	mid := len(values) / 2

	if len(values)%2 == 0 {
		return (values[mid-1] + values[mid]) / 2
	}
	return values[mid]
}

// Mode calculates the mode(s) for []T or *list.List
func Mode[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) []float64 {
	freq := make(map[float64]int)
	var maxCount int

	switch v := input.(type) {
	case []T: // Handle slice
		for _, item := range v {
			val := float64(item)
			freq[val]++
			if freq[val] > maxCount {
				maxCount = freq[val]
			}
		}

	case *list.List: // Handle container/list.List
		for e := v.Front(); e != nil; e = e.Next() {
			item := float64(e.Value.(T))
			freq[item]++
			if freq[item] > maxCount {
				maxCount = freq[item]
			}
		}

	default:
		panic("unsupported input type: must be []T or *list.List")
	}

	if len(freq) == 0 {
		return nil
	}

	var modes []float64
	for val, count := range freq {
		if count == maxCount {
			modes = append(modes, val)
		}
	}

	return modes
}
