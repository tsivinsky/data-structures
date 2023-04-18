package main

import "fmt"

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) Pop() {
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("%+v", s.data)
}

func NewStack[T any](data []T) *Stack[T] {
	return &Stack[T]{
		data: data,
	}
}

func main() {
	stack := NewStack([]int{})

	stack.Push(69)
	stack.Push(420)
	stack.Push(1337)

	fmt.Printf("stack.Len(): %v\n", stack.Len())
	fmt.Printf("stack: %v\n", stack)

	stack.Pop()

	fmt.Printf("stack.Len(): %v\n", stack.Len())
	fmt.Printf("stack: %v\n", stack)
}
