// Copyright 2020 hashset Author(https://github.com/yudeguang/hashset). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/hashset.
package hashset

import (
	"fmt"
	"strings"
	"sync"
)

// 线程安全
type SetThreadSafe struct {
	Items sync.Map
}

// New instantiates a new empty set and adds the passed values, if any, to the set
func NewThreadSafe(values ...interface{}) *SetThreadSafe {
	set := new(SetThreadSafe)
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the Items (one or more) to the set.
func (set *SetThreadSafe) Add(Items ...interface{}) {
	for _, item := range Items {
		set.Items.Store(item, itemExists)
	}
}

// Remove removes the Items (one or more) from the set.
func (set *SetThreadSafe) Remove(Items ...interface{}) {
	for _, item := range Items {
		set.Items.Delete(item)
	}
}

// Contains check if Items (one or more) are present in the set.
// All Items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *SetThreadSafe) Contains(Items ...interface{}) bool {
	for _, item := range Items {
		if _, contains := set.Items.Load(item); !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *SetThreadSafe) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *SetThreadSafe) Size() int {
	i := 0
	set.Items.Range(func(k, v interface{}) bool {
		i++
		return true
	})
	return i
}

// Clear clears all values in the set.
func (set *SetThreadSafe) Clear() {
	set = new(SetThreadSafe)
}

// Values returns all Items in the set.
func (set *SetThreadSafe) ToSlice() []interface{} {
	var ret []interface{}
	set.Items.Range(func(k, v interface{}) bool {
		ret = append(ret, k)
		return true
	})
	return ret
}

// String returns a string representation of container
func (set *SetThreadSafe) String() string {
	str := "HashSet\n"
	Items := []string{}
	set.Items.Range(func(k, v interface{}) bool {
		Items = append(Items, fmt.Sprintf("%v", k))
		return true
	})
	str += strings.Join(Items, ", ")
	return str
}

//两个集合的合集
func UnionThreadSafe(m, n *SetThreadSafe) *SetThreadSafe {
	resultSet := new(SetThreadSafe)
	m.Items.Range(func(k, v interface{}) bool {
		resultSet.Items.Load(k)
		return true
	})
	n.Items.Range(func(k, v interface{}) bool {
		resultSet.Items.Load(k)
		return true
	})
	return resultSet
}

//两集合的交集,这里不能同其它类型一样，先用Size函数得到长度，再计算，性能会较差
func InnerJoinThreadSafe(m, n *SetThreadSafe) *SetThreadSafe {
	resultSet := new(SetThreadSafe)
	m.Items.Range(func(k, v interface{}) bool {
		if _, exist := n.Items.Load(k); exist {
			resultSet.Items.Store(k, itemExists)
		}
		resultSet.Items.Load(k)
		return true
	})
	return resultSet
}
