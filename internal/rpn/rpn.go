package rpn

type Stack struct {
	Contents []float64
	Len      int
}

func NewStack() *Stack {
	return &Stack{
		make([]float64, 0, 100),
		-1,
	}
}

func (s *Stack) Pop() float64 {
	if s.Len == -1 {
		panic("empty stack")
	}
	e := s.Contents[s.Len]
	s.Contents = s.Contents[:s.Len]
	s.Len--
	return e
}

func (s *Stack) Push(e float64) {
	s.Len++
	s.Contents = append(s.Contents, e)
}

type Operation func(*Stack)

func (s *Stack) Apply(chain ...Operation) {
	for _, operation := range chain {
		operation(s)
	}
}
