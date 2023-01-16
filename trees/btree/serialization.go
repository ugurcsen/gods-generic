// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package btree

import (
	"encoding/json"
	"github.com/ugurcsen/gods-generic/containers"
	"github.com/ugurcsen/gods-generic/utils"
)

// Assert Serialization implementation
var _ containers.JSONSerializer = (*Tree[int, int])(nil)
var _ containers.JSONDeserializer = (*Tree[int, int])(nil)

// ToJSON outputs the JSON representation of the tree.
func (tree *Tree[K, T]) ToJSON() ([]byte, error) {
	elements := make(map[string]T)
	it := tree.Iterator()
	for it.Next() {
		elements[utils.ToString(it.Key())] = it.Value()
	}
	return json.Marshal(&elements)
}

// FromJSON populates the tree from the input JSON representation.
func (tree *Tree[K, T]) FromJSON(data []byte) error {
	elements := make(map[K]T)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		tree.Clear()
		for key, value := range elements {
			tree.Put(key, value)
		}
	}
	return err
}

// UnmarshalJSON @implements json.Unmarshaler
func (tree *Tree[K, T]) UnmarshalJSON(bytes []byte) error {
	return tree.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (tree *Tree[K, T]) MarshalJSON() ([]byte, error) {
	return tree.ToJSON()
}
