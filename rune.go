package hashset

import (
	"fmt"
	"strings"
)

//rune类型,内部导出，方便直接操作
type SetRune struct {
	Items map[rune]struct{}
}

// 初始化Set，支持在初始化的时候插入任意数量的元素
func Newrune(values ...rune) *SetRune {
	set := &SetRune{Items: make(map[rune]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// 向set中添加一个或者多元素
func (set *SetRune) Add(items ...rune) {
	for _, item := range items {
		set.Items[item] = itemExists
	}
}

//删除一个或者多个元素
func (set *SetRune) Remove(items ...rune) {
	for _, item := range items {
		delete(set.Items, item)
	}
}

//判断时候包含传入的所有元素，不传入任何元素也被认为是true
func (set *SetRune) Contains(items ...rune) bool {
	for _, item := range items {
		if _, contains := set.Items[item]; !contains {
			return false
		}
	}
	return true
}

//判断set是否为空
func (set *SetRune) Empty() bool {
	return set.Size() == 0
}

//返回set元素数量
func (set *SetRune) Size() int {
	return len(set.Items)
}

// 清空set
func (set *SetRune) Clear() {
	set.Items = make(map[rune]struct{})
}

// 转换成切片
func (set *SetRune) ToSlice() []rune {
	values := make([]rune, set.Size())
	count := 0
	for item := range set.Items {
		values[count] = item
		count++
	}
	return values
}

// 实现string接口
func (set *SetRune) String() string {
	str := "HashSet\n"
	items := []string{}
	for k := range set.Items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

//两个集合的合集
func Unionrune(m, n *SetRune) *SetRune {
	resultSet := &SetRune{Items: make(map[rune]struct{}, m.Size()+n.Size())}
	for item := range m.Items {
		resultSet.Items[item] = itemExists
	}
	for item := range n.Items {
		resultSet.Items[item] = itemExists
	}
	return resultSet
}

//两集合的交集
func InnerJoinrune(m, n *SetRune) *SetRune {
	if m.Size() > n.Size() {
		//最大也只能是n
		resultSet := &SetRune{Items: make(map[rune]struct{}, n.Size())}
		for item := range n.Items {
			if _, contains := m.Items[item]; contains {
				resultSet.Items[item] = itemExists
			}
		}
		return resultSet
	} else {
		//最大也只能是m
		resultSet := &SetRune{Items: make(map[rune]struct{}, n.Size())}
		for item := range m.Items {
			if _, contains := n.Items[item]; contains {
				resultSet.Items[item] = itemExists
			}
		}
		return resultSet
	}
}
