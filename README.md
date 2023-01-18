[![GoDoc](https://godoc.org/github.com/ugurcsen/gods-generic?status.svg)](https://godoc.org/github.com/ugurcsen/gods-generic)
[![Go Report Card](https://goreportcard.com/badge/github.com/ugurcsen/gods-generic)](https://goreportcard.com/report/github.com/ugurcsen/gods-generic)
[![Sourcegraph](https://sourcegraph.com/github.com/ugurcsen/gods-generic/-/badge.svg)](https://sourcegraph.com/github.com/ugurcsen/gods-generic?badge)
[![Release](https://img.shields.io/github/release/ugurcsen/gods-generic.svg?style=flat-square)](https://github.com/ugurcsen/gods-generic/releases)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ugurcsen_gods-generic&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=gods)
[![PyPI](https://img.shields.io/badge/License-BSD_2--Clause-green.svg)](https://github.com/ugurcsen/gods-generic/blob/master/LICENSE)

# GoDS-Generic (Go Data Structures With Generics)

Implementation of various data structures and algorithms in Go.

## Installation

### Generic version

```bash
go get github.com/ugurcsen/gods-generic
```

### Non Generic version
[https://github.com/emirpasic/gods](https://github.com/emirpasic/gods)
```bash
go get github.com/emirpasic/gods
```

## Data Structures(Generic)

- [x] [Containers](#containers)
  - [x] [Lists](#lists)
    - [x] [ArrayList](#arraylist)
    - [x] [SinglyLinkedList](#singlylinkedlist)
    - [x] [DoublyLinkedList](#doublylinkedlist)
  - [x] [Sets](#sets)
    - [x] [HashSet](#hashset)
    - [x] [TreeSet](#treeset)
    - [x] [LinkedHashSet](#linkedhashset)
  - [x] [Stacks](#stacks)
    - [x] [LinkedListStack](#linkedliststack)
    - [x] [ArrayStack](#arraystack)
  - [x] [Maps](#maps)
    - [x] [HashMap](#hashmap)
    - [x] [TreeMap](#treemap)
    - [x] [LinkedHashMap](#linkedhashmap)
    - [x] [HashBidiMap](#hashbidimap)
    - [x] [TreeBidiMap](#treebidimap)
  - [x] [Trees](#trees)
    - [x] [RedBlackTree](#redblacktree)
    - [x] [AVLTree](#avltree)
    - [x] [BTree](#btree)
    - [x] [BinaryHeap](#binaryheap)
  - [x] [Queues](#queues)
    - [x] [LinkedListQueue](#linkedlistqueue)
    - [x] [ArrayQueue](#arrayqueue)
    - [x] [CircularBuffer](#circularbuffer)
    - [x] [PriorityQueue](#priorityqueue)
- [x] [Functions](#functions)
    - [x] [Comparator](#comparator)
    - [x] [Iterator](#iterator)
      - [x] [IteratorWithIndex](#iteratorwithindex)
      - [x] [IteratorWithKey](#iteratorwithkey)
      - [x] [ReverseIteratorWithIndex](#reverseiteratorwithindex)
      - [x] [ReverseIteratorWithKey](#reverseiteratorwithkey)
    - [x] [Enumerable](#enumerable)
      - [x] [EnumerableWithIndex](#enumerablewithindex)
      - [x] [EnumerableWithKey](#enumerablewithkey)
    - [x] [Serialization](#serialization)
      - [x] [JSONSerializer](#jsonserializer)
      - [x] [JSONDeserializer](#jsondeserializer)
    - [x] [Sort](#sort)
    - [x] [Container](#container)
- [x] [Appendix](#appendix)


## Containers

All data structures implement the container interface with the following methods:

```go
type Container[T comparable] interface {
    Empty() bool
    Size() int
    Clear()
    Values() []T
    String() string
}
```

Containers are either ordered or unordered. All ordered containers provide [stateful iterators](#iterator) and some of them allow [enumerable functions](#enumerable).

| **Data** | **Structure**                         | **Ordered** | **[Iterator](#iterator)** | **[Enumerable](#enumerable)** | **Referenced by** |
| :--- |:--------------------------------------| :---: | :---: | :---: | :---: |
| [Lists](#lists) |
|   | [ArrayList](#arraylist)               | yes | yes* | yes | index |
|   | [SinglyLinkedList](#singlylinkedlist) | yes | yes | yes | index |
|   | [DoublyLinkedList](#doublylinkedlist) | yes | yes* | yes | index |
| [Sets](#sets) |
|   | [HashSet](#hashset)                   | no | no | no | index |
|   | [TreeSet](#treeset)                   | yes | yes* | yes | index |
|   | [LinkedHashSet](#linkedhashset)       | yes | yes* | yes | index |
| [Stacks](#stacks) |
|   | [LinkedListStack](#linkedliststack)   | yes | yes | no | index |
|   | [ArrayStack](#arraystack)             | yes | yes* | no | index |
| [Maps](#maps) |
|   | [HashMap](#hashmap)                   | no | no | no | key |
|   | [TreeMap](#treemap)                   | yes | yes* | yes | key |
|   | [LinkedHashMap](#linkedhashmap)       | yes | yes* | yes | key |
|   | [HashBidiMap](#hashbidimap)           | no | no | no | key* |
|   | [TreeBidiMap](#treebidimap)           | yes | yes* | yes | key* |
| [Trees](#trees) |
|   | [RedBlackTree](#redblacktree)         | yes | yes* | no | key |
|   | [AVLTree](#avltree)                   | yes | yes* | no | key |
|   | [BTree](#btree)                       | yes | yes* | no | key |
|   | [BinaryHeap](#binaryheap)             | yes | yes* | no | index |
| [Queues](#queues) |
|   | [LinkedListQueue](#linkedlistqueue)   | yes | yes | no | index |
|   | [ArrayQueue](#arrayqueue)             | yes | yes* | no | index |
|   | [CircularBuffer](#circularbuffer)     | yes | yes* | no | index |
|   | [PriorityQueue](#priorityqueue)       | yes | yes* | no | index |
|   |                                       |  | <sub><sup>*reversible</sup></sub> |  | <sub><sup>*bidirectional</sup></sub> |

### Lists

A list is a data structure that stores values and may have repeated values.

Implements [Container](#containers) interface.

```go
type List[T comparable] interface {
    Get(index int) (T, bool)
    Remove(index int)
    Add(values ...T)
    Contains(values ...T) bool
    Sort(comparator utils.Comparator[T])
    Swap(index1, index2 int)
    Insert(index int, values ...T)
    Set(index int, value T)
    
    containers.Container[T]
    // Empty() bool
    // Size() int
    // Clear()
    // Values() []interface{}
    // String() string
}
```

#### ArrayList

A [list](#lists) backed by a dynamic array that grows and shrinks implicitly.

Implements [List](#lists), [ReverseIteratorWithIndex](#reverseiteratorwithindex), [EnumerableWithIndex](#enumerablewithindex), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import (
	"github.com/ugurcsen/gods-generic/lists/arraylist"
	"github.com/ugurcsen/gods-generic/utils"
)

func main() {
    list := arraylist.New[string]()
    list.Add("a")                         // ["a"]
    list.Add("c", "b")                    // ["a","c","b"]
    list.Sort(utils.StringComparator)     // ["a","b","c"]
    _, _ = list.Get(0)                    // "a",true
    _, _ = list.Get(100)                  // nil,false
    _ = list.Contains("a", "b", "c")      // true
    _ = list.Contains("a", "b", "c", "d") // false
    list.Swap(0, 1)                       // ["b","a",c"]
    list.Remove(2)                        // ["b","a"]
    list.Remove(1)                        // ["b"]
    list.Remove(0)                        // []
    list.Remove(0)                        // [] (ignored)
    _ = list.Empty()                      // true
    _ = list.Size()                       // 0
    list.Add("a")                         // ["a"]
    list.Clear()                          // []
	list.Insert(0, "b")                   // ["b"]
	list.Insert(0, "a")                   // ["a","b"]
}
```

#### SinglyLinkedList

A [list](#lists) where each element points to the next element in the list.

Implements [List](#lists), [IteratorWithIndex](#iteratorwithindex), [EnumerableWithIndex](#enumerablewithindex), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import (
	sll "github.com/ugurcsen/gods-generic/lists/singlylinkedlist"
	"github.com/ugurcsen/gods-generic/utils"
)

func main() {
    list := sll.New[string]()
    list.Add("a")                         // ["a"]
    list.Append("b")                      // ["a","b"] (same as Add())
    list.Prepend("c")                     // ["c","a","b"]
    list.Sort(utils.StringComparator)     // ["a","b","c"]
    _, _ = list.Get(0)                    // "a",true
    _, _ = list.Get(100)                  // nil,false
    _ = list.Contains("a", "b", "c")      // true
    _ = list.Contains("a", "b", "c", "d") // false
    list.Remove(2)                        // ["a","b"]
    list.Remove(1)                        // ["a"]
    list.Remove(0)                        // []
    list.Remove(0)                        // [] (ignored)
    _ = list.Empty()                      // true
    _ = list.Size()                       // 0
    list.Add("a")                         // ["a"]
    list.Clear()                          // []
	list.Insert(0, "b")                   // ["b"]
	list.Insert(0, "a")                   // ["a","b"]
}
```

#### DoublyLinkedList

A [list](#lists) where each element points to the next and previous elements in the list.

Implements [List](#lists), [ReverseIteratorWithIndex](#reverseiteratorwithindex), [EnumerableWithIndex](#enumerablewithindex), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import (
	dll "github.com/ugurcsen/gods-generic/lists/doublylinkedlist"
	"github.com/ugurcsen/gods-generic/utils"
)

func main() {
    list := dll.New[string]()
    list.Add("a")                         // ["a"]
    list.Append("b")                      // ["a","b"] (same as Add())
    list.Prepend("c")                     // ["c","a","b"]
    list.Sort(utils.StringComparator)     // ["a","b","c"]
    _, _ = list.Get(0)                    // "a",true
    _, _ = list.Get(100)                  // nil,false
    _ = list.Contains("a", "b", "c")      // true
    _ = list.Contains("a", "b", "c", "d") // false
    list.Remove(2)                        // ["a","b"]
    list.Remove(1)                        // ["a"]
    list.Remove(0)                        // []
    list.Remove(0)                        // [] (ignored)
    _ = list.Empty()                      // true
    _ = list.Size()                       // 0
    list.Add("a")                         // ["a"]
    list.Clear()                          // []
	list.Insert(0, "b")                   // ["b"]
	list.Insert(0, "a")                   // ["a","b"]
}
```

### Sets

A set is a data structure that can store elements and has no repeated values. It is a computer implementation of the mathematical concept of a finite set. Unlike most other collection types, rather than retrieving a specific element from a set, one typically tests an element for membership in a set. This structure is often used to ensure that no duplicates are present in a container.

Set additionally allow set operations such as [intersection](https://en.wikipedia.org/wiki/Intersection_(set_theory)), [union](https://en.wikipedia.org/wiki/Union_(set_theory)), [difference](https://proofwiki.org/wiki/Definition:Set_Difference), etc.

Implements [Container](#containers) interface.

```go
type Set[T comparable] interface {
    Add(elements ...T)
    Remove(elements ...T)
    Contains(elements ...T) bool
    // Intersection(another *Set) *Set
    // Union(another *Set) *Set
    // Difference(another *Set) *Set
    
    containers.Container[T]
    // Empty() bool
    // Size() int
    // Clear()
    // Values() []interface{}
    // String() string
}

```

#### HashSet

A [set](#sets) backed by a hash table (actually a Go's map). It makes no guarantees as to the iteration order of the set.

Implements [Set](#sets), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import "github.com/ugurcsen/gods-generic/sets/hashset"

func main() {
    set := hashset.New[int]() // empty (keys are of type int)
    set.Add(1)                // 1
    set.Add(2, 2, 3, 4, 5)    // 3, 1, 2, 4, 5 (random order, duplicates ignored)
    set.Remove(4)             // 5, 3, 2, 1 (random order)
    set.Remove(2, 3)          // 1, 5 (random order)
    set.Contains(1)           // true
    set.Contains(1, 5)        // true
    set.Contains(1, 6)        // false
    _ = set.Values()          // []int{5,1} (random order)
    set.Clear()               // empty
    set.Empty()               // true
    set.Size()                // 0
}
```

#### TreeSet

A [set](#sets) backed by a [red-black tree](#redblacktree) to keep the elements ordered with respect to the [comparator](#comparator).

Implements [Set](#sets), [ReverseIteratorWithIndex](#reverseiteratorwithindex), [EnumerableWithIndex](#enumerablewithindex), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import "github.com/ugurcsen/gods-generic/sets/treeset"

func main() {
	set := treeset.NewWithNumberComparator() // empty (keys are of type int)
	set.Add(1)                            // 1
	set.Add(2, 2, 3, 4, 5)                // 1, 2, 3, 4, 5 (in order, duplicates ignored)
	set.Remove(4)                         // 1, 2, 3, 5 (in order)
	set.Remove(2, 3)                      // 1, 5 (in order)
	set.Contains(1)                       // true
	set.Contains(1, 5)                    // true
	set.Contains(1, 6)                    // false
	_ = set.Values()                      // []int{1,5} (in order)
	set.Clear()                           // empty
	set.Empty()                           // true
	set.Size()                            // 0
}
```

#### LinkedHashSet

A [set](#sets) that preserves insertion-order. Data structure is backed by a hash table to store values and [doubly-linked list](#doublylinkedlist) to store insertion ordering.

Implements [Set](#sets), [ReverseIteratorWithIndex](#reverseiteratorwithindex), [EnumerableWithIndex](#enumerablewithindex), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import "github.com/ugurcsen/gods-generic/sets/linkedhashset"

func main() {
	set := linkedhashset.New[int]() // empty
	set.Add(5)                 // 5
	set.Add(4, 4, 3, 2, 1)     // 5, 4, 3, 2, 1 (in insertion-order, duplicates ignored)
	set.Add(4)                 // 5, 4, 3, 2, 1 (duplicates ignored, insertion-order unchanged)
	set.Remove(4)              // 5, 3, 2, 1 (in insertion-order)
	set.Remove(2, 3)           // 5, 1 (in insertion-order)
	set.Contains(1)            // true
	set.Contains(1, 5)         // true
	set.Contains(1, 6)         // false
	_ = set.Values()           // []int{5, 1} (in insertion-order)
	set.Clear()                // empty
	set.Empty()                // true
	set.Size()                 // 0
}
```

### Stacks

A stack that represents a last-in-first-out (LIFO) data structure. The usual push and pop operations are provided, as well as a method to peek at the top item on the stack.

Implements [Container](#containers) interface.

```go
type Stack[T comparable] interface {
    Push(value T)
    Pop() (value T, ok bool)
    Peek() (value T, ok bool)
    
    containers.Container[T]
    // Empty() bool
    // Size() int
    // Clear()
    // Values() []interface{}
    // String() string
}
```

#### LinkedListStack

A [stack](#stacks) based on a [linked list](#singlylinkedlist).

Implements [Stack](#stacks), [IteratorWithIndex](#iteratorwithindex), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import lls "github.com/ugurcsen/gods-generic/stacks/linkedliststack"

func main() {
	stack := lls.New[int]()  // empty
	stack.Push(1)       // 1
	stack.Push(2)       // 1, 2
	stack.Values()      // 2, 1 (LIFO order)
	_, _ = stack.Peek() // 2,true
	_, _ = stack.Pop()  // 2, true
	_, _ = stack.Pop()  // 1, true
	_, _ = stack.Pop()  // nil, false (nothing to pop)
	stack.Push(1)       // 1
	stack.Clear()       // empty
	stack.Empty()       // true
	stack.Size()        // 0
}
```

#### ArrayStack

A [stack](#stacks) based on a [array list](#arraylist).

Implements [Stack](#stacks), [IteratorWithIndex](#iteratorwithindex), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import "github.com/ugurcsen/gods-generic/stacks/arraystack"

func main() {
	stack := arraystack.New[int]() // empty
	stack.Push(1)             // 1
	stack.Push(2)             // 1, 2
	stack.Values()            // 2, 1 (LIFO order)
	_, _ = stack.Peek()       // 2,true
	_, _ = stack.Pop()        // 2, true
	_, _ = stack.Pop()        // 1, true
	_, _ = stack.Pop()        // nil, false (nothing to pop)
	stack.Push(1)             // 1
	stack.Clear()             // empty
	stack.Empty()             // true
	stack.Size()              // 0
}
```

### Maps

A Map is a data structure that maps keys to values. A map cannot contain duplicate keys and each key can map to at most one value.

Implements [Container](#containers) interface.

```go
type Map[K, T comparable] interface {
    Put(key K, value T)
    Get(key K) (value T, found bool)
    Remove(key K)
    Keys() []K
    
    containers.Container[T]
    // Empty() bool
    // Size() int
    // Clear()
    // Values() []interface{}
    // String() string
}
```

A BidiMap is an extension to the Map. A bidirectional map (BidiMap), also called a hash bag, is an associative data structure in which the key-value pairs form a one-to-one relation. This relation works in both directions by allow the value to also act as a key to key, e.g. a pair (a,b) thus provides a coupling between 'a' and 'b' so that 'b' can be found when 'a' is used as a key and 'a' can be found when 'b' is used as a key.

```go
type BidiMap[K, T comparable] interface {
    GetKey(value T) (key K, found bool)
    
    Map[K, T]
}
```

#### HashMap

A [map](#maps) based on hash tables. Keys are unordered.

Implements [Map](#maps), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import "github.com/ugurcsen/gods-generic/maps/hashmap"

func main() {
	m := hashmap.New[int, string]() // empty
	m.Put(1, "x")                   // 1->x
	m.Put(2, "b")                   // 2->b, 1->x (random order)
	m.Put(1, "a")                   // 2->b, 1->a (random order)
	_, _ = m.Get(2)                 // b, true
	_, _ = m.Get(3)                 // nil, false
	_ = m.Values()                  // []string{"b", "a"} (random order)
	_ = m.Keys()                    // []int{1, 2} (random order)
	m.Remove(1)                     // 2->b
	m.Clear()                       // empty
	m.Empty()                       // true
	m.Size()                        // 0
}
```

#### TreeMap

A [map](#maps) based on [red-black tree](#redblacktree). Keys are ordered with respect to the [comparator](#comparator).

Implements [Map](#maps), [ReverseIteratorWithIndex](#reverseiteratorwithindex), [EnumerableWithKey](#enumerablewithkey), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import "github.com/ugurcsen/gods-generic/maps/treemap"

func main() {
	m := treemap.NewWithNumberComparator[string]() // empty (keys are of type int)
	m.Put(1, "x")                                  // 1->x
	m.Put(2, "b")                                  // 1->x, 2->b (in order)
	m.Put(1, "a")                                  // 1->a, 2->b (in order)
	_, _ = m.Get(2)                                // b, true
	_, _ = m.Get(3)                                // nil, false
	_ = m.Values()                                 // []string{"a", "b"} (in order)
	_ = m.Keys()                                   // []int{1, 2} (in order)
	m.Remove(1)                                    // 2->b
	m.Clear()                                      // empty
	m.Empty()                                      // true
	m.Size()                                       // 0

	// Other:
	m.Min() // Returns the minimum key and its value from map.
	m.Max() // Returns the maximum key and its value from map.
}
```

#### LinkedHashMap

A [map](#maps) that preserves insertion-order. It is backed by a hash table to store values and [doubly-linked list](doublylinkedlist) to store ordering.

Implements [Map](#maps), [ReverseIteratorWithIndex](#reverseiteratorwithindex), [EnumerableWithKey](#enumerablewithkey), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import "github.com/ugurcsen/gods-generic/maps/linkedhashmap"

func main() {
    m := linkedhashmap.New[int, string]() // empty (keys are of type int)
    m.Put(2, "b")                         // 2->b
    m.Put(1, "x")                         // 2->b, 1->x (insertion-order)
    m.Put(1, "a")                         // 2->b, 1->a (insertion-order)
    _, _ = m.Get(2)                       // b, true
    _, _ = m.Get(3)                       // nil, false
    _ = m.Values()                        // []string{"b", "a"} (insertion-order)
    _ = m.Keys()                          // []int{2, 1} (insertion-order)
    m.Remove(1)                           // 2->b
    m.Clear()                             // empty
    m.Empty()                             // true
    m.Size()                              // 0
}

```

#### HashBidiMap

A [map](#maps) based on two hashmaps. Keys are unordered.

Implements [BidiMap](#maps), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import "github.com/ugurcsen/gods-generic/maps/hashbidimap"

func main() {
    m := hashbidimap.New[int, string]() // empty
    m.Put(1, "x")                       // 1->x
    m.Put(3, "b")                       // 1->x, 3->b (random order)
    m.Put(1, "a")                       // 1->a, 3->b (random order)
    m.Put(2, "b")                       // 1->a, 2->b (random order)
    _, _ = m.GetKey("a")                // 1, true
    _, _ = m.Get(2)                     // b, true
    _, _ = m.Get(3)                     // nil, false
    _ = m.Values()                      // []string{"a", "b"} (random order)
    _ = m.Keys()                        // []int{1, 2} (random order)
    m.Remove(1)                         // 2->b
    m.Clear()                           // empty
    m.Empty()                           // true
    m.Size()                            // 0
}
```

#### TreeBidiMap

A [map](#maps) based on red-black tree. This map guarantees that the map will be in both ascending key and value order.  Other than key and value ordering, the goal with this structure is to avoid duplication of elements (unlike in [HashBidiMap](#hashbidimap)), which can be significant if contained elements are large.

Implements [BidiMap](#maps), [ReverseIteratorWithIndex](#reverseiteratorwithindex), [EnumerableWithKey](#enumerablewithkey), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import (
	"github.com/ugurcsen/gods-generic/maps/treebidimap"
	"github.com/ugurcsen/gods-generic/utils"
)

func main() {
    m := treebidimap.NewWith[int, string](utils.NumberComparator[int], utils.StringComparator)
    m.Put(1, "x")        // 1->x
    m.Put(3, "b")        // 1->x, 3->b (ordered)
    m.Put(1, "a")        // 1->a, 3->b (ordered)
    m.Put(2, "b")        // 1->a, 2->b (ordered)
    _, _ = m.GetKey("a") // 1, true
    _, _ = m.Get(2)      // b, true
    _, _ = m.Get(3)      // nil, false
    _ = m.Values()       // []string{"a", "b"} (ordered)
    _ = m.Keys()         // []int{1, 2} (ordered)
    m.Remove(1)          // 2->b
    m.Clear()            // empty
    m.Empty()            // true
    m.Size()             // 0
}
```

### Trees

A tree is a widely used data data structure that simulates a hierarchical tree structure, with a root value and subtrees of children, represented as a set of linked nodes; thus no cyclic links.

Implements [Container](#containers) interface.

```go
type Tree[T comparable] interface {
    containers.Container[T]
    // Empty() bool
    // Size() int
    // Clear()
    // Values() []interface{}
    // String() string
}
```

#### RedBlackTree

A red–black [tree](#trees) is a binary search tree with an extra bit of data per node, its color, which can be either red or black. The extra bit of storage ensures an approximately balanced tree by constraining how nodes are colored from any path from the root to the leaf. Thus, it is a data structure which is a type of self-balancing binary search tree.

The balancing of the tree is not perfect but it is good enough to allow it to guarantee searching in O(log n) time, where n is the total number of elements in the tree. The insertion and deletion operations, along with the tree rearrangement and recoloring, are also performed in O(log n) time. <sub><sup>[Wikipedia](http://en.wikipedia.org/wiki/Red%E2%80%93black_tree)</sup></sub>

Implements [Tree](#trees), [ReverseIteratorWithKey](#reverseiteratorwithkey), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

<p align="center"><img src="http://upload.wikimedia.org/wikipedia/commons/thumb/6/66/Red-black_tree_example.svg/500px-Red-black_tree_example.svg.png" width="400px" height="200px" /></p>

```go
package main

import (
	"fmt"
	rbt "github.com/ugurcsen/gods-generic/trees/redblacktree"
)

func main() {
	tree := rbt.NewWithNumberComparator[string]() // empty (keys are of type int)

	tree.Put(1, "x") // 1->x
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
	tree.Put(3, "c") // 1->a, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	tree.Put(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)

	fmt.Println(tree)
	//
	//  RedBlackTree
	//  │           ┌── 6
	//	│       ┌── 5
	//	│   ┌── 4
	//	│   │   └── 3
	//	└── 2
	//		└── 1

	_ = tree.Values() // []string{"a", "b", "c", "d", "e", "f"} (in order)
	_ = tree.Keys()   // []int{1, 2, 3, 4, 5, 6} (in order)

	tree.Remove(2) // 1->a, 3->c, 4->d, 5->e, 6->f (in order)
	fmt.Println(tree)
	//
	//  RedBlackTree
	//  │       ┌── 6
	//  │   ┌── 5
	//  └── 4
	//      │   ┌── 3
	//      └── 1

	tree.Clear() // empty
	tree.Empty() // true
	tree.Size()  // 0

	// Other:
	tree.Left() // gets the left-most (min) node
	tree.Right() // get the right-most (max) node
	tree.Floor(1) // get the floor node
	tree.Ceiling(1) // get the ceiling node
}
```

Extending the red-black tree's functionality  has been demonstrated in the following [example](https://github.com/ugurcsen/gods-generic/blob/master/examples/redblacktreeextended/redblacktreeextended.go).

#### AVLTree

AVL [tree](#trees) is a self-balancing binary search tree. In an AVL tree, the heights of the two child subtrees of any node differ by at most one; if at any time they differ by more than one, rebalancing is done to restore this property. Lookup, insertion, and deletion all take O(log n) time in both the average and worst cases, where n is the number of nodes in the tree prior to the operation. Insertions and deletions may require the tree to be rebalanced by one or more tree rotations.

AVL trees are often compared with red–black trees because both support the same set of operations and take O(log n) time for the basic operations. For lookup-intensive applications, AVL trees are faster than red–black trees because they are more strictly balanced. <sub><sup>[Wikipedia](https://en.wikipedia.org/wiki/AVL_tree)</sup></sub>

Implements [Tree](#trees), [ReverseIteratorWithKey](#reverseiteratorwithkey), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

<p align="center"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/a/ad/AVL-tree-wBalance_K.svg/262px-AVL-tree-wBalance_K.svg.png" width="300px" height="180px" /><br/><sub>AVL tree with balance factors (green)</sub></p>

```go
package main

import (
	"fmt"
	avl "github.com/ugurcsen/gods-generic/trees/avltree"
)

func main() {
	tree := avl.NewWithNumberComparator[string]() // empty(keys are of type int)

	tree.Put(1, "x") // 1->x
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
	tree.Put(3, "c") // 1->a, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	tree.Put(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)

	fmt.Println(tree)
	//
	//  AVLTree
	//  │       ┌── 6
	//  │   ┌── 5
	//  └── 4
	//      │   ┌── 3
	//      └── 2
	//          └── 1


	_ = tree.Values() // []string{"a", "b", "c", "d", "e", "f"} (in order)
	_ = tree.Keys()   // []int{1, 2, 3, 4, 5, 6} (in order)

	tree.Remove(2) // 1->a, 3->c, 4->d, 5->e, 6->f (in order)
	fmt.Println(tree)
	//
	//  AVLTree
	//  │       ┌── 6
	//  │   ┌── 5
	//  └── 4
	//      └── 3
	//          └── 1

	tree.Clear() // empty
	tree.Empty() // true
	tree.Size()  // 0
}
```

#### BTree

B-tree is a self-balancing tree data structure that keeps data sorted and allows searches, sequential access, insertions, and deletions in logarithmic time. The B-tree is a generalization of a binary search tree in that a node can have more than two children.

According to Knuth's definition, a B-tree of order m is a tree which satisfies the following properties:

- Every node has at most m children.
- Every non-leaf node (except root) has at least ⌈m/2⌉ children.
- The root has at least two children if it is not a leaf node.
- A non-leaf node with k children contains k−1 keys.
- All leaves appear in the same level

Each internal node’s keys act as separation values which divide its subtrees. For example, if an internal node has 3 child nodes (or subtrees) then it must have 2 keys: a1 and a2. All values in the leftmost subtree will be less than a1, all values in the middle subtree will be between a1 and a2, and all values in the rightmost subtree will be greater than a2.<sub><sup>[Wikipedia](http://en.wikipedia.org/wiki/Red%E2%80%93black_tree)</sub></sup>

Implements [Tree](#trees), [ReverseIteratorWithKey](#reverseiteratorwithkey), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

<p align="center"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/6/65/B-tree.svg/831px-B-tree.svg.png" width="400px" height="111px" /></p>

```go
package main

import (
	"fmt"
	"github.com/ugurcsen/gods-generic/trees/btree"
)

func main() {
	tree := btree.NewWithNumberComparator[string](3) // empty (keys are of type int)

	tree.Put(1, "x") // 1->x
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
	tree.Put(3, "c") // 1->a, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	tree.Put(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)
	tree.Put(7, "g") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f, 7->g (in order)

	fmt.Println(tree)
	// BTree
	//         1
	//     2
	//         3
	// 4
	//         5
	//     6
	//         7

	_ = tree.Values() // []string{"a", "b", "c", "d", "e", "f", "g"} (in order)
	_ = tree.Keys()   // []string{1, 2, 3, 4, 5, 6, 7} (in order)

	tree.Remove(2) // 1->a, 3->c, 4->d, 5->e, 6->f, 7->g (in order)
	fmt.Println(tree)
	// BTree
	//     1
	//     3
	// 4
	//     5
	// 6
	//     7

	tree.Clear() // empty
	tree.Empty() // true
	tree.Size()  // 0

	// Other:
	tree.Height() // gets the height of the tree
	tree.Left() // gets the left-most (min) node
	tree.LeftKey() // get the left-most (min) node's key
	tree.LeftValue() // get the left-most (min) node's value
	tree.Right() // get the right-most (max) node
	tree.RightKey() // get the right-most (max) node's key
	tree.RightValue() // get the right-most (max) node's value
}
```

#### BinaryHeap

A binary heap is a [tree](#trees) created using a binary tree. It can be seen as a binary tree with two additional constraints:

- Shape property:

  A binary heap is a complete binary tree; that is, all levels of the tree, except possibly the last one (deepest) are fully filled, and, if the last level of the tree is not complete, the nodes of that level are filled from left to right.
- Heap property:

  All nodes are either greater than or equal to or less than or equal to each of its children, according to a comparison predicate defined for the heap. <sub><sup>[Wikipedia](http://en.wikipedia.org/wiki/Binary_heap)</sub></sup>

Implements [Tree](#trees), [ReverseIteratorWithIndex](#reverseiteratorwithindex), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

<p align="center"><img src="http://upload.wikimedia.org/wikipedia/commons/thumb/3/38/Max-Heap.svg/501px-Max-Heap.svg.png" width="300px" height="200px" /></p>

```go
package main

import (
	"github.com/ugurcsen/gods-generic/trees/binaryheap"
	"github.com/ugurcsen/gods-generic/utils"
)

func main() {

    // Min-heap
    heap := binaryheap.NewWithNumberComparator[int]() // empty (min-heap)
    heap.Push(2)                                      // 2
    heap.Push(3)                                      // 2, 3
    heap.Push(1)                                      // 1, 3, 2
    heap.Values()                                     // 1, 3, 2
    _, _ = heap.Peek()                                // 1,true
    _, _ = heap.Pop()                                 // 1, true
    _, _ = heap.Pop()                                 // 2, true
    _, _ = heap.Pop()                                 // 3, true
    _, _ = heap.Pop()                                 // nil, false (nothing to pop)
    heap.Push(1)                                      // 1
    heap.Clear()                                      // empty
    heap.Empty()                                      // true
    heap.Size()                                       // 0
  
    // Max-heap
    inverseIntComparator := func(a, b int) int {
        return -utils.NumberComparator[int](a, b)
    }
    heap = binaryheap.NewWith(inverseIntComparator) // empty (min-heap)
    heap.Push(2)                                    // 2
    heap.Push(3)                                    // 3, 2
    heap.Push(1)                                    // 3, 2, 1
    heap.Values()                                   // 3, 2, 1
}
```

### Queues

A queue that represents a first-in-first-out (FIFO) data structure. The usual enqueue and dequeue operations are provided, as well as a method to peek at the first item in the queue.

<p align="center"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/5/52/Data_Queue.svg/300px-Data_Queue.svg.png" width="200px" height="120px" /></p>

Implements [Container](#containers) interface.

```go
type Queue[T comparable] interface {
    Enqueue(value T)
    Dequeue() (value T, ok bool)
    Peek() (value T, ok bool)
    
    containers.Container[T]
    // Empty() bool
    // Size() int
    // Clear()
    // Values() []interface{}
    // String() string
}
```

#### LinkedListQueue

A [queue](#queues) based on a [linked list](#singlylinkedlist).

Implements [Queue](#queues), [IteratorWithIndex](#iteratorwithindex), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import llq "github.com/ugurcsen/gods-generic/queues/linkedlistqueue"

// LinkedListQueueExample to demonstrate basic usage of LinkedListQueue
func main() {
    queue := llq.New[int]()     // empty
    queue.Enqueue(1)       // 1
    queue.Enqueue(2)       // 1, 2
    _ = queue.Values()     // 1, 2 (FIFO order)
    _, _ = queue.Peek()    // 1,true
    _, _ = queue.Dequeue() // 1, true
    _, _ = queue.Dequeue() // 2, true
    _, _ = queue.Dequeue() // nil, false (nothing to deque)
    queue.Enqueue(1)       // 1
    queue.Clear()          // empty
    queue.Empty()          // true
    _ = queue.Size()       // 0
}
```

#### ArrayQueue

A [queue](#queues) based on a [array list](#arraylist).

Implements [Queue](#queues), [ReverseIteratorWithIndex](#iteratorwithindex), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import aq "github.com/ugurcsen/gods-generic/queues/arrayqueue"

// ArrayQueueExample to demonstrate basic usage of ArrayQueue
func main() {
    queue := aq.New[int]()      // empty
    queue.Enqueue(1)       // 1
    queue.Enqueue(2)       // 1, 2
    _ = queue.Values()     // 1, 2 (FIFO order)
    _, _ = queue.Peek()    // 1,true
    _, _ = queue.Dequeue() // 1, true
    _, _ = queue.Dequeue() // 2, true
    _, _ = queue.Dequeue() // nil, false (nothing to deque)
    queue.Enqueue(1)       // 1
    queue.Clear()          // empty
    queue.Empty()          // true
    _ = queue.Size()       // 0
}
```

#### CircularBuffer

A circular buffer, circular [queue](#queues), cyclic buffer or ring buffer is a data structure that uses a single, fixed-size buffer as if it were connected end-to-end. This structure lends itself easily to buffering data streams.

<p align="center"><img src="https://upload.wikimedia.org/wikipedia/commons/thumb/f/fd/Circular_Buffer_Animation.gif/400px-Circular_Buffer_Animation.gif" width="300px" height="300px" /></p>

Implements [Queue](#queues), [ReverseIteratorWithIndex](#iteratorwithindex), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import cb "github.com/ugurcsen/gods-generic/queues/circularbuffer"

// CircularBufferExample to demonstrate basic usage of CircularBuffer
func main() {
    queue := cb.New[int](3)     // empty (max size is 3)
    queue.Enqueue(1)       // 1
    queue.Enqueue(2)       // 1, 2
    queue.Enqueue(3)       // 1, 2, 3
    _ = queue.Values()     // 1, 2, 3
    queue.Enqueue(3)       // 4, 2, 3
    _, _ = queue.Peek()    // 4,true
    _, _ = queue.Dequeue() // 4, true
    _, _ = queue.Dequeue() // 2, true
    _, _ = queue.Dequeue() // 3, true
    _, _ = queue.Dequeue() // nil, false (nothing to deque)
    queue.Enqueue(1)       // 1
    queue.Clear()          // empty
    queue.Empty()          // true
    _ = queue.Size()       // 0
}
```

#### PriorityQueue

A priority queue is a special type of [queue](#queues) in which each element is associated with a priority value. And, elements are served on the basis of their priority. That is, higher priority elements are served first. However, if elements with the same priority occur, they are served according to their order in the queue.

Implements [Queue](#queues), [ReverseIteratorWithIndex](#iteratorwithindex), [JSONSerializer](#jsonserializer) and [JSONDeserializer](#jsondeserializer) interfaces.

```go
package main

import (
  pq "github.com/ugurcsen/gods-generic/queues/priorityqueue"
  "github.com/ugurcsen/gods-generic/utils"
)

// Element is an entry in the priority queue
type Element struct {
    name     string
    priority int
}

// Comparator function (sort by element's priority value in descending order)
func byPriority(a, b Element) int {
    priorityA := a.priority
    priorityB := b.priority
    return -utils.NumberComparator[int](priorityA, priorityB) // "-" descending order
}

// PriorityQueueExample to demonstrate basic usage of BinaryHeap
func main() {
    a := Element{name: "a", priority: 1}
    b := Element{name: "b", priority: 2}
    c := Element{name: "c", priority: 3}
  
    queue := pq.NewWith[Element](byPriority) // empty
    queue.Enqueue(a)                         // {a 1}
    queue.Enqueue(c)                         // {c 3}, {a 1}
    queue.Enqueue(b)                         // {c 3}, {b 2}, {a 1}
    _ = queue.Values()                       // [{c 3} {b 2} {a 1}]
    _, _ = queue.Peek()                      // {c 3} true
    _, _ = queue.Dequeue()                   // {c 3} true
    _, _ = queue.Dequeue()                   // {b 2} true
    _, _ = queue.Dequeue()                   // {a 1} true
    _, _ = queue.Dequeue()                   // <nil> false (nothing to dequeue)
    queue.Clear()                            // empty
    _ = queue.Empty()                        // true
    _ = queue.Size()                         // 0
}
```

## Functions

Various helper functions used throughout the library.

### Comparator

Some data structures (e.g. TreeMap, TreeSet) require a comparator function to automatically keep their elements sorted upon insertion. This comparator is necessary during the initalization.

Comparator is defined as:

Return values (int):

```go
negative , if a < b
zero     , if a == b
positive , if a > b
```

Comparator signature:

```go
type ComparableNumber interface {
    int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Comparator[T comparable] func(a, b T) int
```

All common comparators for builtin types are included in the library:

```go
func StringComparator(a, b string) int

func NumberComparator[T ComparableNumber](a, b T) int  // For int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64

func ByteComparator(a, b byte) int

func RuneComparator(a, b rune) int

func TimeComparator(a, b time.Time) int
```

Writing custom comparators is easy:

```go
package main

import (
	"fmt"
	"github.com/ugurcsen/gods-generic/sets/treeset"
)

type User struct {
	id   int
	name string
}

// Custom comparator (sort by IDs)
func byID(a, b User) int {

	// Type assertion, program will panic if this is not respected
	c1 := a
	c2 := b

	switch {
	case c1.id > c2.id:
		return 1
	case c1.id < c2.id:
		return -1
	default:
		return 0
	}
}

func main() {
	set := treeset.NewWith[User](byID)

	set.Add(User{2, "Second"})
	set.Add(User{3, "Third"})
	set.Add(User{1, "First"})
	set.Add(User{4, "Fourth"})

	fmt.Println(set) // {1 First}, {2 Second}, {3 Third}, {4 Fourth}
}
```

### Iterator

All ordered containers have stateful iterators. Typically an iterator is obtained by _Iterator()_ function of an ordered container. Once obtained, iterator's _Next()_ function moves the iterator to the next element and returns true if there was a next element. If there was an element, then element's can be obtained by iterator's _Value()_ function. Depending on the ordering type, it's position can be obtained by iterator's _Index()_ or _Key()_ functions. Some containers even provide reversible iterators, essentially the same, but provide another extra _Prev()_ function that moves the iterator to the previous element and returns true if there was a previous element.

Note: it is unsafe to remove elements from container while iterating.

#### IteratorWithIndex

An [iterator](#iterator) whose elements are referenced by an index.

Typical usage:
```go
it := list.Iterator()
for it.Next() {
	index, value := it.Index(), it.Value()
	...
}
```

Other usages:
```go
if it.First() {
	firstIndex, firstValue := it.Index(), it.Value()
	...
}
```

```go
for it.Begin(); it.Next(); {
	...
}
```

Seeking to a specific element:

```go
// Seek function, i.e. find element starting with "b"
seek := func(index K, value T) bool {
    return strings.HasSuffix(value, "b")
}

// Seek to the condition and continue traversal from that point (forward).
// assumes it.Begin() was called.
for found := it.NextTo(seek); found; found = it.Next() {
    index, value := it.Index(), it.Value()
    ...
}
```

#### IteratorWithKey

An [iterator](#iterator) whose elements are referenced by a key.

Typical usage:
```go
it := tree.Iterator()
for it.Next() {
	key, value := it.Key(), it.Value()
	...
}
```

Other usages:
```go
if it.First() {
	firstKey, firstValue := it.Key(), it.Value()
	...
}
```

```go
for it.Begin(); it.Next(); {
	...
}
```

Seeking to a specific element from the current iterator position:

```go
// Seek function, i.e. find element starting with "b"
seek := func(key K, value T) bool {
    return strings.HasSuffix(value, "b")
}

// Seek to the condition and continue traversal from that point (forward).
// assumes it.Begin() was called.
for found := it.NextTo(seek); found; found = it.Next() {
    key, value := it.Key(), it.Value()
    ...
}
```

#### ReverseIteratorWithIndex

An [iterator](#iterator) whose elements are referenced by an index. Provides all functions as [IteratorWithIndex](#iteratorwithindex), but can also be used for reverse iteration.

Typical usage of iteration in reverse:
```go
it := list.Iterator()
for it.End(); it.Prev(); {
	index, value := it.Index(), it.Value()
	...
}
```

Other usages:
```go
if it.Last() {
	lastIndex, lastValue := it.Index(), it.Value()
	...
}
```

Seeking to a specific element:

```go
// Seek function, i.e. find element starting with "b"
seek := func(index K, value T) bool {
    return strings.HasSuffix(value, "b")
}

// Seek to the condition and continue traversal from that point (in reverse).
// assumes it.End() was called.
for found := it.PrevTo(seek); found; found = it.Prev() {
    index, value := it.Index(), it.Value()
	...
}
```

#### ReverseIteratorWithKey

An [iterator](#iterator) whose elements are referenced by a key. Provides all functions as [IteratorWithKey](#iteratorwithkey), but can also be used for reverse iteration.

Typical usage of iteration in reverse:
```go
it := tree.Iterator()
for it.End(); it.Prev(); {
	key, value := it.Key(), it.Value()
	...
}
```

Other usages:
```go
if it.Last() {
	lastKey, lastValue := it.Key(), it.Value()
	...
}
```

```go
// Seek function, i.e. find element starting with "b"
seek := func(key K, value T) bool {
    return strings.HasSuffix(value, "b")
}

// Seek to the condition and continue traversal from that point (in reverse).
// assumes it.End() was called.
for found := it.PrevTo(seek); found; found = it.Prev() {
    key, value := it.Key(), it.Value()
	...
}
```

### Enumerable

Enumerable functions for ordered containers that implement [EnumerableWithIndex](#enumerablewithindex) or [EnumerableWithKey](#enumerablewithkey) interfaces.

#### EnumerableWithIndex

[Enumerable](#enumerable) functions for ordered containers whose values can be fetched by an index.

**Each**

Calls the given function once for each element, passing that element's index and value.

```go
Each(func(index K, value T))
```

**Map**

Invokes the given function once for each element and returns a container containing the values returned by the given function.

```go
Map(func(index K, value T) T) Container
```

**Select**

Returns a new container containing all elements for which the given function returns a true value.

```go
Select(func(index K, value T) bool) Container
```

**Any**

Passes each element of the container to the given function and returns true if the function ever returns true for any element.

```go
Any(func(index K, value T) bool) bool
```

**All**

Passes each element of the container to the given function and returns true if the function returns true for all elements.

```go
All(func(index K, value T) bool) bool
```

**Find**

Passes each element of the container to the given function and returns the first (index,value) for which the function is true or -1,nil otherwise if no element matches the criteria.

```go
Find(func(index K, value T) bool) (K, T)}
```

**Example:**

```go
package main

import (
	"fmt"
	"github.com/ugurcsen/gods-generic/sets/treeset"
)

func printSet(txt string, set *treeset.Set[int]) {
    fmt.Print(txt, "[ ")
    set.Each(func(index int, value int) {
        fmt.Print(value, " ")
    })
    fmt.Println("]")
}

// EnumerableWithIndexExample to demonstrate basic usage of EnumerableWithIndex
func main() {
    set := treeset.NewWithNumberComparator()
    set.Add(2, 3, 4, 2, 5, 6, 7, 8)
    printSet("Initial", set) // [ 2 3 4 5 6 7 8 ]
  
    even := set.Select(func(index int, value int) bool {
        return value%2 == 0
    })
    printSet("Even numbers", even) // [ 2 4 6 8 ]
  
    foundIndex, foundValue := set.Find(func(index int, value int) bool {
        return value%2 == 0 && value%3 == 0
    })
    if foundIndex != -1 {
        fmt.Println("Number divisible by 2 and 3 found is", foundValue, "at index", foundIndex) // value: 6, index: 4
    }
  
    square := set.Map(func(index int, value int) int {
        return value * value
    })
    printSet("Numbers squared", square) // [ 4 9 16 25 36 49 64 ]
  
    bigger := set.Any(func(index int, value int) bool {
        return value > 5
    })
    fmt.Println("Set contains a number bigger than 5 is ", bigger) // true
  
    positive := set.All(func(index int, value int) bool {
        return value > 0
    })
    fmt.Println("All numbers are positive is", positive) // true
  
    evenNumbersSquared := set.Select(func(index int, value int) bool {
        return value%2 == 0
    }).Map(func(index int, value int) int {
        return value * value
    })
    printSet("Chaining", evenNumbersSquared) // [ 4 16 36 64 ]
}
```

#### EnumerableWithKey

Enumerable functions for ordered containers whose values whose elements are key/value pairs.

**Each**

Calls the given function once for each element, passing that element's key and value.

```go
Each(func(key K, value T))
```

**Map**

Invokes the given function once for each element and returns a container containing the values returned by the given function as key/value pairs.

```go
Map(func(key K, value T) (K, T)) Container
```

**Select**

Returns a new container containing all elements for which the given function returns a true value.

```go
Select(func(key K, value T) bool) Container
```

**Any**

Passes each element of the container to the given function and returns true if the function ever returns true for any element.

```go
Any(func(key K, value T) bool) bool
```

**All**

Passes each element of the container to the given function and returns true if the function returns true for all elements.

```go
All(func(key K, value T) bool) bool
```

**Find**

Passes each element of the container to the given function and returns the first (key,value) for which the function is true or nil,nil otherwise if no element matches the criteria.

```go
Find(func(key K, value T) bool) (K, T)
```

**Example:**

```go
package main

import (
	"fmt"
	"github.com/ugurcsen/gods-generic/maps/treemap"
)

func printMap(txt string, m *treemap.Map[string, int]) {
    fmt.Print(txt, " { ")
    m.Each(func(key string, value int) {
        fmt.Print(key, ":", value, " ")
    })
    fmt.Println("}")
}

// EunumerableWithKeyExample to demonstrate basic usage of EunumerableWithKey
func main() {
    m := treemap.NewWithStringComparator[int]()
    m.Put("g", 7)
    m.Put("f", 6)
    m.Put("e", 5)
    m.Put("d", 4)
    m.Put("c", 3)
    m.Put("b", 2)
    m.Put("a", 1)
    printMap("Initial", m) // { a:1 b:2 c:3 d:4 e:5 f:6 g:7 }
  
    even := m.Select(func(key string, value int) bool {
        return value%2 == 0
    })
    printMap("Elements with even values", even) // { b:2 d:4 f:6 }
  
    foundKey, foundValue := m.Find(func(key string, value int) bool {
        return value%2 == 0 && value%3 == 0
    })
    var empty string
    if foundKey != empty {
      fmt.Println("Element with value divisible by 2 and 3 found is", foundValue, "with key", foundKey) // value: 6, index: 4
    }
  
    square := m.Map(func(key string, value int) (string, int) {
        return key + key, value * value
    })
    printMap("Elements' values squared and letters duplicated", square) // { aa:1 bb:4 cc:9 dd:16 ee:25 ff:36 gg:49 }
  
    bigger := m.Any(func(key string, value int) bool {
        return value > 5
    })
    fmt.Println("Map contains element whose value is bigger than 5 is", bigger) // true
  
    positive := m.All(func(key string, value int) bool {
        return value > 0
    })
    fmt.Println("All map's elements have positive values is", positive) // true
  
    evenNumbersSquared := m.Select(func(key string, value int) bool {
        return value%2 == 0
    }).Map(func(key string, value int) (string, int) {
        return key, value * value
    })
    printMap("Chaining", evenNumbersSquared) // { b:4 d:16 f:36 }
}
```

### Serialization

All data structures can be serialized (marshalled) and deserialized (unmarshalled). Currently, only JSON support is available.

#### JSONSerializer

Outputs the container into its JSON representation.

Typical usage for key-value structures:

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/ugurcsen/gods-generic/maps/hashmap"
)

func main() {
    m := hashmap.New[string, string]()
    m.Put("a", "1")
    m.Put("b", "2")
    m.Put("c", "3")
  
    // Serialization (marshalling)
    json, err := m.ToJSON()
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println(string(json)) // {"a":"1","b":"2","c":"3"}
  
    // Deserialization (unmarshalling)
    json = []byte(`{"a":"1","b":"2"}`)
    err = m.FromJSON(json)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println(m) // HashMap {"a":"1","b":"2"}
}
```

Typical usage for value-only structures:

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/ugurcsen/gods-generic/lists/arraylist"
)

func main() {
	list := arraylist.New[string]()
	list.Add("a", "b", "c")

	bytes, err := json.Marshal(list) // Same as "list.ToJSON(list)"
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes)) // ["a","b","c"]
}
```

#### JSONDeserializer

Populates the container with elements from the input JSON representation.

Typical usage for key-value structures:

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/ugurcsen/gods-generic/maps/hashmap"
)

func main() {
	hm := hashmap.New[string, string]()

	bytes := []byte(`{"a":"1","b":"2"}`)
	err := json.Unmarshal(bytes, &hm) // Same as "hm.FromJSON(bytes)"
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hm) // HashMap map[b:2 a:1]
}
```

Typical usage for value-only structures:

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/ugurcsen/gods-generic/lists/arraylist"
)

func main() {
	list := arraylist.New[string]()

	bytes := []byte(`["a","b"]`)
	err := json.Unmarshal(bytes, &list) // Same as "list.FromJSON(bytes)"
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list) // ArrayList ["a","b"]
}
```

### Sort

Sort is a general purpose sort function.

Lists have an in-place _Sort()_ function and all containers can return their sorted elements via _containers.GetSortedValues()_ function.

Internally these all use the _utils.Sort()_ method:

```go
package main

import "github.com/ugurcsen/gods-generic/utils"

func main() {
	strings := []string{}                  // []
	strings = append(strings, "d")              // ["d"]
	strings = append(strings, "a")              // ["d","a"]
	strings = append(strings, "b")              // ["d","a",b"
	strings = append(strings, "c")              // ["d","a",b","c"]
	utils.Sort(strings, utils.StringComparator) // ["a","b","c","d"]
}
```

### Container

Container specific operations:

```go
// Returns sorted container''s elements with respect to the passed comparator.
// Does not affect the ordering of elements within the container.
func GetSortedValues[T comparable](container Container[T], comparator utils.Comparator[T]) []T
```

Usage:

```go
package main

import (
	"github.com/ugurcsen/gods-generic/lists/arraylist"
	"github.com/ugurcsen/gods-generic/utils"
)

func main() {
	list := arraylist.New[int]()
	list.Add(2, 1, 3)
	values := GetSortedValues(list, utils.NumberComparator[int]) // [1, 2, 3]
}
```

## Appendix

### Motivation

Collections and data structures found in other languages: Java Collections, C++ Standard Template Library (STL) containers, Qt Containers, Ruby Enumerable etc.

### Goals

**Fast algorithms**:

  - Based on decades of knowledge and experiences of other libraries mentioned above.

**Memory efficient algorithms**:

  - Avoiding to consume memory by using optimal algorithms and data structures for the given set of problems, e.g. red-black tree in case of TreeMap to avoid keeping redundant sorted array of keys in memory.

**Easy to use library**:

  - Well-structured library with minimalistic set of atomic operations from which more complex operations can be crafted.

**Stable library**:

  - Only additions are permitted keeping the library backward compatible.

**Solid documentation and examples**:

  - Learning by example.

**Production ready**:

  - Not used in production yet.

**No dependencies**:

  - No external imports.

There is often a tug of war between speed and memory when crafting algorithms. We choose to optimize for speed in most cases within reasonable limits on memory consumption.

Thread safety is not a concern of this project, this should be handled at a higher level.

### Testing and Benchmarking

This takes a while, so test within sub-packages:

`go test -run=NO_TEST -bench . -benchmem  -benchtime 1s ./...`

Non Generic version is [https://github.com/emirpasic/gods](https://github.com/emirpasic/gods)

| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
|--------------------------------|-----------------------|-----------------------|----------|----------------------|----------------------|----------|---------------------|--------------------|----------|
| ArrayList Get 100              | 96 ns/op              | 56,78 ns/op           | %69      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayList Get 1000             | 954,7 ns/op           | 522,8 ns/op           | %83      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayList Get 10000            | 9.563 ns/op           | 5.184 ns/op           | %84      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayList Get 100000           | 95.477 ns/op          | 50.836 ns/op          | %88      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayList Add 100              | 2.977 ns/op           | 567,9 ns/op           | %424     | 6.276 B/op           | 3.125 B/op           | %101     | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayList Add 1000             | 24.379 ns/op          | 6.798 ns/op           | %259     | 45.458 B/op          | 23.602 B/op          | %93      | 744 allocs/op       | 0 allocs/op        | ∞        |
| ArrayList Add 10000            | 280.743 ns/op         | 51.363 ns/op          | %447     | 503.260 B/op         | 242.948 B/op         | %107     | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| ArrayList Add 100000           | 3.542.995 ns/op       | 637.222 ns/op         | %456     | 6.921.637 B/op       | 2.378.902 B/op       | %191     | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| ArrayList Remove 100           | 217,2 ns/op           | 250,7 ns/op           | -%13     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayList Remove 1000          | 2.092 ns/op           | 2.765 ns/op           | -%24     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayList Remove 10000         | 20.951 ns/op          | 25.810 ns/op          | -%19     | 1 B/op               | 0 B/op               | ∞        | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayList Remove 100000        | 4.765.332 ns/op       | 450.777 ns/op         | %957     | 3.328 B/op           | 129 B/op             | %2.480   | 0 allocs/op         | 0 allocs/op        | %0       |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| DoublyLinkedList Get 100       | 1.508 ns/op           | 1.090 ns/op           | %38      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| DoublyLinkedList Get 1000      | 210.991 ns/op         | 201.415 ns/op         | %5       | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| DoublyLinkedList Get 10000     | 28.638.542 ns/op      | 23.120.196 ns/op      | %24      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| DoublyLinkedList Get 100000    | 3.238.820.667 ns/op   | 2.441.314.958 ns/op   | %33      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| DoublyLinkedList Add 100       | 3.079 ns/op           | 2.632 ns/op           | %17      | 3.200 B/op           | 2.400 B/op           | %33      | 100 allocs/op       | 100 allocs/op      | %0       |
| DoublyLinkedList Add 1000      | 40.018 ns/op          | 26.092 ns/op          | %53      | 37.952 B/op          | 24.000 B/op          | %58      | 1.744 allocs/op     | 1.000 allocs/op    | %74      |
| DoublyLinkedList Add 10000     | 414.397 ns/op         | 253.128 ns/op         | %64      | 397.952 B/op         | 240.000 B/op         | %66      | 19.744 allocs/op    | 10.000 allocs/op   | %97      |
| DoublyLinkedList Add 100000    | 4.097.912 ns/op       | 2.672.412 ns/op       | %53      | 3.997.952 B/op       | 2.400.000 B/op       | %67      | 199.744 allocs/op   | 100.000 allocs/op  | %100     |
| DoublyLinkedList Remove 100    | 216,4 ns/op           | 215,9 ns/op           | %0       | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| DoublyLinkedList Remove 1000   | 2.090 ns/op           | 2.071 ns/op           | %1       | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| DoublyLinkedList Remove 10000  | 21.138 ns/op          | 20.824 ns/op          | %2       | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| DoublyLinkedList Remove 100000 | 1.479.976.834 ns/op   | 634.645.521 ns/op     | %133     | 96 B/op              | 0 B/op               | ∞        | 1 allocs/op         | 0 allocs/op        | ∞        |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| SinglyLinkedList Get 100       | 2.718 ns/op           | 2.434 ns/op           | %12      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| SinglyLinkedList Get 1000      | 443.806 ns/op         | 432.509 ns/op         | %3       | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| SinglyLinkedList Get 10000     | 70.921.002 ns/op      | 48.080.932 ns/op      | %48      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| SinglyLinkedList Get 100000    | 5.487.166.668 ns/op   | 4.922.933.792 ns/op   | %11      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| SinglyLinkedList Add 100       | 3.423 ns/op           | 3.201 ns/op           | %7       | 2.400 B/op           | 1.600 B/op           | %50      | 100 allocs/op       | 100 allocs/op      | %0       |
| SinglyLinkedList Add 1000      | 41.118 ns/op          | 23.598 ns/op          | %74      | 29.952 B/op          | 16.000 B/op          | %87      | 1.744 allocs/op     | 1.000 allocs/op    | %74      |
| SinglyLinkedList Add 10000     | 392.871 ns/op         | 226.340 ns/op         | %74      | 317.952 B/op         | 160.000 B/op         | %99      | 19.744 allocs/op    | 10.000 allocs/op   | %97      |
| SinglyLinkedList Add 100000    | 4.032.446 ns/op       | 2.336.021 ns/op       | %73      | 3.197.952 B/op       | 1.600.000 B/op       | %100     | 199.744 allocs/op   | 100.000 allocs/op  | %100     |
| SinglyLinkedList Remove 100    | 217,9 ns/op           | 215,7 ns/op           | %1       | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| SinglyLinkedList Remove 1000   | 2.093 ns/op           | 2.106 ns/op           | -%1      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| SinglyLinkedList Remove 10000  | 21.384 ns/op          | 21.316 ns/op          | %0       | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| SinglyLinkedList Remove 100000 | 1.659.551.583 ns/op   | 1.497.199.208 ns/op   | %11      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| HashBidiMap Get 100            | 1.375 ns/op           | 606,8 ns/op           | %127     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashBidiMap Get 1000           | 10.147 ns/op          | 5.569 ns/op           | %82      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashBidiMap Get 10000          | 244.080 ns/op         | 177.146 ns/op         | %38      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashBidiMap Get 100000         | 2.973.754 ns/op       | 2.110.579 ns/op       | %41      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashBidiMap Put 100            | 10.123 ns/op          | 4.829 ns/op           | %110     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashBidiMap Put 1000           | 77.706 ns/op          | 36.435 ns/op          | %113     | 11.904 B/op          | 0 B/op               | ∞        | 1.488 allocs/op     | 0 allocs/op        | ∞        |
| HashBidiMap Put 10000          | 1.518.899 ns/op       | 922.936 ns/op         | %65      | 155.904 B/op         | 0 B/op               | ∞        | 19.488 allocs/op    | 0 allocs/op        | ∞        |
| HashBidiMap Put 100000         | 17.534.874 ns/op      | 10.306.088 ns/op      | %70      | 1.595.905 B/op       | 0 B/op               | ∞        | 199.488 allocs/op   | 0 allocs/op        | ∞        |
| HashBidiMap Remove 100         | 503,6 ns/op           | 230,1 ns/op           | %119     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashBidiMap Remove 1000        | 4.962 ns/op           | 2.238 ns/op           | %122     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashBidiMap Remove 10000       | 49.393 ns/op          | 24.633 ns/op          | %101     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashBidiMap Remove 100000      | 499.347 ns/op         | 224.164 ns/op         | %123     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| HashMap Get 100                | 1.335 ns/op           | 605,8 ns/op           | %120     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashMap Get 1000               | 10.309 ns/op          | 5.415 ns/op           | %90      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashMap Get 10000              | 249.199 ns/op         | 163.026 ns/op         | %53      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashMap Get 100000             | 2.895.422 ns/op       | 2.025.766 ns/op       | %43      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashMap Put 100                | 2.425 ns/op           | 849,9 ns/op           | %185     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashMap Put 1000               | 19.716 ns/op          | 7.192 ns/op           | %174     | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| HashMap Put 10000              | 370.781 ns/op         | 186.410 ns/op         | %99      | 77.952 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| HashMap Put 100000             | 4.207.168 ns/op       | 2.267.878 ns/op       | %86      | 797.955 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| HashMap Remove 100             | 494,2 ns/op           | 219,3 ns/op           | %125     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashMap Remove 1000            | 4.947 ns/op           | 2.100 ns/op           | %136     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashMap Remove 10000           | 47.827 ns/op          | 21.152 ns/op          | %126     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashMap Remove 100000          | 495.454 ns/op         | 209.361 ns/op         | %137     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| TreeBidiMap Get 100            | 1.773 ns/op           | 1.970 ns/op           | -%10     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| TreeBidiMap Get 1000           | 53.093 ns/op          | 48.321 ns/op          | %10      | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| TreeBidiMap Get 10000          | 686.818 ns/op         | 572.178 ns/op         | %20      | 77.954 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| TreeBidiMap Get 100000         | 8.346.804 ns/op       | 6.604.697 ns/op       | %26      | 797.952 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| TreeBidiMap Put 100            | 23.555 ns/op          | 23.457 ns/op          | %0       | 9.600 B/op           | 6.400 B/op           | %50      | 200 allocs/op       | 200 allocs/op      | %0       |
| TreeBidiMap Put 1000           | 339.843 ns/op         | 315.734 ns/op         | %8       | 107.904 B/op         | 64.000 B/op          | %69      | 3.488 allocs/op     | 2.000 allocs/op    | %74      |
| TreeBidiMap Put 10000          | 4.118.347 ns/op       | 3.623.625 ns/op       | %14      | 1.115.920 B/op       | 640.004 B/op         | %74      | 39.488 allocs/op    | 20.000 allocs/op   | %97      |
| TreeBidiMap Put 100000         | 44.971.067 ns/op      | 40.995.770 ns/op      | %10      | 11.195.917 B/op      | 6.400.010 B/op       | %75      | 399.488 allocs/op   | 200.000 allocs/op  | %100     |
| TreeBidiMap Remove 100         | 389,3 ns/op           | 357,2 ns/op           | %9       | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| TreeBidiMap Remove 1000        | 8.585 ns/op           | 3.510 ns/op           | %145     | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| TreeBidiMap Remove 10000       | 99.369 ns/op          | 35.065 ns/op          | %183     | 77.952 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| TreeBidiMap Remove 100000      | 1.023.437 ns/op       | 354.091 ns/op         | %189     | 797.957 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| TreeMap Get 100                | 1.783 ns/op           | 2.024 ns/op           | -%12     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| TreeMap Get 1000               | 51.320 ns/op          | 48.677 ns/op          | %5       | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| TreeMap Get 10000              | 683.409 ns/op         | 575.812 ns/op         | %19      | 77.954 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| TreeMap Get 100000             | 7.958.002 ns/op       | 6.489.285 ns/op       | %23      | 797.955 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| TreeMap Put 100                | 2.182 ns/op           | 2.155 ns/op           | %1       | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| TreeMap Put 1000               | 58.258 ns/op          | 50.631 ns/op          | %15      | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| TreeMap Put 10000              | 753.072 ns/op         | 662.433 ns/op         | %14      | 77.953 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| TreeMap Put 100000             | 9.180.666 ns/op       | 6.852.508 ns/op       | %34      | 797.956 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| TreeMap Remove 100             | 358,9 ns/op           | 325,6 ns/op           | %10      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| TreeMap Remove 1000            | 8.206 ns/op           | 3.191 ns/op           | %157     | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| TreeMap Remove 10000           | 96.557 ns/op          | 31.861 ns/op          | %203     | 77.952 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| TreeMap Remove 100000          | 982.151 ns/op         | 320.021 ns/op         | %207     | 797.957 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| ArrayQueue Dequeue 100         | 218 ns/op             | 214,3 ns/op           | %2       | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayQueue Dequeue 1000        | 2.522 ns/op           | 2.060 ns/op           | %22      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayQueue Dequeue 10000       | 26.175 ns/op          | 20.559 ns/op          | %27      | 1 B/op               | 0 B/op               | ∞        | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayQueue Dequeue 100000      | 1.448.044.834 ns/op   | 742.191 ns/op         | %195.004 | 699.008 B/op         | 250 B/op             | %279.503 | 8 allocs/op         | 0 allocs/op        | ∞        |
| ArrayQueue Enqueue 100         | 2.232 ns/op           | 497,3 ns/op           | %349     | 6.014 B/op           | 2.082 B/op           | %189     | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayQueue Enqueue 1000        | 29.074 ns/op          | 8.659 ns/op           | %236     | 64.213 B/op          | 30.969 B/op          | %107     | 744 allocs/op       | 0 allocs/op        | ∞        |
| ArrayQueue Enqueue 10000       | 256.823 ns/op         | 72.981 ns/op          | %252     | 502.756 B/op         | 259.789 B/op         | %94      | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| ArrayQueue Enqueue 100000      | 3.456.854 ns/op       | 742.063 ns/op         | %366     | 6.718.641 B/op       | 2.532.538 B/op       | %165     | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| BinaryQueue Dequeue 100        | 219,3 ns/op           | 247,8 ns/op           | -%12     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| BinaryQueue Dequeue 1000       | 2.071 ns/op           | 2.594 ns/op           | -%20     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| BinaryQueue Dequeue 10000      | 20.551 ns/op          | 22.606 ns/op          | -%9      | 1 B/op               | 0 B/op               | ∞        | 0 allocs/op         | 0 allocs/op        | %0       |
| BinaryQueue Dequeue 100000     | 205.545 ns/op         | 257.967 ns/op         | -%20     | 126 B/op             | 66 B/op              | %91      | 0 allocs/op         | 0 allocs/op        | %0       |
| BinaryQueue Enqueue 100        | 4.834 ns/op           | 1.549 ns/op           | %212     | 6.405 B/op           | 2.783 B/op           | %130     | 100 allocs/op       | 0 allocs/op        | ∞        |
| BinaryQueue Enqueue 1000       | 53.640 ns/op          | 13.898 ns/op          | %286     | 74.019 B/op          | 24.219 B/op          | %206     | 1.000 allocs/op     | 0 allocs/op        | ∞        |
| BinaryQueue Enqueue 10000      | 468.757 ns/op         | 133.306 ns/op         | %252     | 628.706 B/op         | 230.784 B/op         | %172     | 10.000 allocs/op    | 0 allocs/op        | ∞        |
| BinaryQueue Enqueue 100000     | 5.096.278 ns/op       | 1.323.769 ns/op       | %285     | 7.050.207 B/op       | 2.326.883 B/op       | %203     | 100.000 allocs/op   | 0 allocs/op        | ∞        |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| HashSet Contains 100           | 1.465 ns/op           | 708,4 ns/op           | %107     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashSet Contains 1000          | 10.611 ns/op          | 6.248 ns/op           | %70      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashSet Contains 10000         | 256.814 ns/op         | 179.416 ns/op         | %43      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashSet Contains 100000        | 2.877.410 ns/op       | 2.183.116 ns/op       | %32      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashSet Add 100                | 1.862 ns/op           | 724,7 ns/op           | %157     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashSet Add 1000               | 17.684 ns/op          | 6.263 ns/op           | %182     | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| HashSet Add 10000              | 347.682 ns/op         | 195.879 ns/op         | %77      | 77.952 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| HashSet Add 100000             | 3.987.649 ns/op       | 2.363.547 ns/op       | %69      | 797.962 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| HashSet Remove 100             | 644,2 ns/op           | 295,3 ns/op           | %118     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| HashSet Remove 1000            | 6.391 ns/op           | 2.882 ns/op           | %122     | 46 B/op              | 10 B/op              | %360     | 0 allocs/op         | 0 allocs/op        | %0       |
| HashSet Remove 10000           | 77.642 ns/op          | 30.941 ns/op          | %151     | 56.950 B/op          | 11.612 B/op          | %390     | 0 allocs/op         | 0 allocs/op        | %0       |
| HashSet Remove 100000          | 26.227.605.125 ns/op  | 10.762.679.624 ns/op  | %144     | 80.403.460.904 B/op  | 40.397.527.296 B/op  | %99      | 100.057 allocs/op   | 101.802 allocs/op  | -%2      |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| TreeSet Contains 100           | 1.844 ns/op           | 2.117 ns/op           | -%13     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| TreeSet Contains 1000          | 52.671 ns/op          | 50.012 ns/op          | %5       | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| TreeSet Contains 10000         | 688.225 ns/op         | 571.616 ns/op         | %20      | 77.954 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| TreeSet Contains 100000        | 8.346.695 ns/op       | 6.569.518 ns/op       | %27      | 797.955 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| TreeSet Add 100                | 2.328 ns/op           | 2.265 ns/op           | %3       | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| TreeSet Add 1000               | 58.447 ns/op          | 52.293 ns/op          | %12      | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| TreeSet Add 10000              | 751.150 ns/op         | 606.855 ns/op         | %24      | 77.952 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| TreeSet Add 100000             | 8.659.309 ns/op       | 6.904.362 ns/op       | %25      | 797.955 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| TreeSet Remove 100             | 389,4 ns/op           | 389,5 ns/op           | -%0      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| TreeSet Remove 1000            | 8.524 ns/op           | 3.835 ns/op           | %122     | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| TreeSet Remove 10000           | 100.427 ns/op         | 38.285 ns/op          | %162     | 77.952 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| TreeSet Remove 100000          | 1.013.442 ns/op       | 383.574 ns/op         | %164     | 797.956 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| ArrayStack Pop 100             | 228,3 ns/op           | 231 ns/op             | -%1      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayStack Pop 1000            | 2.125 ns/op           | 2.235 ns/op           | -%5      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayStack Pop 10000           | 20.972 ns/op          | 22.305 ns/op          | -%6      | 1 B/op               | 0 B/op               | ∞        | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayStack Pop 100000          | 219.957 ns/op         | 222.922 ns/op         | -%1      | 132 B/op             | 65 B/op              | %103     | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayStack Push 100            | 2.362 ns/op           | 419,1 ns/op           | %464     | 6.116 B/op           | 1.683 B/op           | %263     | 0 allocs/op         | 0 allocs/op        | %0       |
| ArrayStack Push 1000           | 26.772 ns/op          | 4.571 ns/op           | %486     | 57.667 B/op          | 28.212 B/op          | %104     | 744 allocs/op       | 0 allocs/op        | ∞        |
| ArrayStack Push 10000          | 249.355 ns/op         | 68.846 ns/op          | %262     | 474.070 B/op         | 257.870 B/op         | %84      | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| ArrayStack Push 100000         | 3.144.425 ns/op       | 576.439 ns/op         | %445     | 6.559.482 B/op       | 2.453.667 B/op       | %167     | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| LinkedListStack Pop 100        | 232,3 ns/op           | 262,4 ns/op           | -%11     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| LinkedListStack Pop 1000       | 2.246 ns/op           | 2.555 ns/op           | -%12     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| LinkedListStack Pop 10000      | 22.320 ns/op          | 25.690 ns/op          | -%13     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| LinkedListStack Pop 100000     | 223.584 ns/op         | 254.731 ns/op         | -%12     | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| LinkedListStack Push 100       | 3.479 ns/op           | 3.145 ns/op           | %11      | 2.400 B/op           | 1.600 B/op           | %50      | 100 allocs/op       | 100 allocs/op      | %0       |
| LinkedListStack Push 1000      | 37.550 ns/op          | 34.586 ns/op          | %9       | 29.952 B/op          | 16.000 B/op          | %87      | 1.744 allocs/op     | 1.000 allocs/op    | %74      |
| LinkedListStack Push 10000     | 374.092 ns/op         | 226.045 ns/op         | %65      | 317.952 B/op         | 160.000 B/op         | %99      | 19.744 allocs/op    | 10.000 allocs/op   | %97      |
| LinkedListStack Push 100000    | 3.930.625 ns/op       | 3.932.688 ns/op       | -%0      | 3.197.952 B/op       | 1.600.000 B/op       | %100     | 199.744 allocs/op   | 100.000 allocs/op  | %100     |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| AVLTree Get 100                | 1.794 ns/op           | 1.735 ns/op           | %3       | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| AVLTree Get 1000               | 52.251 ns/op          | 45.476 ns/op          | %15      | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| AVLTree Get 10000              | 682.255 ns/op         | 553.651 ns/op         | %23      | 77.954 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| AVLTree Get 100000             | 7.930.279 ns/op       | 6.286.729 ns/op       | %26      | 797.956 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| AVLTree Put 100                | 2.729 ns/op           | 2.475 ns/op           | %10      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| AVLTree Put 1000               | 69.900 ns/op          | 63.906 ns/op          | %9       | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| AVLTree Put 10000              | 850.449 ns/op         | 724.575 ns/op         | %17      | 77.952 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| AVLTree Put 100000             | 10.163.641 ns/op      | 8.392.132 ns/op       | %21      | 797.956 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| AVLTree Remove 100             | 320,5 ns/op           | 220,4 ns/op           | %45      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| AVLTree Remove 1000            | 7.265 ns/op           | 2.096 ns/op           | %247     | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| AVLTree Remove 10000           | 85.562 ns/op          | 20.735 ns/op          | %313     | 77.952 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| AVLTree Remove 100000          | 871.629 ns/op         | 209.643 ns/op         | %316     | 797.956 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| BinaryHeap Pop 100             | 217,5 ns/op           | 219,8 ns/op           | -%1      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| BinaryHeap Pop 1000            | 2.097 ns/op           | 2.099 ns/op           | -%0      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| BinaryHeap Pop 10000           | 20.908 ns/op          | 20.907 ns/op          | %0       | 1 B/op               | 0 B/op               | ∞        | 0 allocs/op         | 0 allocs/op        | %0       |
| BinaryHeap Pop 100000          | 211.406 ns/op         | 213.072 ns/op         | -%1      | 124 B/op             | 62 B/op              | %100     | 0 allocs/op         | 0 allocs/op        | %0       |
| BinaryHeap Push 100            | 3.129 ns/op           | 1.344 ns/op           | %133     | 4.916 B/op           | 2.443 B/op           | %101     | 0 allocs/op         | 0 allocs/op        | %0       |
| BinaryHeap Push 1000           | 31.756 ns/op          | 12.078 ns/op          | %163     | 38.026 B/op          | 21.049 B/op          | %81      | 744 allocs/op       | 0 allocs/op        | ∞        |
| BinaryHeap Push 10000          | 337.796 ns/op         | 113.516 ns/op         | %198     | 418.223 B/op         | 214.722 B/op         | %95      | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| BinaryHeap Push 100000         | 3.836.439 ns/op       | 1.131.129 ns/op       | %239     | 4.327.811 B/op       | 2.008.788 B/op       | %115     | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| BTree Get 100                  | 2.192 ns/op           | 2.088 ns/op           | %5       | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| BTree Get 1000                 | 67.574 ns/op          | 60.504 ns/op          | %12      | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| BTree Get 10000                | 799.961 ns/op         | 692.648 ns/op         | %15      | 77.954 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| BTree Get 100000               | 8.812.033 ns/op       | 7.600.964 ns/op       | %16      | 797.957 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| BTree Put 100                  | 6.049 ns/op           | 3.893 ns/op           | %55      | 3.200 B/op           | 1.600 B/op           | %100     | 100 allocs/op       | 100 allocs/op      | %0       |
| BTree Put 1000                 | 91.949 ns/op          | 70.728 ns/op          | %30      | 37.952 B/op          | 16.000 B/op          | %137     | 1.744 allocs/op     | 1.000 allocs/op    | %74      |
| BTree Put 10000                | 1.017.101 ns/op       | 788.238 ns/op         | %29      | 397.958 B/op         | 160.001 B/op         | %149     | 19.744 allocs/op    | 10.000 allocs/op   | %97      |
| BTree Put 100000               | 11.689.052 ns/op      | 8.940.382 ns/op       | %31      | 3.997.973 B/op       | 1.600.003 B/op       | %150     | 199.744 allocs/op   | 100.000 allocs/op  | %100     |
| BTree Remove 100               | 389 ns/op             | 231,4 ns/op           | %68      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| BTree Remove 1000              | 8.517 ns/op           | 2.235 ns/op           | %281     | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| BTree Remove 10000             | 99.440 ns/op          | 22.802 ns/op          | %336     | 77.952 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| BTree Remove 100000            | 1.019.614 ns/op       | 224.342 ns/op         | %354     | 797.961 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| RedBlackTree Get 100           | 1.914 ns/op           | 1.906 ns/op           | %0       | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| RedBlackTree Get 1000          | 52.516 ns/op          | 47.911 ns/op          | %10      | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| RedBlackTree Get 10000         | 691.151 ns/op         | 564.442 ns/op         | %22      | 77.954 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| RedBlackTree Get 100000        | 8.007.029 ns/op       | 6.396.226 ns/op       | %25      | 797.956 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| RedBlackTree Put 100           | 2.354 ns/op           | 2.057 ns/op           | %14      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| RedBlackTree Put 1000          | 58.733 ns/op          | 51.150 ns/op          | %15      | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| RedBlackTree Put 10000         | 752.802 ns/op         | 594.066 ns/op         | %27      | 77.953 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| RedBlackTree Put 100000        | 9.052.725 ns/op       | 7.033.469 ns/op       | %29      | 797.954 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| RedBlackTree Remove 100        | 357,7 ns/op           | 265,1 ns/op           | %35      | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| RedBlackTree Remove 1000       | 8.245 ns/op           | 2.586 ns/op           | %219     | 5.952 B/op           | 0 B/op               | ∞        | 744 allocs/op       | 0 allocs/op        | ∞        |
| RedBlackTree Remove 10000      | 97.490 ns/op          | 25.482 ns/op          | %283     | 77.952 B/op          | 0 B/op               | ∞        | 9.744 allocs/op     | 0 allocs/op        | ∞        |
| RedBlackTree Remove 100000     | 990.313 ns/op         | 255.761 ns/op         | %287     | 797.955 B/op         | 0 B/op               | ∞        | 99.744 allocs/op    | 0 allocs/op        | ∞        |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| GoSortRandom                   | 0,015 ns/op           | 0,013 ns/op           | %9       | 0 B/op               | 0 B/op               | %0       | 0 allocs/op         | 0 allocs/op        | %0       |
| **Benchmarks**                 | **Non Generic**       | **Generic**           | **Gain** | **Non Generic**      | **Generic**          | **Gain** | **Non Generic**     | **Generic**        | **Gain** |
| Averages                       | 484.677.830,044 ns/op | 158.477.047,019 ns/op | %206     | 816.695.882,091 B/op | 410.270.167,112 B/op | %99      | 22.310,80 allocs/op | 4.416,76 allocs/op | %405     |


### Contributing

Biggest contribution towards this library is to use it and give us feedback for further improvements and additions.

For direct contributions, _pull request_ into master branch or ask to become a contributor.

Coding style:

```shell
# Install tooling and set path:
go install gotest.tools/gotestsum@latest
go install golang.org/x/lint/golint@latest
go install github.com/kisielk/errcheck@latest
export PATH=$PATH:$GOPATH/bin

# Fix errors and warnings:
go fmt ./... &&
go test -v ./... && 
golint -set_exit_status ./... && 
! go fmt ./... 2>&1 | read &&
go vet -v ./... &&
gocyclo -avg -over 15 ../gods &&
errcheck ./...
```

### License

This library is distributed under the BSD-style license found in the [LICENSE](https://github.com/ugurcsen/gods-generic/blob/master/LICENSE) file.
