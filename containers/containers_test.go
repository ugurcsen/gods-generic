// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// All data structures must implement the container structure

package containers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/roehrijn/g-gods/utils"
)

// For testing purposes
type ContainerTest[T any] struct {
	values []T
}

func (container ContainerTest[T]) Empty() bool {
	return len(container.values) == 0
}

func (container ContainerTest[T]) Size() int {
	return len(container.values)
}

func (container ContainerTest[T]) Clear() {
	container.values = []T{}
}

func (container ContainerTest[T]) Values() []T {
	return container.values
}

func (container ContainerTest[T]) String() string {
	str := "ContainerTest\n"
	var values []string
	for _, value := range container.values {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

func TestGetSortedValuesInts(t *testing.T) {
	container := ContainerTest[int]{}
	GetSortedValues[int](container, utils.NumberComparator[int])
	container.values = []int{5, 1, 3, 2, 4}
	values := GetSortedValues[int](container, utils.NumberComparator[int])
	for i := 1; i < container.Size(); i++ {
		if values[i-1] > values[i] {
			t.Errorf("Not sorted!")
		}
	}
}

func TestGetSortedValuesStrings(t *testing.T) {
	container := ContainerTest[string]{}
	GetSortedValues[string](container, utils.StringComparator)
	container.values = []string{"g", "a", "d", "e", "f", "c", "b"}
	values := GetSortedValues[string](container, utils.StringComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1] > values[i] {
			t.Errorf("Not sorted!")
		}
	}
}
