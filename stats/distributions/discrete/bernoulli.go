package discrete_dists

import (
	"fmt"
	"math"
)

type bernoulli struct {
	P float64
}

func Bernoulli(p float64) (*bernoulli, error) {
	if p < 0 || p > 1 {
		return nil, fmt.Errorf("probability (P) must be between 0 and 1, got %f", p)
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

func (b bernoulli) Mean() float64 {
	return b.P
}

func (b bernoulli) Median() []float64 {
	if b.P < 0.5 {
		return []float64{0}
	}
	if b.P > 0.5 {
		return []float64{1}
	}
	return []float64{0, 1}
}

func (b bernoulli) Mode() []float64 {
	if b.P < 0.5 {
		return []float64{0}
	}
	if b.P > 0.5 {
		return []float64{1}
	}
	return []float64{0, 1}
}

func (b bernoulli) Variance() float64 {
	q := 1 - b.P
	return b.P * q
}

func (b bernoulli) MedianAbsoluteDeviation() float64 {
	q := 1 - b.P
	return 2 * b.P * q
}

func (b bernoulli) Skewness() float64 {
	p := b.P
	q := 1 - b.P
	return (q - p) / math.Pow(p*q, 0.5)
}

func (b bernoulli) ExcessKurtosis() float64 {
	p := b.P
	q := 1 - b.P
	return (1 - 6*p*q) / (p * q)
}

func (b bernoulli) Entropy() float64 {
	p := b.P
	q := 1 - b.P
	return -q*math.Log(q) - p*math.Log(p)
}

func (b bernoulli) MomentGeneratingFunction(t float64) float64 {
	p := b.P
	q := 1 - b.P
	return q + p*math.Pow(math.E, t)
}

func (b bernoulli) FisherInformation() float64 {
	p := b.P
	q := 1 - b.P
	return 1 / (p * q)
}
