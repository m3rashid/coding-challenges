package stack

import "fmt"

type Stack[T interface{}] struct {
	items        []T
	stackPointer int
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		items:        []T{},
		stackPointer: -1,
	}
}

func (s *Stack[T]) Push(item T) {
	s.stackPointer++
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.stackPointer < 0 {
		var res T
		return res, fmt.Errorf("stack is empty")
	}

	item := s.items[s.stackPointer]
	s.items = s.items[:s.stackPointer]
	s.stackPointer--
	return item, nil
}

func (s *Stack[T]) Top() T {
	return s.items[s.stackPointer]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.stackPointer < 0
}

func (s *Stack[T]) Size() int {
	return s.stackPointer + 1
}

func (s *Stack[T]) Clear() {
	s.items = []T{}
	s.stackPointer = -1
}
