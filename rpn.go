package rpn

type Stack struct {
	a []float64
	i int
}

func NewStack() *Stack {
	return &Stack{
		make([]float64, 0, 100),
		-1,
	}
}

func (s *Stack) Pop() float64 {
	if s.i == -1 {
		panic("empty stack")
	}
	e := s.a[s.i]
	s.a = s.a[:s.i]
	s.i--
	return e
}

func (s *Stack) Push(e float64) {
	s.i++
	s.a = append(s.a, e)
}

type Operation func(*Stack)

func (s *Stack) Apply(chain ...Operation) {
	for _, operation := range chain {
		operation(s)
	}
}
