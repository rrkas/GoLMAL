package tests

import (
	"fmt"

	"github.com/rrkas/GoLMAL/tests/data"
	"gorgonia.org/tensor"
)

func TestMatrices() {
	m1 := data.RandomTensor(0, 10, 3, 3)
	m2 := data.RandomTensor(0, 10, 3, 3)
	r, err := tensor.MatMul(m1, m2)
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(r)
	fmt.Println(err)
}
