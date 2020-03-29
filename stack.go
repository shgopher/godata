package godata

type Stack struct {
	leng int
	Val    []interface{}
}

func NewStack() *Stack {
	return &Stack{
		leng: 0,
		Val:    nil,
	}
}
func(s *Stack)Length()int{
	return s.leng
}
func (s *Stack) Push(v interface{}) int {
	if v == nil {
		return -1
	}
	s.Val = append(s.Val, v)
	s.leng++
	return s.Length()
}
func (s *Stack) Pop() interface{} {
	if s.Length() > 0 {
		value := s.Val[s.leng-1]
		s.Val = s.Val[:s.leng-1]
		s.leng--
		return value
	}

	return nil
}
func (s *Stack) Top() interface{} {
	if s.Length() > 0 {
		return s.Val[s.leng-1]
	}

	return nil
}
