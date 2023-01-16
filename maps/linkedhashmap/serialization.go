// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashmap

import (
	"bytes"
	"encoding/json"
	"github.com/ugurcsen/gods-generic/containers"
	"github.com/ugurcsen/gods-generic/utils"
)

// Assert Serialization implementation
var _ containers.JSONSerializer = (*Map[int, int])(nil)
var _ containers.JSONDeserializer = (*Map[int, int])(nil)

// ToJSON outputs the JSON representation of map.
func (m *Map[K, T]) ToJSON() ([]byte, error) {
	var b []byte
	buf := bytes.NewBuffer(b)

	buf.WriteRune('{')

	it := m.Iterator()
	lastIndex := m.Size() - 1
	index := 0

	for it.Next() {
		km, err := json.Marshal(it.Key())
		if err != nil {
			return nil, err
		}
		buf.Write(km)

		buf.WriteRune(':')

		vm, err := json.Marshal(it.Value())
		if err != nil {
			return nil, err
		}
		buf.Write(vm)

		if index != lastIndex {
			buf.WriteRune(',')
		}

		index++
	}

	buf.WriteRune('}')

	return buf.Bytes(), nil
}

// FromJSON populates map from the input JSON representation.
//func (m *Map[K, T) FromJSON(data []byte) error {
//	elements := make(map[string]interface{})
//	err := json.Unmarshal(data, &elements)
//	if err == nil {
//		m.Clear()
//		for key, value := range elements {
//			m.Put(key, value)
//		}
//	}
//	return err
//}

// FromJSON populates map from the input JSON representation.
func (m *Map[K, T]) FromJSON(data []byte) error {
	elements := make(map[K]T)
	err := json.Unmarshal(data, &elements)
	if err != nil {
		return err
	}

	index := make(map[K]int)
	var keys []K
	for key := range elements {
		keys = append(keys, key)
		esc, _ := json.Marshal(key)
		index[key] = bytes.Index(data, esc)
	}

	byIndex := func(a, b K) int {
		key1 := a
		key2 := b
		index1 := index[key1]
		index2 := index[key2]
		return index1 - index2
	}

	utils.Sort(keys, byIndex)

	m.Clear()

	for _, key := range keys {
		m.Put(key, elements[key])
	}

	return nil
}

// UnmarshalJSON @implements json.Unmarshaler
func (m *Map[K, T]) UnmarshalJSON(bytes []byte) error {
	return m.FromJSON(bytes)
}

// MarshalJSON @implements json.Marshaler
func (m *Map[K, T]) MarshalJSON() ([]byte, error) {
	return m.ToJSON()
}
