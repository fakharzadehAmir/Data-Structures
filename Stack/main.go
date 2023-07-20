package stack

import (
	"fmt"
	"reflect"
)

type MyStack struct {
	data []interface{}
	size int
}

func NewMyStack() *MyStack {
	return &MyStack{}
}

func (ms *MyStack) Push(newData interface{}) error {
	if len(ms.data) > 0 {
		if reflect.TypeOf(ms.data[0]) != reflect.TypeOf(newData) {
			return fmt.Errorf("invalid type: all elements must be of type %T", ms.data[0])
		}
	}
	ms.data = append(ms.data, newData)
	ms.size++
	return nil
}

func (ms *MyStack) Pop() (interface{}, error) {
	if ms.size == 0 {
		*ms = MyStack{}
		return nil, fmt.Errorf("stack is Empty")
	}
	lastElement := ms.data[ms.size-1]
	ms.size--
	return lastElement, nil
}

func (ms *MyStack) Peek() (interface{}, error) {
	if ms.size == 0 {
		return nil, fmt.Errorf("stack is Empty")
	}
	return ms.data[ms.size-1], nil
}

func (ms *MyStack) IsEmpty() bool {
	return ms.size == 0
}

func (ms *MyStack) Clear() {
	if ms.size != 0 && len(ms.data) > 0 {
		*ms = MyStack{}
	}
}
