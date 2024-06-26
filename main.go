package main

import "dsa/circularlinkedlist"

func main() {
	list := circularlinkedlist.CircularLinkedList{}

	list.AddData(5)
	list.AddData(6)
	list.AddData(8)
	list.AddData(9)
	list.AddData(10)

	list.Display()

	list.DeleteData(9)

	list.Display()

}
