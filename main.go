package main

import (
	"fmt"
	"linear_list"
)

func main() {
	list := linear_list.SequentialList[int]{}

	list.Init()

	for i := 0; i < 10; i++ {
		list.Add(i + 1)
	}

	fmt.Println(list.ToString())

	list.RemoveAt(5)

	fmt.Println(list.ToString())
}
