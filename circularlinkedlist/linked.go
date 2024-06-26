package circularlinkedlist

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

	} else {
		cll.tail.next = newNode
		newNode.next = cll.head
		cll.tail = newNode
	}
}
