package rpn

func Add(s *Stack) {
	s.Push(s.Pop() + s.Pop())
}
