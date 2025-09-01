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

// Quantile calculates the q-th quantile (0 <= q <= 1) for []T or *list.List
// q=0.25 -> 25th percentile, q=0.5 -> median, q=0.75 -> 75th percentile
func Quantile[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any, q float64) float64 {
	if q < 0 || q > 1 {
		panic("quantile must be between 0 and 1")
	}

	var values []float64

	switch v := input.(type) {
	case []T: // Handle slice
		for _, item := range v {
			values = append(values, float64(item))
		}

	case *list.List: // Handle container/list.List
		for e := v.Front(); e != nil; e = e.Next() {
			item := e.Value.(T)
			values = append(values, float64(item))
		}

	default:
		panic("unsupported input type: must be []T or *list.List")
	}

	if len(values) == 0 {
		return 0
	}

	sort.Float64s(values)

	// Rank-based interpolation (R-7 method, same as NumPy / R default)
	pos := q * float64(len(values)-1)
	lower := int(pos)
	upper := lower + 1
	if upper >= len(values) {
		return values[lower]
	}
	weight := pos - float64(lower)

	return values[lower]*(1-weight) + values[upper]*weight
}

// Quartile calculates the quartiles for []T or *list.List
func Quartile[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) []float64 {
	vals := []float64{}

	for i := 0.0; i <= 1.0; i += 0.25 {
		vals = append(vals, float64(Quantile[T](input, i)))
	}

	return vals
}

// Decile calculates the deciles for []T or *list.List
func Decile[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) []float64 {
	vals := []float64{}

	for i := 0.0; i <= 1.0; i += 0.1 {
		vals = append(vals, float64(Quantile[T](input, i)))
	}

	return vals
}

// Percentile calculates the percentiles for []T or *list.List
func Percentile[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) []float64 {
	vals := []float64{}

	for i := 0.0; i <= 1.0; i += 0.01 {
		vals = append(vals, float64(Quantile[T](input, i)))
	}

	return vals
}

// Range calculates the range (max-min) for []T or *list.List
// returns -1 if input is empty
func Range[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) float64 {
	var values []float64

	switch v := input.(type) {
	case []T: // Handle slice
		for _, item := range v {
			values = append(values, float64(item))
		}

	case *list.List: // Handle container/list.List
		for e := v.Front(); e != nil; e = e.Next() {
			item := e.Value.(T)
			values = append(values, float64(item))
		}

	default:
		panic("unsupported input type: must be []T or *list.List")
	}

	sort.Float64s(values)

	if len(values) == 0 {
		return -1
	}

	return values[len(values)-1] - values[0]
}

// InterquartileRange calculates the inter-quartile range (Q3-Q1) for []T or *list.List
func InterquartileRange[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) float64 {
	return Quantile[T](input, 0.75) - Quantile[T](input, 0.25)
}
