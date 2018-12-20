package hashset

import (
	"fmt"
	"strings"
)

//int64类型,内部导出，方便直接操作
type SetInt64 struct {
	Items map[int64]struct{}
}

// 初始化Set，支持在初始化的时候插入任意数量的元素
func NewInt64(values ...int64) *SetInt64 {
	set := &SetInt64{Items: make(map[int64]struct{}, len(values))}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// 向set中添加一个或者多元素
func (set *SetInt64) Add(items ...int64) {
	for _, item := range items {
		set.Items[item] = itemExists
	}
}

//删除一个或者多个元素
func (set *SetInt64) Remove(items ...int64) {
	for _, item := range items {
		delete(set.Items, item)
	}
}

//判断时候包含传入的所有元素，不传入任何元素也被认为是true
func (set *SetInt64) Contains(items ...int64) bool {
	for _, item := range items {
		if _, contains := set.Items[item]; !contains {
			return false
		}
	}
	return true
}

//判断set是否为空
func (set *SetInt64) Empty() bool {
	return set.Size() == 0
}

//返回set元素数量
func (set *SetInt64) Size() int {
	return len(set.Items)
}

// 清空set
func (set *SetInt64) Clear() {
	set.Items = make(map[int64]struct{})
}

// 转换成切片
func (set *SetInt64) ToSlice() []int64 {
	values := make([]int64, set.Size())
	count := 0
	for item := range set.Items {
		values[count] = item
		count++
	}
	return values
}

// 实现string接口
func (set *SetInt64) String() string {
	str := "HashSet\n"
	items := []string{}
	for k := range set.Items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

//两个集合的合集
func Unionint64(m, n *SetInt64) *SetInt64 {
	resultSet := &SetInt64{Items: make(map[int64]struct{}, m.Size()+n.Size())}
	for item := range m.Items {
		resultSet.Items[item] = itemExists
	}
	for item := range n.Items {
		resultSet.Items[item] = itemExists
	}
	return resultSet
}

//两集合的交集
func InnerJoinint64(m, n *SetInt64) *SetInt64 {
	if m.Size() > n.Size() {
		//最大也只能是n
		resultSet := &SetInt64{Items: make(map[int64]struct{}, n.Size())}
		for item := range n.Items {
			if _, contains := m.Items[item]; contains {
				resultSet.Items[item] = itemExists
			}
		}
		return resultSet
	} else {
		//最大也只能是m
		resultSet := &SetInt64{Items: make(map[int64]struct{}, n.Size())}
		for item := range m.Items {
			if _, contains := n.Items[item]; contains {
				resultSet.Items[item] = itemExists
			}
		}
		return resultSet
	}
}
