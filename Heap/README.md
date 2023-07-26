# Heap Implementation in Go

This repository contains implementations for both MinHeap and MaxHeap data structures in Go. A heap is a specialized tree-based data structure that satisfies the heap property, where the key of each node is either always greater than or equal to (MaxHeap) or always less than or equal to (MinHeap) the keys of its children.

## Heap Operations
Both MinHeap and MaxHeap share the following operations:  

### 1. Insert

```go
func (h *Heap) Insert(newData float64)
```  

The `Insert` function adds a new element with the specified data to the Heap while maintaining the respective Heap property (MinHeap or MaxHeap).

### 2. Extract

```go
func (h *Heap) Extract() (*Heap, float64)
```  

The `Extract` function removes and returns the root element (minimum in MinHeap, maximum in MaxHeap) from the Heap, along with the updated Heap after reheapifying.

### 3. Peek

```go
func (h *Heap) Peek() float64
```  

The `Peek` function returns the root element (minimum in MinHeap, maximum in MaxHeap) from the Heap without removing it.

### 4. Delete

```go
func (h *Heap) Delete(data float64)
```  

The `Delete` function removes the specified element from the Heap while preserving the respective Heap property (MinHeap or MaxHeap).  


## Running Test  

To run the tests for the MinHeap and MaxHeap implementations, use the following command:  

```bash
go test -v
```  

The tests cover various scenarios, including inserting elements, extracting the minimum/maximum value, deleting elements, and performing different types of binary tree traversals. The test files use the `github.com/stretchr/testify/assert` package to provide concise and readable test assertions.

Feel free to explore the code and run the tests to see how the MinHeap and MaxHeap implementations work! If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request. Happy coding!