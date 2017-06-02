package rpn

type Stack struct {
	a [100]float64
	i int
}

func (s *Stack) Pop() float64 {
	if s.i == 0 {
		panic("empty stack")
	}
	defer func() { s.i-- }()
	return s.a[s.i]
}

func (s *Stack) Push(e float64) {
	if s.i == 100 {
		panic("stack overflow")
	}
	s.i++
	s.a[s.i] = e
}

type Operation func(*Stack)

func (s *Stack) Apply(chain ...Operation) {
	for _, operation := range chain {
		operation(s)
	}
}

func Sum(s *Stack) {
	s.Push(s.Pop() + s.Pop())
}
