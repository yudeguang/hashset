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

//uint类型
type SetUint struct {
	Items map[uint]struct{}
}

// 初始化Set，支持在初始化的时候插入任意数量的元素
func NewUint(values ...uint) *SetUint {
	set := &SetUint{Items: make(map[uint]struct{}, len(values))}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// 向set中添加一个或者多元素
func (set *SetUint) Add(items ...uint) {
	for _, item := range items {
		set.Items[item] = itemExists
	}
}

//删除一个或者多个元素
func (set *SetUint) Remove(items ...uint) {
	for _, item := range items {
		delete(set.Items, item)
	}
}

//判断时候包含传入的所有元素，不传入任何元素也被认为是true
func (set *SetUint) Contains(items ...uint) bool {
	for _, item := range items {
		if _, contains := set.Items[item]; !contains {
			return false
		}
	}
	return true
}

//判断set是否为空
func (set *SetUint) Empty() bool {
	return set.Size() == 0
}

//返回set元素数量
func (set *SetUint) Size() int {
	return len(set.Items)
}

// 清空set
func (set *SetUint) Clear() {
	set.Items = make(map[uint]struct{})
}

// 转换成切片
func (set *SetUint) ToSlice() []uint {
	values := make([]uint, set.Size())
	count := 0
	for item := range set.Items {
		values[count] = item
		count++
	}
	return values
}

// 实现string接口
func (set *SetUint) String() string {
	str := "HashSet\n"
	items := []string{}
	for k := range set.Items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

//两个集合的合集
func UnionUint(m, n *SetUint) *SetUint {
	resultSet := &SetUint{Items: make(map[uint]struct{}, m.Size()+n.Size())}
	for item := range m.Items {
		resultSet.Items[item] = itemExists
	}
	for item := range n.Items {
		resultSet.Items[item] = itemExists
	}
	return resultSet
}

//两集合的交集
func InnerJoinUint(m, n *SetUint) *SetUint {
	if m.Size() > n.Size() {
		//最大也只能是n
		resultSet := &SetUint{Items: make(map[uint]struct{}, n.Size())}
		for item := range n.Items {
			if _, contains := m.Items[item]; contains {
				resultSet.Items[item] = itemExists
			}
		}
		return resultSet
	} else {
		//最大也只能是m
		resultSet := &SetUint{Items: make(map[uint]struct{}, n.Size())}
		for item := range m.Items {
			if _, contains := n.Items[item]; contains {
				resultSet.Items[item] = itemExists
			}
		}
		return resultSet
	}
}
