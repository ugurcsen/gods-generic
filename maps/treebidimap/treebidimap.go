// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treebidimap implements a bidirectional map backed by two red-black tree.
//
// This structure guarantees that the map will be in both ascending key and value order.
//
// Other than key and value ordering, the goal with this structure is to avoid duplication of elements, which can be significant if contained elements are large.
//
// A bidirectional map, or hash bag, is an associative data structure in which the (key,value) pairs form a one-to-one correspondence.
// Thus the binary relation is functional in each direction: value can also act as a key to key.
// A pair (a,b) thus provides a unique coupling between 'a' and 'b' so that 'b' can be found when 'a' is used as a key and 'a' can be found when 'b' is used as a key.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Bidirectional_map
package treebidimap

import (
	"fmt"
	"github.com/ugurcsen/gods-generic/maps"
	"github.com/ugurcsen/gods-generic/trees/redblacktree"
	"github.com/ugurcsen/gods-generic/utils"
	"strings"
)

// Assert Map implementation
var _ maps.BidiMap[int, int] = (*Map[int, int])(nil)

// Map holds the elements in two red-black trees.
type Map[K, T comparable] struct {
	forwardMap      redblacktree.Tree[K, *data[K, T]]
	inverseMap      redblacktree.Tree[T, *data[K, T]]
	keyComparator   utils.Comparator[K]
	valueComparator utils.Comparator[T]
}

type data[K, T comparable] struct {
	key   K
	value T
}

// NewWith instantiates a bidirectional map.
func NewWith[K, T comparable](keyComparator utils.Comparator[K], valueComparator utils.Comparator[T]) *Map[K, T] {
	return &Map[K, T]{
		forwardMap:      *redblacktree.NewWith[K, *data[K, T]](keyComparator),
		inverseMap:      *redblacktree.NewWith[T, *data[K, T]](valueComparator),
		keyComparator:   keyComparator,
		valueComparator: valueComparator,
	}
}

// NewWithNumberComparators instantiates a bidirectional map with the IntComparator for key and value, i.e. keys and values are of type int.
func NewWithNumberComparators() *Map[int, int] {
	return NewWith(utils.NumberComparator[int], utils.NumberComparator[int])
}

// NewWithStringComparators instantiates a bidirectional map with the StringComparator for key and value, i.e. keys and values are of type string.
func NewWithStringComparators() *Map[string, string] {
	return NewWith(utils.StringComparator, utils.StringComparator)
}

// Put inserts element into the map.
func (m *Map[K, T]) Put(key K, value T) {
	if d, ok := m.forwardMap.Get(key); ok {
		m.inverseMap.Remove(d.value)
	}
	if d, ok := m.inverseMap.Get(value); ok {
		m.forwardMap.Remove(d.key)
	}
	d := &data[K, T]{key: key, value: value}
	m.forwardMap.Put(key, d)
	m.inverseMap.Put(value, d)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[K, T]) Get(key K) (value T, found bool) {
	if d, ok := m.forwardMap.Get(key); ok {
		return d.value, true
	}
	var empty T
	return empty, false
}

// GetKey searches the element in the map by value and returns its key or nil if value is not found in map.
// Second return parameter is true if value was found, otherwise false.
func (m *Map[K, T]) GetKey(value T) (key K, found bool) {
	if d, ok := m.inverseMap.Get(value); ok {
		return d.key, true
	}
	var empty K
	return empty, false
}

// Remove removes the element from the map by key.
func (m *Map[K, T]) Remove(key K) {
	if d, found := m.forwardMap.Get(key); found {
		m.forwardMap.Remove(key)
		m.inverseMap.Remove(d.value)
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

// Keys returns all keys (ordered).
func (m *Map[K, T]) Keys() []K {
	return m.forwardMap.Keys()
}

// Values returns all values (ordered).
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
	str := "TreeBidiMap\nmap["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"
}
