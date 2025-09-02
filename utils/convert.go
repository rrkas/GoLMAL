package utils

import "container/list"

func ToFloat64s[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) []float64 {
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

	return values
}
