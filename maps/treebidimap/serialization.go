// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treebidimap

import (
	"encoding/json"
	"github.com/ugurcsen/gods-generic/containers"
	"github.com/ugurcsen/gods-generic/utils"
)

// Assert Serialization implementation
var _ containers.JSONSerializer = (*Map[int, int])(nil)
var _ containers.JSONDeserializer = (*Map[int, int])(nil)

// ToJSON outputs the JSON representation of the map.
func (m *Map[K, T]) ToJSON() ([]byte, error) {
	elements := make(map[string]interface{})
	it := m.Iterator()
	for it.Next() {
		elements[utils.ToString(it.Key())] = it.Value()
	}
	return json.Marshal(&elements)
}

// FromJSON populates the map from the input JSON representation.
func (m *Map[K, T]) FromJSON(data []byte) error {
	elements := make(map[K]T)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		m.Clear()
		for key, value := range elements {
			m.Put(key, value)
		}
	}
	return err
}

// UnmarshalJSON @implements json.Unmarshaler
func (m *Map[K, T]) UnmarshalJSON(bytes []byte) error {
	return m.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (m *Map[K, T]) MarshalJSON() ([]byte, error) {
	return m.ToJSON()
}
