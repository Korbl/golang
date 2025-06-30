package main

import "fmt"

// Implement a generic stack where we can push and pop elements of type T
type Stack[T any] interface {
	Push(v T)
	Pop() T
}

func main() {
	intStack := &NewStack[int]{}
	intStack.Push(42)
	fmt.Println(intStack.Pop())

	stringStack := &NewStack[string]{}
	stringStack.Push("42")
	fmt.Println(stringStack.Pop())
}

type NewStack[T any] struct {
	items []T
}

func (n *NewStack[T]) Push(v T) {
	n.items = append(n.items, v)
}

func (n *NewStack[T]) Pop() T {
	if len(n.items) == 0 {
		var zero T
		return zero
	}
	item := n.items[len(n.items)-1]
	n.items = n.items[:len(n.items)-1]
	return item
}

func NewStringStack() Stack[string] {
	return nil
}
