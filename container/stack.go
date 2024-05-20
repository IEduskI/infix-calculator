package container

import (
	"errors"
)

var (
	ErrEmptyStack = errors.New("stack is empty")
)

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Peek() (T, error) {
	if s.Length() == 0 {
		var empty T
		return empty, ErrEmptyStack
	}

	return s.items[s.Length()-1], nil
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var empty T
		return empty, ErrEmptyStack
	}

	last := s.items[s.Length()-1]

	s.items = s.items[:s.Length()-1]

	return last, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Length() == 0
}

func (s *Stack[T]) Length() int {
	return len(s.items)
}
