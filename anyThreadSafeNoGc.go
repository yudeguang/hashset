// Copyright 2020 hashset Author(https://github.com/yudeguang/hashset). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/hashset.
package hashset

import (
	"github.com/yudeguang/noGCMap"
)

// 线程安全
type SetThreadSafeNoGC struct {
	Items *noGCMap.NoGcMapAny //核心部分
}

// New instantiates a new empty set and adds the passed values, if any, to the set
func NewThreadSafeNoGC() *SetThreadSafeNoGC {
	var set SetThreadSafeNoGC
	set.Items = noGCMap.New()
	return &set
}

// Add adds the Items (one or more) to the set.
func (set *SetThreadSafeNoGC) Add(Items ...string) {
	for _, item := range Items {
		set.Items.SetString(item, "")
	}
}

// Remove removes the Items (one or more) from the set.
func (set *SetThreadSafeNoGC) Remove(Items ...string) {
	for _, item := range Items {
		set.Items.DeleteString(item)
	}
}

// Contains check if Items (one or more) are present in the set.
// All Items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *SetThreadSafeNoGC) Contains(Items ...string) bool {
	for _, item := range Items {
		if _, contains := set.Items.GetString(item); !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *SetThreadSafeNoGC) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *SetThreadSafeNoGC) Size() int {
	return set.Items.Len()
}

// Clear clears all values in the set.
func (set *SetThreadSafeNoGC) Clear() {
	set.Items = noGCMap.New()
}
