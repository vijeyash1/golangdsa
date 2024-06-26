package circularlinkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type CircularLinkedList struct {
	head *Node
	tail *Node
}

func (cll *CircularLinkedList) IsEmpty() bool {
	return cll.head == nil
}

func (cll *CircularLinkedList) AddData(data int) {
	newNode := &Node{data: data}
	if cll.IsEmpty() {
		cll.head = newNode
		cll.tail = newNode
		newNode.next = newNode
	} else {
		newNode.next = cll.head
		cll.tail.next = newNode
		cll.tail = newNode
	}
}

func (cll *CircularLinkedList) DeleteData(deletevalue int) {
	if cll.IsEmpty() {
		return
	} else {
		current := cll.head
		// specialcase
		if current.data == deletevalue {
			cll.head = current.next
			cll.tail.next = cll.head
			return
		}

		for {
			previous := current
			current = current.next

			if current.data == deletevalue {
				previous.next = current.next
				return
			}
		}
	}

}

func (cll *CircularLinkedList) Display() {
	if cll.IsEmpty() {
		fmt.Println("Circular List is empty")
	} else {
		current := cll.head
		for {
			fmt.Printf("%d", current.data)
			current = current.next
			if current == cll.head {
				break
			}
			if current != nil {
				fmt.Print("->")
			}
		}
	}
	fmt.Println()

}
