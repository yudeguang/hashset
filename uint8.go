package hashset

import (
	"fmt"
	"strings"
)

//uint8类型,内部导出，方便直接操作
type SetUint8 struct {
	Items map[uint8]struct{}
}

// 初始化Set，支持在初始化的时候插入任意数量的元素
func NewUint8(values ...uint8) *SetUint8 {
	set := &SetUint8{Items: make(map[uint8]struct{}, len(values))}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// 向set中添加一个或者多元素
func (set *SetUint8) Add(items ...uint8) {
	for _, item := range items {
		set.Items[item] = itemExists
	}
}

//删除一个或者多个元素
func (set *SetUint8) Remove(items ...uint8) {
	for _, item := range items {
		delete(set.Items, item)
	}
}

//判断时候包含传入的所有元素，不传入任何元素也被认为是true
func (set *SetUint8) Contains(items ...uint8) bool {
	for _, item := range items {
		if _, contains := set.Items[item]; !contains {
			return false
		}
	}
	return true
}

//判断set是否为空
func (set *SetUint8) Empty() bool {
	return set.Size() == 0
}

//返回set元素数量
func (set *SetUint8) Size() int {
	return len(set.Items)
}

// 清空set
func (set *SetUint8) Clear() {
	set.Items = make(map[uint8]struct{})
}

// 转换成切片
func (set *SetUint8) ToSlice() []uint8 {
	values := make([]uint8, set.Size())
	count := 0
	for item := range set.Items {
		values[count] = item
		count++
	}
	return values
}

// 实现string接口
func (set *SetUint8) String() string {
	str := "HashSet\n"
	items := []string{}
	for k := range set.Items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

//两个集合的合集
func UnionUint8(m, n *SetUint8) *SetUint8 {
	resultSet := &SetUint8{Items: make(map[uint8]struct{}, m.Size()+n.Size())}
	for item := range m.Items {
		resultSet.Items[item] = itemExists
	}
	for item := range n.Items {
		resultSet.Items[item] = itemExists
	}
	return resultSet
}

//两集合的交集
func InnerJoinUint8(m, n *SetUint8) *SetUint8 {
	if m.Size() > n.Size() {
		//最大也只能是n
		resultSet := &SetUint8{Items: make(map[uint8]struct{}, n.Size())}
		for item := range n.Items {
			if _, contains := m.Items[item]; contains {
				resultSet.Items[item] = itemExists
			}
		}
		return resultSet
	} else {
		//最大也只能是m
		resultSet := &SetUint8{Items: make(map[uint8]struct{}, n.Size())}
		for item := range m.Items {
			if _, contains := n.Items[item]; contains {
				resultSet.Items[item] = itemExists
			}
		}
		return resultSet
	}
}
