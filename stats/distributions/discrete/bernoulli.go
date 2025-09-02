package stats

import "fmt"

type bernoulli struct {
	P float64
}

func Bernoulli(p float64) (*bernoulli, error) {
	if p < 0 || p > 1 {
		return nil, fmt.Errorf("probability must be between 0 and 1, got %f", p)
	}
	return &bernoulli{P: p}, nil
}

func (b bernoulli) ProbabilityMassFunction(k int) float64 {
	switch k {
	case 0:
		return 1 - b.P
	case 1:
		return b.P
	default:
		panic(fmt.Sprintf("k(%v) must be 0/1", k))
	}
}

func (b bernoulli) CumulativeDensityFunction(k int) float64 {
	if k < 0 {
		return 0
	}
	if k < 1 {
		return 1 - b.P
	}
	return 1
}
