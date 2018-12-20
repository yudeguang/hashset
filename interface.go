//直接引用 github.com/emirpasic/gods/sets/hashset
package hashset

import (
	"fmt"
	"strings"
)

// Set holds elements in go's native map
type Set struct {
	Items map[interface{}]struct{}
}

// New instantiates a new empty set and adds the passed values, if any, to the set
func New(values ...interface{}) *Set {
	set := &Set{Items: make(map[interface{}]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the Items (one or more) to the set.
func (set *Set) Add(Items ...interface{}) {
	for _, item := range Items {
		set.Items[item] = itemExists
	}
}

// Remove removes the Items (one or more) from the set.
func (set *Set) Remove(Items ...interface{}) {
	for _, item := range Items {
		delete(set.Items, item)
	}
}

// Contains check if Items (one or more) are present in the set.
// All Items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set) Contains(Items ...interface{}) bool {
	for _, item := range Items {
		if _, contains := set.Items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set) Size() int {
	return len(set.Items)
}

// Clear clears all values in the set.
func (set *Set) Clear() {
	set.Items = make(map[interface{}]struct{})
}

// Values returns all Items in the set.
func (set *Set) ToSlice() []interface{} {
	values := make([]interface{}, set.Size())
	count := 0
	for item := range set.Items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *Set) String() string {
	str := "HashSet\n"
	Items := []string{}
	for k := range set.Items {
		Items = append(Items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(Items, ", ")
	return str
}

//两个集合的合集
func Union(m, n *Set) *Set {
	resultSet := &Set{Items: make(map[interface{}]struct{}, m.Size()+n.Size())}
	for item := range m.Items {
		resultSet.Items[item] = itemExists
	}
	for item := range n.Items {
		resultSet.Items[item] = itemExists
	}
	return resultSet
}

//两集合的交集
func InnerJoin(m, n *Set) *Set {
	if m.Size() > n.Size() {
		//最大也只能是n
		resultSet := &Set{Items: make(map[interface{}]struct{}, n.Size())}
		for item := range n.Items {
			if _, contains := m.Items[item]; contains {
				resultSet.Items[item] = itemExists
			}
		}
		return resultSet
	} else {
		//最大也只能是m
		resultSet := &Set{Items: make(map[interface{}]struct{}, n.Size())}
		for item := range m.Items {
			if _, contains := n.Items[item]; contains {
				resultSet.Items[item] = itemExists
			}
		}
		return resultSet
	}
}
