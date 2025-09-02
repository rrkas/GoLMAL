package discrete_dists

import (
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
	"math"
)

type binomial struct {
	P float64
	N int
}

func Binomial(n int, p float64) (*binomial, error) {
	if n < 0 {
		return nil, fmt.Errorf("number of trials (N) must be >0, got %f", p)
	}
	if p < 0 || p > 1 {
		return nil, fmt.Errorf("probability (P) must be between 0 and 1, got %f", p)
	}
	return &binomial{P: p, N: n}, nil
}

func (b binomial) ProbabilityMassFunction(k int) float64 {
	// nCk * p^k * q^(n-k)
	return float64(combin.Binomial(b.N, k)) * math.Pow(b.P, float64(k)) * math.Pow(1-b.P, float64(b.N-k))
}

func (b binomial) CumulativeDensityFunction(k int) float64 {
	s := 0.0

	for i := 0; i <= k; i++ {
		s += b.ProbabilityMassFunction(i)
	}

	return s
}

func (b binomial) Mean() float64 {
	return float64(b.N) * b.P
}

func (b binomial) Median() []float64 {
	if b.P < 0.5 {
		return []float64{math.Floor(float64(b.N) * b.P)}
	}
	if b.P > 0.5 {
		return []float64{math.Ceil(float64(b.N) * b.P)}
	}
	return []float64{float64(b.N) * b.P}
}

func (b binomial) Mode() []float64 {
	l := math.Floor(float64(b.N+1) * b.P)
	u := math.Ceil(float64(b.N+1)*b.P) - 1
	return []float64{l, u}
}

func (b binomial) Variance() float64 {
	q := 1 - b.P
	return float64(b.N) * b.P * q
}

func (b binomial) Skewness() float64 {
	p := b.P
	q := 1 - b.P
	return (q - p) / math.Pow(float64(b.N)*p*q, 0.5)
}

func (b binomial) ExcessKurtosis() float64 {
	p := b.P
	q := 1 - b.P
	return (1 - 6*p*q) / (p * q * float64(b.N))
}

func (b binomial) MomentGeneratingFunction(t float64) float64 {
	p := b.P
	q := 1 - b.P
	return math.Pow(q+p*math.Pow(math.E, t), float64(b.N))
}

func (b binomial) FisherInformation() float64 {
	p := b.P
	q := 1 - b.P
	return float64(b.N) / (p * q)
}
