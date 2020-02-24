// Copyright 2020 hashset Author(https://github.com/yudeguang/hashset). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/hashset.
package hashset

import (
	"fmt"
	"strings"
)

// Set holds elements in go's native map
type SetAny struct {
	Items map[interface{}]struct{}
}

// New instantiates a new empty set and adds the passed values, if any, to the set
func New(values ...interface{}) *SetAny {
	set := &SetAny{Items: make(map[interface{}]struct{}, len(values))}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the Items (one or more) to the set.
func (set *SetAny) Add(Items ...interface{}) {
	for _, item := range Items {
		set.Items[item] = itemExists
	}
}

// Remove removes the Items (one or more) from the set.
func (set *SetAny) Remove(Items ...interface{}) {
	for _, item := range Items {
		delete(set.Items, item)
	}
}

// Contains check if Items (one or more) are present in the set.
// All Items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *SetAny) Contains(Items ...interface{}) bool {
	for _, item := range Items {
		if _, contains := set.Items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *SetAny) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *SetAny) Size() int {
	return len(set.Items)
}

// Clear clears all values in the set.
func (set *SetAny) Clear() {
	set.Items = make(map[interface{}]struct{})
}

// Values returns all Items in the set.
func (set *SetAny) ToSlice() []interface{} {
	values := make([]interface{}, set.Size())
	count := 0
	for item := range set.Items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *SetAny) String() string {
	str := "HashSet\n"
	Items := []string{}
	for k := range set.Items {
		Items = append(Items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(Items, ", ")
	return str
}

//两个集合的合集
func UnionAny(m, n *SetAny) *SetAny {
	resultSet := &SetAny{Items: make(map[interface{}]struct{}, m.Size()+n.Size())}
	for item := range m.Items {
		resultSet.Items[item] = itemExists
	}
	for item := range n.Items {
		resultSet.Items[item] = itemExists
	}
	return resultSet
}

//两集合的交集
func InnerJoinAny(m, n *SetAny) *SetAny {
	if m.Size() > n.Size() {
		//最大也只能是n
		resultSet := &SetAny{Items: make(map[interface{}]struct{}, n.Size())}
		for item := range n.Items {
			if _, contains := m.Items[item]; contains {
				resultSet.Items[item] = itemExists
			}
		}
		return resultSet
	} else {
		//最大也只能是m
		resultSet := &SetAny{Items: make(map[interface{}]struct{}, n.Size())}
		for item := range m.Items {
			if _, contains := n.Items[item]; contains {
				resultSet.Items[item] = itemExists
			}
		}
		return resultSet
	}
}
