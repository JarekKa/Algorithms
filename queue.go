package algorithms

import (
	"errors"
)

type Queue struct {
	Elements []int
	Max      int
}

func NewQueue(opts ...func(*Queue)) *Queue {
	s := &Queue{}
	for _, o := range opts {
		o(s)
	}
	s.Elements = make([]int, s.Max)
	return s
}

func SetMax(i int) func(*Queue) {
	return func(s *Queue) {
		s.Max = i
	}
}

func (s *Queue) Dequeue() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	s.Elements = s.Elements[:len(s.Elements)-1]
	return s.Elements[len(s.Elements)-1], nil
}
func (s *Queue) Enqueue(i int) error {
	if s.IsFull() {
		return errors.New("queue is full")
	}
	s.Elements = append(s.Elements[:0], append([]int{i}, s.Elements[0:]...)...)
	return nil
}
func (s *Queue) Peep() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	return s.Elements[0], nil
}
func (s *Queue) IsFull() bool {
	return len(s.Elements) == s.Max
}
func (s *Queue) IsEmpty() bool {
	return len(s.Elements) == 0
}
