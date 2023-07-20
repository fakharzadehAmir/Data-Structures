package stack

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	var stack MyStack
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	for index := range stack.data {
		if stack.data[index] != index+1 {
			t.Errorf("Push function for stack doesn't work correctly! index(%d)", index)
		}
	}

	stack = MyStack{}
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	if stack.data[0] != "a" {
		t.Errorf("Push function for stack doesn't work correctly! index(0)")
	}
	if stack.data[1] != "b" {
		t.Errorf("Push function for stack doesn't work correctly! index(1)")
	}
	if stack.data[2] != "c" {
		t.Errorf("Push function for stack doesn't work correctly! index(2)")
	}

	stack = MyStack{}
	stack.Push(false)
	stack.Push(true)
	stack.Push(true)
	if stack.data[0] != false {
		t.Errorf("Push function for stack doesn't work correctly! index(0)")
	}
	if stack.data[1] != true {
		t.Errorf("Push function for stack doesn't work correctly! index(1)")
	}
	if stack.data[2] != true {
		t.Errorf("Push function for stack doesn't work correctly! index(2)")
	}
}

func TestStackPushError(t *testing.T) {
	var stack MyStack
	stack.Push(1)
	err := stack.Push("a")
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func TestStackPop(t *testing.T) {
	stack := &MyStack{
		data: []interface{}{1, 2, 3, 4, 5},
		size: len([]interface{}{1, 2, 3, 4, 5}),
	}
	lastElem, _ := stack.Pop()
	if lastElem != 5 && stack.size != 4 {
		t.Errorf("lastElement must be 5 but it's %v, and size of stack must be 4 but it's %v", lastElem, stack.size)
	}
	stack.Pop()
	stack.Pop()
	lastElem, _ = stack.Pop()
	if lastElem != 2 && stack.size != 1 {
		t.Errorf("lastElement must be 2 but it's %v, and size of stack must be 1 but it's %v", lastElem, stack.size)
	}

	stack = &MyStack{
		data: []interface{}{"a", "b", "c", "d"},
		size: len([]interface{}{"a", "b", "c", "d"}),
	}
	lastElem, _ = stack.Pop()
	if lastElem != "d" && stack.size != 3 {
		t.Errorf("lastElement must be d but it's %v, and size of stack must be 3 but it's %v", lastElem, stack.size)
	}
	stack.Pop()
	stack.Pop()
	lastElem, _ = stack.Pop()
	if lastElem != "a" && stack.size != 0 {
		t.Errorf("lastElement must be a but it's %v, and size of stack must be 0 but it's %v", lastElem, stack.size)
	}

}
func TestStackPeek(t *testing.T) {
	stack := &MyStack{
		data: []interface{}{1, 2, 3, 4, 5},
		size: len([]interface{}{1, 2, 3, 4, 5}),
	}
	lastElem, _ := stack.Peek()
	if lastElem != 5 && stack.size != 5 {
		t.Errorf("lastElement must be 5 but it's %v, and size of stack must be 5 but it's %v", lastElem, stack.size)
	}
	stack.Peek()
	stack.Peek()
	stack.Peek()
	stack.Peek()
	stack.Peek()
	lastElem, _ = stack.Peek()
	if lastElem != 5 && stack.size != 5 {
		t.Errorf("lastElement must be 5 but it's %v, and size of stack must be 5 but it's %v", lastElem, stack.size)
	}
}
func TestStackPopPeekError(t *testing.T) {
	stack := &MyStack{
		data: []interface{}{1, 2, 3, 4, 5},
		size: len([]interface{}{1, 2, 3, 4, 5}),
	}
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	_, err := stack.Pop()
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
	_, err = stack.Peek()
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
	stack.Push(1)
	_, err = stack.Peek()
	if err != nil {
		t.Errorf("Stack is not empty")
	}
	_, err = stack.Pop()
	if err != nil {
		t.Errorf("Stack is not empty")
	}
	_, err = stack.Pop()
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func TestIsEmpty(t *testing.T) {
	stack := &MyStack{
		data: []interface{}{1, 2},
		size: len([]interface{}{1, 2}),
	}
	if stack.IsEmpty() {
		t.Errorf("Stack is not empty")
	}
	stack.Pop()
	stack.Pop()
	if !stack.IsEmpty() {
		t.Errorf("Stack is empty")
	}
	stack.Pop()
	if !stack.IsEmpty() {
		t.Errorf("Stack is empty")
	}
}

func TestClear(t *testing.T) {
	stack := &MyStack{
		data: []interface{}{1, 2},
		size: len([]interface{}{1, 2}),
	}
	if stack.size == 0 {
		t.Errorf("Stack is not empty")
	}
	stack.Clear()
	if stack.size != 0 {
		t.Errorf("Stack is empty")
	}
}
