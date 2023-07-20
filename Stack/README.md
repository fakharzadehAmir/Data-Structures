# Stack Implementation in Go

This repository contains a simple Stack implementation in Go. The `MyStack` struct is designed to store a stack of elements, with each element being of type `interface{}` to allow for flexibility in the type of data stored.

## Stack Operations

### 1. NewMyStack

```go
func NewMyStack() *MyStack
```  

The `NewMyStack` function returns a new instance of `MyStack`.  

### 2. Push  

```go
func (ms *MyStack) Push(newData interface{}) error
```  

The `Push` function adds an element to the top of the stack. It checks the type of the new element and ensures that all elements in the stack have the same type. If the new element has a different type from the existing elements, it returns an error.  

### 3. Pop  

```go
func (ms *MyStack) Pop() (interface{}, error)
```  

The `Pop` function removes and returns the top element from the stack. If the stack is empty, it returns an error.  

### 4. Peek  

```go
func (ms *MyStack) Peek() (interface{}, error)
```  

The `Peek` function returns the top element from the stack without removing it. If the stack is empty, it returns an error.  

### 5. IsEmpty  

```go
func (ms *MyStack) IsEmpty() bool
```  

The `IsEmpty` function returns `true` if the stack is empty and `false` otherwise.  

### 6. Clear  

```go
func (ms *MyStack) Clear()
```  

The `Clear` function clears all elements from the stack, making it empty.   

## Running Tests  

To run the tests for the Stack implementation, use the following command:
```bash
go test -v
```  

The tests cover various scenarios, including pushing elements of different types, popping elements from an empty stack, and checking the size and emptiness of the stack after various operations.

Feel free to explore the code and run the tests to see how the Stack implementation works!  


**Note:** This implementation uses an `interface{}` to store elements, allowing for flexibility in the type of data. However, this also means that type safety is not guaranteed at compile-time, and you need to handle type assertions when working with elements retrieved from the stack.