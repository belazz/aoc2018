package datastruct

import (
	"log"
)

type LinkedList struct {
	Head *Node
}

type Node struct {
	Value int
	next  *Node
}

/*
	Counts the amount of nodes inside linked list
*/
func (list LinkedList) Count() int {
	count := 0
	traversalNode := list.Head
	for traversalNode != nil {
		count++
		traversalNode = traversalNode.next
	}

	return count
}

/*
	Inserts element at position @index

*/
func (list LinkedList) InsertAt(index int, newNode *Node) LinkedList {
	if list.Head == nil && index == 0 {
		list.Head = newNode
	}

	if index == 0 {
		newNode.next = list.Head
		list.Head = newNode

		return list
	}

	if index == list.Count()-1 {
		list = list.Append(newNode)
		return list
	}
	currentIndex := 0
	traversalNode := list.Head
	for currentIndex != index-1 {
		currentIndex++
		traversalNode = traversalNode.next
		if traversalNode.next == nil {
			log.Panic("Specified index is out of bounds")
		}
	}

	newNode.next = traversalNode.next
	traversalNode.next = newNode

	return list
}

func (list LinkedList) DeleteAt(index int) LinkedList {
	// delete first item
	if index == 0 {
		list.Head = list.Head.next
		return list
	}

	// delete last item
	if index == list.Count()-1 {
		currentIndex := 0
		traversalNode := list.Head
		for currentIndex != index-1 {
			currentIndex++
			traversalNode = traversalNode.next
		}
		traversalNode.next = nil
		return list
	}

	// delete item somewhere in between
	currentIndex := 0
	traversalNode := list.Head
	for currentIndex != index-1 {
		currentIndex++
		traversalNode = traversalNode.next
	}

	traversalNode.next = traversalNode.next.next
	return list
}

func (list LinkedList) Append(newNode *Node) LinkedList {
	if list.Head == nil {
		list.Head = newNode
		return list
	}
	traversalNode := list.Head
	for traversalNode.next != nil {
		traversalNode = traversalNode.next
	}

	traversalNode.next = newNode

	return list
}
