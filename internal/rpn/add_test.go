package rpn_test

import (
	"testing"
	"testing/quick"

	"github.com/pcasaretto/rpn"
)

func TestChain(t *testing.T) {
	var tests = []struct {
		prev   []float64
		chain  []rpn.Operation
		result float64
	}{
		{
			[]float64{1, 1},
			[]rpn.Operation{rpn.Add},
			2,
		},
		{
			[]float64{1, 1, 1},
			[]rpn.Operation{rpn.Add, rpn.Add},
			3,
		},
	}

	for _, test := range tests {
		stack := rpn.NewStack()
		for _, e := range test.prev {
			stack.Push(e)
		}
		stack.Apply(test.chain...)
		result := stack.Pop()
		if result != test.result {
			t.Errorf("expected %f, got %f", test.result, result)
		}
	}
}

func add(x, y float64) float64 {
	stack := rpn.NewStack()
	stack.Push(x)
	stack.Push(y)
	rpn.Add(stack)
	return stack.Pop()
}

func TestAddZero(t *testing.T) {
	assertion := func(x float64) bool {
		return add(x, 0) == x
	}

	// Run the assertion through the quick checker
	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}

// func TestAssociativity(t *testing.T) {
// 	assertion := func(x, y, z float64) bool {
// 		return add(add(x, y), z) == add(add(z, y), x)
// 	}

// 	if err := quick.Check(assertion, nil); err != nil {
// 		t.Error(err)
// 	}
// }
