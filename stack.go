package godata

type Stack struct {
	leng int
	val  []interface{}
}
//return the cap of the slice.
func NewStack(count int) *Stack {
	return &Stack{
		leng: 0,
		val:  make([]interface{}, 0, count),
	}
}
func (s *Stack) Length() int {
	return s.leng
}
func (s *Stack) Push(v interface{}) int {
	if v == nil {
		return -1
	}
	s.val = append(s.val, v)
	s.leng++
	return s.Length()
}
func (s *Stack) Pop() interface{} {
	if s.Length() > 0 {
		value := s.val[s.leng-1]
		s.val = s.val[:s.leng-1]
		s.leng--
		return value
	}

	return nil
}
func (s *Stack) Top() interface{} {
	if s.Length() > 0 {
		return s.val[s.leng-1]
	}

	return nil
}
