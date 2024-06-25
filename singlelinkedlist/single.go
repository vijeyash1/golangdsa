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
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func (l *LinkedList) AddtoHead(data int) {
	new := &Node{
		data: data,
	}
	if l.head == nil {
		l.head = new
	} else {
		current := l.head
		new.next = current
		l.head = new
	}
}

func (l *LinkedList) Addtotail(data int) {
	new := &Node{
		data: data,
	}
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

func (l *LinkedList) DeleteNodeWithValue(data int) {
	if l.head == nil {
		fmt.Println("List is empty")
	}

	current := l.head
	next := current.next
	if current.data == data {
		l.head = next
		return
	}

	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
		} else {
			current = current.next
		}
	}
}

/*
package main

import "dsa/singlelinkedlist"

func main() {
	list := singlelinkedlist.LinkedList{}

	list.InsertData(5)
	list.InsertData(6)
	list.InsertData(7)
	list.InsertData(8)

	list.Display()

	list.DeleteNodeWithValue(7)

	list.Display()
	list.DeleteNodeWithValue(7)
	list.Display()
	list.DeleteNodeWithValue(5)
	list.Display()

}

*/
/*
output

Linked list: 5 6 7 8
Linked list: 5 6 8
Linked list: 5 6 8
Linked list: 6 8

*/
