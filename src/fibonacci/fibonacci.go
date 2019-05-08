package fibonacci

import (
	"errors"
	"math/big"
)

// Computes the nth Fibonacci number using a bottom-up approach
//  Time complexity = O(n)
//  Storage = O(1)
func ComputeNthFibonacci(n int) (*big.Int, error) {
	if n < 0 {
		return nil, errors.New("cannot compute for negative n")
	}
	if n == 0 {
		return big.NewInt(0), nil
	}
	if n == 1 {
		return big.NewInt(1), nil
	}

	a := big.NewInt(0)
	b := big.NewInt(1)

	for i := 2; i <= n; i++ {
		b.Add(a, b)
		a.Sub(b, a)
	}

	return b, nil
}
