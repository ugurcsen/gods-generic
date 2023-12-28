// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treemap implements a map backed by red-black tree.
//
// Elements are ordered by key in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package treemap

import (
	"fmt"
	"github.com/ugurcsen/gods-generic/maps"
	rbt "github.com/ugurcsen/gods-generic/trees/redblacktree"
	"github.com/ugurcsen/gods-generic/utils"
	"strings"
)

// Assert Map implementation
var _ maps.Map[int, int] = (*Map[int, int])(nil)

// Map holds the elements in a red-black tree
type Map[K comparable, T any] struct {
	tree *rbt.Tree[K, T]
}

// NewWith instantiates a tree map with the custom comparator.
func NewWith[K comparable, T any](comparator utils.Comparator[K]) *Map[K, T] {
	return &Map[K, T]{tree: rbt.NewWith[K, T](comparator)}
}

// NewWithNumberComparator instantiates a tree map with the IntComparator, i.e. keys are of type int.
func NewWithNumberComparator[T any]() *Map[int, T] {
	return &Map[int, T]{tree: rbt.NewWithNumberComparator[T]()}
}

// NewWithStringComparator instantiates a tree map with the StringComparator, i.e. keys are of type string.
func NewWithStringComparator[T any]() *Map[string, T] {
	return &Map[string, T]{tree: rbt.NewWithStringComparator[T]()}
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, T]) Put(key K, value T) {
	m.tree.Put(key, value)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, T]) Get(key K) (value T, found bool) {
	return m.tree.Get(key)
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, T]) Remove(key K) {
	m.tree.Remove(key)
}

// Empty returns true if map does not contain any elements
func (m *Map[K, T]) Empty() bool {
	return m.tree.Empty()
}

// Size returns number of elements in the map.
func (m *Map[K, T]) Size() int {
	return m.tree.Size()
}

// Keys returns all keys in-order
func (m *Map[K, T]) Keys() []K {
	return m.tree.Keys()
}

// Values returns all values in-order based on the key.
func (m *Map[K, T]) Values() []T {
	return m.tree.Values()
}

// Clear removes all elements from the map.
func (m *Map[K, T]) Clear() {
	m.tree.Clear()
}

// Min returns the minimum key and its value from the tree map.
// Returns nil, nil if map is empty.
func (m *Map[K, T]) Min() (key K, value T) {
	if node := m.tree.Left(); node != nil {
		return node.Key, node.Value
	}
	var emptyK K
	var emptyT T
	return emptyK, emptyT
}

// Max returns the maximum key and its value from the tree map.
// Returns nil, nil if map is empty.
func (m *Map[K, T]) Max() (key K, value T) {
	if node := m.tree.Right(); node != nil {
		return node.Key, node.Value
	}
	var emptyK K
	var emptyT T
	return emptyK, emptyT
}

// Floor finds the floor key-value pair for the input key.
// In case that no floor is found, then both returned values will be nil.
// It's generally enough to check the first value (key) for nil, which determines if floor was found.
//
// Floor key is defined as the largest key that is smaller than or equal to the given key.
// A floor key may not be found, either because the map is empty, or because
// all keys in the map are larger than the given key.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, T]) Floor(key K) (foundKey K, foundValue T) {
	node, found := m.tree.Floor(key)
	if found {
		return node.Key, node.Value
	}
	var emptyK K
	var emptyT T
	return emptyK, emptyT
}

// Ceiling finds the ceiling key-value pair for the input key.
// In case that no ceiling is found, then both returned values will be nil.
// It's generally enough to check the first value (key) for nil, which determines if ceiling was found.
//
// Ceiling key is defined as the smallest key that is larger than or equal to the given key.
// A ceiling key may not be found, either because the map is empty, or because
// all keys in the map are smaller than the given key.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[K, T]) Ceiling(key K) (foundKey K, foundValue T) {
	node, found := m.tree.Ceiling(key)
	if found {
		return node.Key, node.Value
	}
	var emptyK K
	var emptyT T
	return emptyK, emptyT
}

// String returns a string representation of container
func (m *Map[K, T]) String() string {
	str := "TreeMap\nmap["
	it := m.Iterator()
	for it.Next() {
		str += fmt.Sprintf("%v:%v ", it.Key(), it.Value())
	}
	return strings.TrimRight(str, " ") + "]"

}
