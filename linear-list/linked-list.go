package linear_list

import (
	"fmt"
	"strings"
)

// 链表节点
type LinkedListNode[T comparable] struct {
	// 数据域
	data T
	// 指针域
	next *LinkedListNode[T]
}

// 链表
type LinkedList[T comparable] struct {
	// 表的长度
	count int
	// 表的指针
	head *LinkedListNode[T]
}

func (list *LinkedList[T]) Init() {
	list.count = 0
	list.head = &LinkedListNode[T]{}
}

func (list LinkedList[T]) GetLength() int {
	return list.count
}

func (list LinkedList[T]) IsEmpty() bool {
	return list.count == 0
}

func (list *LinkedList[T]) Add(elem T) {
	node := LinkedListNode[T]{
		data: elem,
	}
	ptr := list.head

	for ptr.next != nil {
		ptr = ptr.next
	}

	ptr.next = &node

	list.count++
}

func (list *LinkedList[T]) AddFirst(elem T) {
	node := LinkedListNode[T]{
		data: elem,
		next: list.head.next,
	}

	list.head.next = &node

	list.count++
}

func (list *LinkedList[T]) Remove(elem T) {
	if list.head.next.data == elem {
		list.head.next = list.head.next.next

		list.count--

		return
	}

	ptr := list.head.next

	for ptr.next != nil {
		if ptr.next.data == elem {
			ptr.next = ptr.next.next

			list.count--

			break
		}

		ptr = ptr.next
	}
}

func (list *LinkedList[T]) RemoveAt(index int) {
	if index > list.count {
		return
	}

	var ptr = list.head
	var times = index - 1

	for i := 0; i < times; i++ {
		ptr = ptr.next
	}

	ptr.next = ptr.next.next

	list.count--
}

func (list LinkedList[T]) IndexOf(elem T) int {
	index := 0
	ptr := list.head.next

	for ptr != nil {
		if ptr.data == elem {
			return index
		}

		ptr = ptr.next
		index++
	}

	return -1
}

func (list LinkedList[T]) LastIndexOf(elem T) int {
	i := 0
	index := -1
	ptr := list.head.next

	for ptr != nil {
		if ptr.data == elem {
			index = i
		}

		ptr = ptr.next
		i++
	}

	return index
}

func (list LinkedList[T]) ToString() string {
	var sb strings.Builder
	var ptr = list.head.next

	for ptr != nil {
		sb.WriteString(fmt.Sprintf("%v\n", ptr.data))

		ptr = ptr.next
	}

	return sb.String()
}
