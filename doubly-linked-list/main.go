package main

import (
	"fmt"

	"github.com/u-shylianok/golang-basics/doubly-linked-list/list"
)

func main() {
	fmt.Println("create new list")
	l := list.New()
	printList(l, true)

	fmt.Println("add values")
	l.InsertBack(20)
	l.InsertFront(10)
	l.InsertBeforeFromBack(10, 11)
	l.InsertAfterFromFront(11, 21)
	l.InsertFront(10)
	fmt.Println("expected: 10, 11, 21, 10, 20")
	printList(l, true)

	fmt.Println("delete from back: 10")
	l.DeleteFromBack(10)
	printList(l, false)

	fmt.Println("insert back: 10")
	l.InsertBack(10)
	printList(l, false)

	fmt.Println("delete from front: 10")
	l.DeleteFromFront(10)
	printList(l, false)

	fmt.Println("delete front and back")
	l.DeleteBack()
	l.DeleteFront()
	printList(l, true)

	fmt.Println("delete front and back")
	l.DeleteBack()
	l.DeleteFront()
	printList(l, true)
}

func printList(list *list.List, isReverseNeeded bool) {
	fmt.Printf("--> %v \n", list.ToSlice())
	if isReverseNeeded {
		fmt.Printf("<-- %v \n", list.ToSliceReverse())
	}
	fmt.Println()
}
