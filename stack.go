package algorithms

import (
	"errors"
)

type Stack struct {
	Elements []int
	Max      int
}

func NewStack(opts ...func(*Stack)) *Stack {
	s := &Stack{}
	for _, o := range opts {
		o(s)
	}
	s.Elements = make([]int, s.Max)
	return s
}

func SetMax(i int) func(*Stack) {
	return func(s *Stack) {
		s.Max = i
	}
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	e := s.Elements[0]
	s.Elements = append(s.Elements[:0], s.Elements[i+0:]...)
	return e, nil
}
func (s *Stack) Push(i int) error {
	if s.IsFull() {
		return errors.New("stack is full")
	}
	s.Elements = append(s.Elements, i)
	return nil
}
func (s *Stack) Peep() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.Elements[len(s.Elements)-1], nil
}
func (s *Stack) IsFull() bool {
	return len(s.Elements) == s.Max
}
func (s *Stack) IsEmpty() bool {
	return len(s.Elements) == 0
}
