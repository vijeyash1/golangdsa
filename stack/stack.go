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

// Push to add member to top of the stack
func (s *Stack) Push(value int) {
	newStack := &StackNode{
		value: value,
		next:  s.top,
	}
	if s.Isempty() {
		s.top = newStack
		return
	} else {
		s.top = newStack
	}
}

// Traverse to disaply all the stack members
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

// Pop to remove the member at the top
func (s *Stack) Pop() {
	if s.Isempty() {
		fmt.Println("Nothing to pop")
		return
	} else {
		s.top = s.top.next
	}
}

func (s *Stack) Peek() int {
	if s.Isempty() {
		fmt.Println("Nothing to peek")
		return 0
	} else {
		return s.top.next.value
	}
}
