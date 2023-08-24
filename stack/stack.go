package stack

import "fmt"

var ErrNoSuchElement = fmt.Errorf("no such element")

type Stack struct {
	arr []int
}

func New() *Stack {
	return &Stack{
		arr: make([]int, 0),
	}
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(s.arr)
}

func (s *Stack) Push(element int) error {
	s.arr = append(s.arr, element)

	return nil
}

func (s *Stack) Pop() (*int, error) {
	if s.IsEmpty() {
		return nil, ErrNoSuchElement
	}

	value := s.getLastElement()
	s.arr = s.arr[:s.Size()-1]

	return &value, nil
}

func (s *Stack) Peek() (*int, error) {
	if s.IsEmpty() {
		return nil, ErrNoSuchElement
	}

	value := s.getLastElement()

	return &value, nil
}

func (s *Stack) getLastElement() int {
	lastIndex := s.Size() - 1
	lastElement := s.arr[lastIndex]

	return lastElement
}
