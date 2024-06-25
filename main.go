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
