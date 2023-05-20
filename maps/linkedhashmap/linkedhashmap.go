// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package linkedhashmap is a map that preserves insertion-order.
//
// It is backed by a hash table to store values and doubly-linked list to store ordering.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package linkedhashmap

import (
	"fmt"
	"github.com/roehrijn/g-gods/lists/doublylinkedlist"
	"github.com/roehrijn/g-gods/maps"
	"strings"
)

// Assert Map implementation
var _ maps.Map[int, int] = (*Map[int, int])(nil)

// Map holds the elements in a regular hash table, and uses doubly-linked list to store key ordering.
type Map[K, T comparable] struct {
	table    map[K]T
	ordering *doublylinkedlist.List[K]
}

// New instantiates a linked-hash-map.
func New[K, T comparable]() *Map[K, T] {
	return &Map[K, T]{
		table:    make(map[K]T),
		ordering: doublylinkedlist.New[K](),
	}
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, T]) Put(key K, value T) {
	if _, contains := m.table[key]; !contains {
		m.ordering.Append(key)
	}
	m.table[key] = value
}

// Get searches the element in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, T]) Get(key K) (value T, found bool) {
	value, found = m.table[key]
	return
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, T]) Remove(key K) {
	if _, contains := m.table[key]; contains {
		delete(m.table, key)
		index := m.ordering.IndexOf(key)
		m.ordering.Remove(index)
	}
}

// Empty returns true if map does not contain any elements
func (m *Map[K, T]) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[K, T]) Size() int {
	return m.ordering.Size()
}

// Keys returns all keys in-order
func (m *Map[K, T]) Keys() []K {
	return m.ordering.Values()
}

// Values returns all values in-order based on the key.
func (m *Map[K, T]) Values() []T {
	values := make([]T, m.Size())
	count := 0
	it := m.Iterator()
	for it.Next() {
		values[count] = it.Value()
		count++
	}
	return values
}

// Clear removes all elements from the map.
func (m *Map[K, T]) Clear() {
	m.table = make(map[K]T)
	m.ordering.Clear()
}

// String returns a string representation of container
func (m *Map[K, T]) String() string {
	str := "LinkedHashMap\nmap["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"

}
