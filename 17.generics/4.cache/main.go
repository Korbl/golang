package main

import (
	"fmt"
)

// Implement the generic cache interface.
type Cache[T comparable, U any] interface {
	Set(k T, v U)
	Get(k T) U
}

func main() {
	cache := NewCache[int, string]{}
	cache.Set(42, "hello cache")
	fmt.Println(cache.Get(42))
}

type NewCache[T, U comparable] struct {
	cache map[T]U
}

func (n *NewCache[T, U]) Set(k T, v U) {
	n.cache[k] = v
}

func (n *NewCache[T, U]) Get(k T) U {
	return n.cache[k]
}
