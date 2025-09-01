package stats

import "container/list"

// AbsoluteFrequency calculates the frequency of elements in either a slice ([]T) or *list.List.
// It works with any comparable type (int, string, float, etc.).
func AbsoluteFrequency[T comparable](input interface{}) map[T]int {
	freq := make(map[T]int)

	switch v := input.(type) {
	case []T: // Handle slice
		for _, item := range v {
			freq[item]++
		}

	case *list.List: // Handle container/list.List
		for e := v.Front(); e != nil; e = e.Next() {
			item := e.Value.(T) // Type assert to T
			freq[item]++
		}

	default:
		panic("unsupported input type: must be []T or *list.List")
	}

	return freq
}

// CumulativeAbsoluteFrequency calculates the cumulative frequency of elements in either a slice ([]T) or *list.List.
// It works with any comparable type (int, string, float, etc.).
func CumulativeAbsoluteFrequency[T comparable](input interface{}) map[T]int {
	freq := make(map[T]int)
	count := 0

	switch v := input.(type) {
	case []T: // Handle slice
		for _, item := range v {
			count++
			freq[item] = count
		}

	case *list.List: // Handle container/list.List
		for e := v.Front(); e != nil; e = e.Next() {
			item := e.Value.(T) // Type assert to T
			count++
			freq[item] = count
		}

	default:
		panic("unsupported input type: must be []T or *list.List")
	}

	return freq
}

// RelativeFrequency calculates the relative frequency of elements in either a slice ([]T) or *list.List.
// It works with any comparable type (int, string, float, etc.).
func RelativeFrequency[T comparable](input interface{}) map[T]float64 {
	freq := make(map[T]float64)
	count := float64(0)

	switch v := input.(type) {
	case []T: // Handle slice
		for _, item := range v {
			count++
			freq[item]++
		}

	case *list.List: // Handle container/list.List
		for e := v.Front(); e != nil; e = e.Next() {
			item := e.Value.(T) // Type assert to T
			count++
			freq[item]++
		}

	default:
		panic("unsupported input type: must be []T or *list.List")
	}

	// Normalize to relative frequency
	for k, v := range freq {
		freq[k] = v / float64(count)
	}

	return freq
}

// CumulativeRelativeFrequency calculates the cumulative relative frequency of elements in either a slice ([]T) or *list.List.
// It works with any comparable type (int, string, float, etc.).
func CumulativeRelativeFrequency[T comparable](input interface{}) map[T]float64 {
	freq := make(map[T]float64)
	count := float64(0)

	switch v := input.(type) {
	case []T: // Handle slice
		for _, item := range v {
			count++
			freq[item] = count
		}

	case *list.List: // Handle container/list.List
		for e := v.Front(); e != nil; e = e.Next() {
			item := e.Value.(T) // Type assert to T
			count++
			freq[item] = count
		}

	default:
		panic("unsupported input type: must be []T or *list.List")
	}

	// Normalize to relative frequency
	for k, v := range freq {
		freq[k] = v / float64(count)
	}

	return freq
}
