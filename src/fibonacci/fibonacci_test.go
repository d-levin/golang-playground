package fibonacci

import (
	"testhelpers"
	"testing"
)

func Test_ComputeNthFibonacci_GivenNonNegativeN_ShouldComputeNthFibonacci(t *testing.T) {
	pairs := []testhelpers.Pair{
		{0, "0"},
		{1, "1"},
		{2, "1"},
		{3, "2"},
		{4, "3"},
		{5, "5"},
		{6, "8"},
		{7, "13"},
		{8, "21"},
		{9, "34"},
		{10, "55"},
		{100, "354224848179261915075"},
		{200, "280571172992510140037611932413038677189525"},
		{300, "222232244629420445529739893461909967206666939096499764990979600"},
	}

	for _, p := range pairs {
		result, _ := ComputeNthFibonacci(p.N)
		if result.String() != p.Expected {
			t.Error("Expected "+p.Expected+", got", result.String())
		}
	}
}
