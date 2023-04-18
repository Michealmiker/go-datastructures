package main

import (
	"fmt"
	"linear_list"
)

func main() {
	testLinkedList()
}

func testLinkedList() {
	list := linear_list.LinkedList[int]{}

	list.Init()

	for i := 0; i < 10; i++ {
		list.Add(i + 1) // list.AddFirst(i + 1)
	}

	fmt.Println(list.ToString())

	list.Remove(5)
	// list.RemoveAt(5)

	fmt.Println(list.ToString())
}

func testSequentialList() {
	list := linear_list.SequentialList[int]{}

	list.Init()

	for i := 0; i < 10; i++ {
		list.Add(i + 1) // list.AddFirst(i + 1)
	}

	fmt.Println(list.ToString())

	list.Remove(5)
	// list.RemoveAt(5)

	fmt.Println(list.ToString())
}
