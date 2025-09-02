package data

import (
	"math/rand"

	"gorgonia.org/tensor"
)

func RandomTensor(min float64, max float64, dims ...int) tensor.Tensor {
	if min > max {
		panic("min must be <= max")
	}

	if len(dims) == 0 {
		panic("dims must not be empty")
	}

	shape := 1
	for _, e := range dims {
		shape *= e
	}

	values := []float64{}
	for i := 0; i < shape; i++ {
		val := min + (max-min)*rand.Float64()
		values = append(values, val)
	}

	return tensor.New(
		tensor.WithShape(dims...),
		tensor.WithBacking(values),
	)
}
