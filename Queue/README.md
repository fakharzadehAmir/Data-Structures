# Queue Implementation in Go

This repository contains a simple Queue implementation in Go. The `Queue` struct is designed to store a queue of elements, with each element being of type `interface{}` to allow for flexibility in the type of data stored.

## Queue Operations

### 1. NewQueue

```go
func NewQueue() *Queue
```

The `NewQueue` function returns a new instance of `Queue`.

### 2. Enqueue

```go
func (q *Queue) Enqueue(newData interface{})
```

The `Enqueue` method adds an element to the end of the queue.

### 3. Dequeue

```go
func (q *Queue) Dequeue() interface{}
```

The `Dequeue` method removes and returns the element from the front of the queue. If the queue is empty, it returns `nil`.

### 4. FrontValue

```go
func (q *Queue) FrontValue() interface{}
```

The `FrontValue` method returns the value of the element at the front of the queue without removing it. If the queue is empty, it returns `nil`.

### 5. RearValue

```go
func (q *Queue) RearValue() interface{}
```

The `RearValue` method returns the value of the element at the rear (end) of the queue without removing it. If the queue is empty, it returns `nil`.

### 6. IsEmpty

```go
func (q *Queue) IsEmpty() bool
```

The `IsEmpty` method checks if the queue is empty and returns `true` if it is, otherwise `false`.

### 7. Count

```go
func (q *Queue) Count() int
```

The `Count` method returns the number of elements currently in the queue.

### 8. Clear

```go
func (q *Queue) Clear()
```

The `Clear` method removes all elements from the queue, making it empty.

### 9. Contains

```go
func (q *Queue) Contains(data interface{}) bool
```

The `Contains` method checks if a given element is present in the queue and returns `true` if it exists, otherwise `false`.

### 10. IndexOf

```go
func (q *Queue) IndexOf(data interface{}) int
```

The `IndexOf` method returns the index of the first occurrence of a given element in the queue. If the element is not found, it returns `-1`.

### 11. LastIndexOf

```go
func (q *Queue) LastIndexOf(data interface{}) int
```

The `LastIndexOf` method returns the index of the last occurrence of a given element in the queue. If the element is not found, it returns `-1`.

## Running Tests

To run the tests for the Queue implementation, use the following command:

```bash
go test -v
```

The tests cover various scenarios, including enqueueing and dequeueing elements, checking the front and rear values, verifying the queue's emptiness, and testing additional edge cases.

Feel free to explore the code and run the tests to see how the Queue implementation works!