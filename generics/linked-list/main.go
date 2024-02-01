package main

import (
	"fmt"
)

type LinkedListElement[T any] struct {
	value T
	next  *LinkedListElement[T]
}

type LinkedList[T any] struct {
	head *LinkedListElement[T]
	tail *LinkedListElement[T]
}

func (linkedList *LinkedList[T]) push(value T) {
	newElement := LinkedListElement[T]{
		value: value,
	}

	if linkedList.head == nil {
		linkedList.head = &newElement
		linkedList.tail = linkedList.head
	} else {
		linkedList.tail.next = &newElement
		linkedList.tail = linkedList.tail.next
	}
}

func (linkedList *LinkedList[T]) getAll() []T {
	items := []T{}

	if linkedList.head != nil {
		currentNode := linkedList.head
		for currentNode != nil {
			items = append(items, currentNode.value)
			currentNode = currentNode.next
		}
	}

	return items
}

func (linkedList *LinkedList[T]) getAllWithFancyForLoop() []T {
	var items []T

	for item := linkedList.head; item != nil; item = item.next {
		items = append(items, item.value)
	}

	return items
}

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println("-- Go Linked List --")
	listOfInts := LinkedList[int]{}
	listOfInts.push(1)
	listOfInts.push(2)
	listOfInts.push(3)
	listOfInts.push(4)
	fmt.Println(listOfInts)
	fmt.Println(listOfInts.getAll())
	fmt.Println(listOfInts.getAllWithFancyForLoop())
}
