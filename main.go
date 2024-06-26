package main

import "dsa/singlelinkedlist"

func main() {
	list := singlelinkedlist.LinkedList{}

	list.InsertData(5)
	list.InsertData(6)
	list.InsertData(8)

	list.Display() // 5,6,8

	list.InsertDataAfterValue(7, 6)

	list.Display()                   //5,6,7,8
	list.InsertDataBeforeValue(4, 5) //4,5,6,7,8
	list.Display()
	list.InsertDataBeforeValue(4, 8)
	list.Display() // 4,5,6,7,4,8

}
