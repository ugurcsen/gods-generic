// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashbidimap implements a bidirectional map backed by two hashmaps.
//
// A bidirectional map, or hash bag, is an associative data structure in which the (key,value) pairs form a one-to-one correspondence.
// Thus the binary relation is functional in each direction: value can also act as a key to key.
// A pair (a,b) thus provides a unique coupling between 'a' and 'b' so that 'b' can be found when 'a' is used as a key and 'a' can be found when 'b' is used as a key.
//
// Elements are unordered in the map.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Bidirectional_map
package hashbidimap

import (
	"fmt"
	"github.com/ugurcsen/gods-generic/maps"
	"github.com/ugurcsen/gods-generic/maps/hashmap"
)

// Assert Map implementation
var _ maps.BidiMap[int, int] = (*Map[int, int])(nil)

// Map holds the elements in two hashmaps.
type Map[K, T comparable] struct {
	forwardMap hashmap.Map[K, T]
	inverseMap hashmap.Map[T, K]
}

// New instantiates a bidirectional map.
func New[K, T comparable]() *Map[K, T] {
	return &Map[K, T]{*hashmap.New[K, T](), *hashmap.New[T, K]()}
}

// Put inserts element into the map.
func (m *Map[K, T]) Put(key K, value T) {
	if valueByKey, ok := m.forwardMap.Get(key); ok {
		m.inverseMap.Remove(valueByKey)
	}
	if keyByValue, ok := m.inverseMap.Get(value); ok {
		m.forwardMap.Remove(keyByValue)
	}
	m.forwardMap.Put(key, value)
	m.inverseMap.Put(value, key)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[K, T]) Get(key K) (value T, found bool) {
	return m.forwardMap.Get(key)
}

// GetKey searches the element in the map by value and returns its key or nil if value is not found in map.
// Second return parameter is true if value was found, otherwise false.
func (m *Map[K, T]) GetKey(value T) (key K, found bool) {
	return m.inverseMap.Get(value)
}

// Remove removes the element from the map by key.
func (m *Map[K, T]) Remove(key K) {
	if value, found := m.forwardMap.Get(key); found {
		m.forwardMap.Remove(key)
		m.inverseMap.Remove(value)
	}
}

// Empty returns true if map does not contain any elements
func (m *Map[K, T]) Empty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[K, T]) Size() int {
	return m.forwardMap.Size()
}

// Keys returns all keys (random order).
func (m *Map[K, T]) Keys() []K {
	return m.forwardMap.Keys()
}

// Values returns all values (random order).
func (m *Map[K, T]) Values() []T {
	return m.inverseMap.Keys()
}

// Clear removes all elements from the map.
func (m *Map[K, T]) Clear() {
	m.forwardMap.Clear()
	m.inverseMap.Clear()
}

// String returns a string representation of container
func (m *Map[K, T]) String() string {
	str := "HashBidiMap\n"
	str += fmt.Sprintf("%v", m.forwardMap)
	return str
}
