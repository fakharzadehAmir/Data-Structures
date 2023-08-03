# Doubly Linked List in Go

This repository contains a simple implementation of a Doubly Linked List in Go. A Doubly Linked List is a data structure that consists of a sequence of elements, each containing a reference to both the next and previous elements, allowing for efficient traversal in both directions.

## Doubly Linked List Operations

### 1. Insertion at the Beginning

```go
func (dll *DoublyLinkedList) InsertBegin(newData interface{})
```

The `InsertBegin` method inserts a new element with the given value at the beginning of the doubly linked list. If the list is empty, the new element becomes both the `Root` and `Tail`.

### 2. Insertion at the End

```go
func (dll *DoublyLinkedList) InsertEnd(newData interface{})
```

The `InsertEnd` method inserts a new element with the given value at the end of the doubly linked list. If the list is empty, the new element becomes both the `Root` and `Tail`.

### 3. Insertion After a Node

```go
func (dll *DoublyLinkedList) InsertAfter(newData, data interface{})
```

The `InsertAfter` method inserts a new element with the given value after a specified node containing the `data`. If the node with the `data` is not found, the new element is inserted at the end.

### 4. Insertion Before a Node

```go
func (dll *DoublyLinkedList) InsertBefore(newData, data interface{})
```

The `InsertBefore` method inserts a new element with the given value before a specified node containing the `data`. If the node with the `data` is not found, the new element is inserted at the beginning.

### 5. Search

```go
func (dll *DoublyLinkedList) Search(data interface{}) (bool, *DoublyNode)
```

The `Search` method searches for a given value in the doubly linked list. If the value is found, it returns `true` and a pointer to the node containing the value. If the value is not present in the list, it returns `false` and a `nil` node pointer.

### 6. Deletion

```go
func (dll *DoublyLinkedList) Deletion(data interface{}) bool
```

The `Deletion` method removes a node with the specified value from the doubly linked list. If the value is not found in the list, it returns `false`.

### 7. Size

```go
func (dll *DoublyLinkedList) Size() int
```

The `Size` method returns the number of elements in the doubly linked list.

## Running Tests

To run the tests for the Doubly Linked List package, use the following command:

```bash
go test -v
```

The tests cover various operations, including insertion at the beginning, insertion at the end, insertion after a node, insertion before a node, search, deletion, and size calculation, ensuring the correctness of the Doubly Linked List implementation. Additional edge cases are also tested.

Feel free to explore the code and run the tests to see how the Doubly Linked List operations work!