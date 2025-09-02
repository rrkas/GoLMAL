package stats

import (
	"container/list"
	"math"
	"sort"

	"github.com/rrkas/GoLMAL/utils"
)

// ArithmeticMean calculates the mean for []T or *list.List
func ArithmeticMean[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) (float64, bool) {
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
		return 0, false
	}

	return sum / float64(count), true
}

// Median calculates the median for []T or *list.List
func Median[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) (float64, bool) {
	values := utils.ToFloat64s[T](input)

	if len(values) == 0 {
		return 0, false
	}

	sort.Float64s(values)
	mid := len(values) / 2

	if len(values)%2 == 0 {
		return (values[mid-1] + values[mid]) / 2, true
	}

	return values[mid], true
}

// Mode calculates the mode(s) for []T or *list.List
func Mode[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) ([]float64, bool) {
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
		return nil, false
	}

	var modes []float64
	for val, count := range freq {
		if count == maxCount {
			modes = append(modes, val)
		}
	}

	return modes, true
}

// Quantile calculates the q-th quantile (0 <= q <= 1) for []T or *list.List
// q=0.25 -> 25th percentile, q=0.5 -> median, q=0.75 -> 75th percentile
func Quantile[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any, q float64) (float64, bool) {
	if q < 0 || q > 1 {
		panic("quantile must be between 0 and 1")
	}

	values := utils.ToFloat64s[T](input)

	if len(values) == 0 {
		return 0, false
	}

	sort.Float64s(values)

	// Rank-based interpolation (R-7 method, same as NumPy / R default)
	pos := q * float64(len(values)-1)
	lower := int(pos)
	upper := lower + 1
	if upper >= len(values) {
		return values[lower], true
	}
	weight := pos - float64(lower)

	return values[lower]*(1-weight) + values[upper]*weight, true
}

// Quartile calculates the quartiles for []T or *list.List
func Quartile[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) ([]float64, bool) {
	vals := []float64{}

	for i := 0; i <= 4; i++ {
		val, success := Quantile[T](input, float64(i)/4)
		if !success {
			return nil, false
		}
		vals = append(vals, float64(val))
	}

	return vals, true
}

// Decile calculates the deciles for []T or *list.List
func Decile[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) ([]float64, bool) {
	vals := []float64{}

	for i := 0; i <= 10; i++ {
		val, success := Quantile[T](input, float64(i)/10)
		if !success {
			return nil, false
		}
		vals = append(vals, float64(val))
	}

	return vals, true
}

// Percentile calculates the percentiles for []T or *list.List
func Percentile[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) ([]float64, bool) {
	vals := []float64{}

	for i := 0; i <= 100; i++ {
		val, success := Quantile[T](input, float64(i)/100)
		if !success {
			return nil, false
		}
		vals = append(vals, float64(val))
	}

	return vals, true
}

// Range calculates the range (max-min) for []T or *list.List
// returns -1 if input is empty
func Range[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) (float64, bool) {
	values := utils.ToFloat64s[T](input)

	if len(values) == 0 {
		return 0, false
	}

	sort.Float64s(values)

	return values[len(values)-1] - values[0], true
}

// InterquartileRange calculates the inter-quartile range (Q3-Q1) for []T or *list.List
func InterquartileRange[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) (float64, bool) {
	q3, success := Quantile[T](input, 0.75)
	if !success {
		return 0, false
	}
	q1, success := Quantile[T](input, 0.25)
	if !success {
		return 0, false
	}
	return q3 - q1, true
}

// MeanAbsoluteDeviation calculates Mean Absolute Deviation for []T or *list.List
func MeanAbsoluteDeviation[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) (float64, bool) {
	s := float64(0)
	n := float64(0)

	mean, success := ArithmeticMean[T](input)
	if !success {
		return 0, false
	}

	switch v := input.(type) {
	case []T: // Handle slice
		for _, item := range v {
			s += math.Abs(float64(item) - mean)
			n += 1
		}

	case *list.List: // Handle container/list.List
		for e := v.Front(); e != nil; e = e.Next() {
			item := e.Value.(T)
			s += math.Abs(float64(item) - mean)
			n += 1
		}

	default:
		panic("unsupported input type: must be []T or *list.List")
	}

	return s / n, true
}

// MeanSquaredDeviation calculates Mean Squared Deviation for []T or *list.List
func MeanSquaredDeviation[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) (float64, bool) {
	s := float64(0)
	n := float64(0)

	mean, success := ArithmeticMean[T](input)
	if !success {
		return 0, false
	}

	switch v := input.(type) {
	case []T: // Handle slice
		for _, item := range v {
			s += math.Pow(float64(item)-mean, 2)
			n += 1
		}

	case *list.List: // Handle container/list.List
		for e := v.Front(); e != nil; e = e.Next() {
			item := e.Value.(T)
			s += math.Pow(float64(item)-mean, 2)
			n += 1
		}

	default:
		panic("unsupported input type: must be []T or *list.List")
	}

	return s / n, true
}

// Variance calculates variance for []T or *list.List
func Variance[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) (float64, bool) {
	variance, success := MeanSquaredDeviation[T](input)
	if !success {
		return 0, false
	}

	return variance, true
}

// StandardDeviation calculates standard deviation for []T or *list.List
func StandardDeviation[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) (float64, bool) {
	variance, success := MeanSquaredDeviation[T](input)
	if !success {
		return 0, false
	}

	return math.Pow(variance, 0.5), true
}

// Skewness calculates skewness for []T or *list.List
func Skewness[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) (float64, bool) {
	mean, success := ArithmeticMean[T](input)
	if !success {
		return 0, false
	}

	std_dev, success := StandardDeviation[T](input)
	if !success {
		return 0, false
	}

	s := float64(0.0)

	values := utils.ToFloat64s[T](input)

	for _, v := range values {
		s += math.Pow((v-mean)/std_dev, 3)
	}

	return s / float64(len(values)), true
}

// Kurtosis calculates kurtosis for []T or *list.List
func Kurtosis[T ~int | ~int32 | ~int64 | ~float32 | ~float64](input any) (float64, bool) {
	mean, success := ArithmeticMean[T](input)
	if !success {
		return 0, false
	}

	std_dev, success := StandardDeviation[T](input)
	if !success {
		return 0, false
	}

	s := float64(0.0)

	values := utils.ToFloat64s[T](input)

	for _, v := range values {
		s += math.Pow((v-mean)/std_dev, 4)
	}

	s -= 3

	return s / float64(len(values)), true
}
