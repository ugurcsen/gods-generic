// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import "time"

// Comparator will make type assertion (see IntComparator for example),
// which will panic if a or b are not of the asserted type.
//
// Should return a number:
//
//	negative , if a < b
//	zero     , if a == b
//	positive , if a > b

type ComparableNumber interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Comparator[T comparable] func(a, b T) int

// StringComparator provides a fast comparison on strings
func StringComparator(a, b string) int {
	s1 := a
	s2 := b
	min := len(s2)
	if len(s1) < len(s2) {
		min = len(s1)
	}
	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}
	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}

// NumberComparator provides a basic comparison on int
func NumberComparator[T ComparableNumber](a, b T) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// ByteComparator provides a basic comparison on byte
func ByteComparator(a, b byte) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// RuneComparator provides a basic comparison on rune
func RuneComparator(a, b rune) int {
	aAsserted := a
	bAsserted := b
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// TimeComparator provides a basic comparison on time.Time
func TimeComparator(a, b time.Time) int {
	aAsserted := a
	bAsserted := b

	switch {
	case aAsserted.After(bAsserted):
		return 1
	case aAsserted.Before(bAsserted):
		return -1
	default:
		return 0
	}
}
