// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashmap implements a map backed by a hash table.
//
// Elements are unordered in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package hashmap

import (
	"fmt"

	"github.com/ugurcsen/gods-generic/maps"
)

// Assert Map implementation
var _ maps.Map[int, int] = (*Map[int, int])(nil)

// Map holds the elements in go's native map
type Map[K comparable, T any] struct {
	m map[K]T
}

// New instantiates a hash map.
func New[K comparable, T any]() *Map[K, T] {
	return &Map[K, T]{m: make(map[K]T)}
}

// Put inserts element into the map.
func (m *Map[K, T]) Put(key K, value T) {
	m.m[key] = value
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[K, T]) Get(key K) (value T, found bool) {
	value, found = m.m[key]
	return
}

// Remove removes the element from the map by key.
func (m *Map[K, T]) Remove(key K) {
	delete(m.m, key)
}

// Empty returns true if map does not contain any elements
func (m *Map[K, T]) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[K, T]) Size() int {
	return len(m.m)
}

// Keys returns all keys (random order).
func (m *Map[K, T]) Keys() []K {
	keys := make([]K, m.Size())
	count := 0
	for key := range m.m {
		keys[count] = key
		count++
	}
	return keys
}

// Values returns all values (random order).
func (m *Map[K, T]) Values() []T {
	values := make([]T, m.Size())
	count := 0
	for _, value := range m.m {
		values[count] = value
		count++
	}
	return values
}

// Clear removes all elements from the map.
func (m *Map[K, T]) Clear() {
	m.m = make(map[K]T)
}

// String returns a string representation of container
func (m *Map[K, T]) String() string {
	str := "HashMap\n"
	str += fmt.Sprintf("%v", m.m)
	return str
}
