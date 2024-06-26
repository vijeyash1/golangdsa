package arraytolinkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Add(data int) {
	new := &Node{data: data}

	if l.head == nil {
		l.head = new
	} else {
		current := l.head

		for current.next != nil {
			current = current.next
		}
		current.next = new
	}
}

func (l *LinkedList) PrintForward() {
	if l.head == nil {
		fmt.Println("list is empty")
		return
	}
	fmt.Println("Linked list:")
	current := l.head
	for current != nil {
		fmt.Printf("%d", current.data)
		current = current.next
		if current != nil {
			fmt.Printf(" ->")
		}
	}
	fmt.Println()
}
func ArrayToLinked(list []int) {
	new := &LinkedList{}
	for _, v := range list {
		new.Add(v)
	}
	new.PrintForward()
}
func PrintLinkedList(head *Node) {
	currentNode := head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}
func SimpleMethod(list []int) *Node {
	var head, tail *Node
	for _, d := range list {
		newNode := &Node{data: d}
		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.next = newNode
			tail = newNode
		}
	}
	return head
}
