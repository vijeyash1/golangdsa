package singlelinkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) InsertData(data int) {
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

func (l *LinkedList) Display() {
	current := l.head

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}
	fmt.Print("Linked list: ")
	for current.next != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}
