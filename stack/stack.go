// Last in first out principle
// example: arranging plates in a kitchen

package stack

import "fmt"

type StackNode struct {
	value int
	next  *StackNode
}

type Stack struct {
	top *StackNode
}

func (s *Stack) Isempty() bool {
	return s.top == nil
}
func (s *Stack) Push(value int) {
	newStack := &StackNode{
		value: value,
	}
	if s.Isempty() {
		s.top = newStack
		return
	} else {
		newStack.next = s.top
		s.top = newStack
	}
}

func (s *Stack) Traverse() {

	if s.Isempty() {
		fmt.Println("Stack is empty")
		return
	} else {
		for {
			current := s.top
			fmt.Printf("--> %d", current.value)
			current = current.next
			if current == nil {
				break
			}
		}
	}
}
