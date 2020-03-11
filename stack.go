package stack

import "sync"

// Stack defines a thread safe stack
type Stack struct {
	mutex    sync.Mutex
	elements []interface{}
}

// New initialises a new Stack
func New() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
	}
}

// Push adds elements to the stack
func (s *Stack) Push(elements ...interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.elements = append(s.elements, elements...)
}

// Size returns the total number of elements in the stack
func (s *Stack) Size() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return len(s.elements)
}

// Pop returns the top element and nil if the stack is empty
func (s *Stack) Pop() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.elements) == 0 {
		return nil
	}

	var (
		topIndex = len(s.elements) - 1
		element  = s.elements[topIndex]
	)

	s.elements[topIndex] = nil
	s.elements = s.elements[:topIndex]

	return element
}
