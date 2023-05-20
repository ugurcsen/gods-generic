// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treebidimap

import (
	"github.com/roehrijn/g-gods/containers"
	rbt "github.com/roehrijn/g-gods/trees/redblacktree"
)

// Assert Iterator implementation
var _ containers.ReverseIteratorWithKey[int, int] = (*Iterator[int, int])(nil)

// Iterator holding the iterator's state
type Iterator[K, T comparable] struct {
	iterator rbt.Iterator[K, *data[K, T]]
}

// Iterator returns a stateful iterator whose elements are key/value pairs.
func (m *Map[K, T]) Iterator() Iterator[K, T] {
	return Iterator[K, T]{iterator: m.forwardMap.Iterator()}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's key and value can be retrieved by Key() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator[K, T]) Next() bool {
	return iterator.iterator.Next()
}

// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
// If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[K, T]) Prev() bool {
	return iterator.iterator.Prev()
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator[K, T]) Value() T {
	return iterator.iterator.Value().value
}

// Key returns the current element's key.
// Does not modify the state of the iterator.
func (iterator *Iterator[K, T]) Key() K {
	return iterator.iterator.Key()
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *Iterator[K, T]) Begin() {
	iterator.iterator.Begin()
}

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (iterator *Iterator[K, T]) End() {
	iterator.iterator.End()
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator
func (iterator *Iterator[K, T]) First() bool {
	return iterator.iterator.First()
}

// Last moves the iterator to the last element and returns true if there was a last element in the container.
// If Last() returns true, then last element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[K, T]) Last() bool {
	return iterator.iterator.Last()
}

// NextTo moves the iterator to the next element from current position that satisfies the condition given by the
// passed function, and returns true if there was a next element in the container.
// If NextTo() returns true, then next element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[K, T]) NextTo(f func(key K, value T) bool) bool {
	for iterator.Next() {
		key, value := iterator.Key(), iterator.Value()
		if f(key, value) {
			return true
		}
	}
	return false
}

// PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the
// passed function, and returns true if there was a next element in the container.
// If PrevTo() returns true, then next element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[K, T]) PrevTo(f func(key K, value T) bool) bool {
	for iterator.Prev() {
		key, value := iterator.Key(), iterator.Value()
		if f(key, value) {
			return true
		}
	}
	return false
}
