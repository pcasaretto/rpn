package rpn_test

import (
	"fmt"
	"testing"
	"testing/quick"

	"github.com/pcasaretto/rpn"
)

func TestSum(t *testing.T) {
	stack := &rpn.Stack{}
	stack.Push(1)
	stack.Push(1)
	fmt.Println(stack)
	rpn.Sum(stack)
	result := stack.Pop()
	if result != 2 {
		t.Error("")
	}
}

func add(x, y float64) float64 {
	stack := &rpn.Stack{}
	stack.Push(x)
	stack.Push(y)
	rpn.Sum(stack)
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

func TestAssociativity(t *testing.T) {
	assertion := func(x, y, z float64) bool {
		return add(add(x, y), z) == add(add(z, y), x)
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error(err)
	}
}
