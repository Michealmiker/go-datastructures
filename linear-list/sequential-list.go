package linear_list

import (
	"fmt"
	"strings"
)

const list_init_size int = 2

// 顺序表
type SequentialList[T comparable] struct {
	// 表的存储空间
	contents []T
}

// 初始化表
func (list *SequentialList[T]) Init() {
	list = &SequentialList[T]{
		contents: make([]T, 0, list_init_size),
	}
}

// 获取表的长度
func (list SequentialList[T]) GetLength() int {
	return len(list.contents)
}

// 获取表的容量
func (list SequentialList[T]) GetCapacity() int {
	return cap(list.contents)
}

// 获取表是否为空
func (list SequentialList[T]) IsEmpty() bool {
	return len(list.contents) == 0
}

// 新增元素
func (list *SequentialList[T]) Add(elem T) {
	list.contents = append(list.contents, elem)
}

func (list *SequentialList[T]) AddFirst(elem T) {
	newList := []T{elem}

	newList = append(newList, list.contents...)

	list.contents = newList
}

func (list *SequentialList[T]) Remove(elem T) {
	index := list.IndexOf(elem)

	if index == -1 {
		return
	}

	list.contents = append(list.contents[:index], list.contents[index+1:]...)
}

func (list *SequentialList[T]) RemoveAt(index int) {
	if index > len(list.contents) {
		return
	}

	list.contents = append(list.contents[:index], list.contents[index+1:]...)
}

func (list SequentialList[T]) IndexOf(elem T) int {
	for i, e := range list.contents {
		if e == elem {
			return i
		}
	}

	return -1
}

func (list SequentialList[T]) LastIndexOf(elem T) int {
	var length = len(list.contents)

	for i := length - 1; i > -1; i-- {
		if list.contents[i] == elem {
			return i
		}
	}

	return -1
}

func (list SequentialList[T]) ToString() string {
	var sb strings.Builder

	for i, elem := range list.contents {
		sb.WriteString(fmt.Sprintf("[%d]: %v\n", i, elem))
	}

	return sb.String()
}
