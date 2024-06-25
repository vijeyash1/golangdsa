package doublelinkedlist

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) AddNode(data int) {
	newnode := &Node{
		data: data,
		next: nil,
		prev: nil,
	}

	if dll.head == nil {
		dll.head = newnode
		dll.tail = newnode
	} else {
		newnode.prev = dll.tail
		dll.tail.next = newnode
		dll.tail = newnode
	}
}

func (dll *DoublyLinkedList) PrintForward() {
	currentNode := dll.head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (dll *DoublyLinkedList) PrintReverse() {
	currentNode := dll.tail
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.prev
	}
	fmt.Println()
}
