package utils

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
	"gorgonia.org/tensor"
)

func InverseTensor(t tensor.Tensor) (tensor.Tensor, error) {
	// Ensure tensor is 2D
	if t.Dims() != 2 {
		return nil, fmt.Errorf("only 2D tensors (matrices) can be inverted")
	}

	r, c := t.Shape()[0], t.Shape()[1]

	// Convert tensor backing to gonum matrix
	data := t.Data().([]float64)
	matA := mat.NewDense(r, c, data)
	var inv mat.Dense

	if r != c {
		// return nil, fmt.Errorf("matrix must be square")
		// perform Moore-Penrose pseudo-inverse

		var ata mat.Dense
		ata.Mul(matA.T(), matA)
		err := inv.Inverse(&ata)
		if err != nil {
			return nil, err
		}
		ata.Mul(&ata, matA.T())

		// Convert back to tensor
		return tensor.New(
			tensor.WithShape(c, r),
			tensor.WithBacking(ata.RawMatrix().Data),
		), nil
	}

	// Compute inverse
	err := inv.Inverse(matA)
	if err != nil {
		return nil, err
	}

	// Convert back to tensor
	return tensor.New(
		tensor.WithShape(r, c),
		tensor.WithBacking(inv.RawMatrix().Data),
	), nil
}

// Rank with custom tolerance
func Rank(m mat.Matrix, tol float64) int {
	var svd mat.SVD
	ok := svd.Factorize(m, mat.SVDThin)
	if !ok {
		panic("SVD factorization failed")
	}

	values := svd.Values(nil)

	rank := 0
	for _, v := range values {
		if v > tol {
			rank++
		}
	}
	return rank
}

// RankDefault chooses tolerance automatically based on matrix size
func RankDefault(m mat.Matrix) int {
	// Heuristic: ε * max(m, n) * largest singular value
	var svd mat.SVD
	ok := svd.Factorize(m, mat.SVDThin)
	if !ok {
		panic("SVD factorization failed")
	}
	values := svd.Values(nil)

	// Machine epsilon for float64
	eps := 2.220446049250313e-16
	maxVal := values[0]
	rows, cols := m.Dims()
	tol := eps * float64(max(rows, cols)) * maxVal

	rank := 0
	for _, v := range values {
		if v > tol {
			rank++
		}
	}
	return rank
}
