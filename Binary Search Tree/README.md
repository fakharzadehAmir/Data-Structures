# Binary Search Tree in Go

This repository contains a simple implementation of a Binary Search Tree (BST) in Go. A Binary Search Tree is a data structure that allows for efficient searching, insertion, and deletion operations.

## Binary Search Tree Operations

### 1. Insertion

```go
func (bst *BinarySearchTree) Insertion(newData float64)
```

The `Insertion` method inserts a new element with the given value into the binary search tree while maintaining the binary search tree property. If the value already exists in the tree, the method does not insert a duplicate.

### 2. Search

```go
func (bst *BinarySearchTree) Search(data float64) (bool, *Node)
```

The `Search` method searches for a given value in the binary search tree. If the value is found, it returns `true` and a pointer to the node containing the value. If the value is not present in the tree, it returns `false` and a `nil` node pointer.

### 3. Deletion

```go
func (bst *BinarySearchTree) Deletion(data float64) error
```

The `Deletion` method removes a node with the specified value from the binary search tree while maintaining the binary search tree property. If the value is not found in the tree, it returns an error.

### 4. Minimum and Maximum

```go
func (bst *BinarySearchTree) Minimum() float64
func (bst *BinarySearchTree) Maximum() float64
```

The `Minimum` method returns the minimum value in the binary search tree, which is the leftmost node in the tree. The `Maximum` method returns the maximum value in the binary search tree, which is the rightmost node in the tree.

### 5. Height

```go
func (bst *BinarySearchTree) Height() int
```

The `Height` method calculates and returns the height of the binary search tree. The height is the maximum number of edges from the root to any leaf node.


### 6. InOrder Traversal

```go
func (bst *BinarySearchTree) InOrderTraverse(capture func(float64))
```

The `InOrderTraverse` method performs an in-order traversal of the binary search tree and calls the `capture` function on each node's value, resulting in values sorted in ascending order.

### 7. PreOrder Traversal

```go
func (bst *BinarySearchTree) PreOrderTraverse(capture func(float64))
```

The `PreOrderTraverse` method performs a pre-order traversal of the binary search tree and calls the `capture` function on each node's value in the order: node, left subtree, right subtree.

### 8. PostOrder Traversal

```go
func (bst *BinarySearchTree) PostOrderTraverse(capture func(float64))
```

The `PostOrderTraverse` method performs a post-order traversal of the binary search tree and calls the `capture` function on each node's value in the order: left subtree, right subtree, node.


### 9. Successor

```go
func (bst *BinarySearchTree) Successor(data float64) float64
```

The `Successor` method finds the next greater value in the binary search tree after the given value. If no such value exists (e.g., the given value is the maximum), it returns 0.0.

### 10. Predecessor

```go
func (bst *BinarySearchTree) Predecessor(data float64) float64
```

The `Predecessor` method finds the next smaller value in the binary search tree before the given value. If no such value exists (e.g., the given value is the minimum), it returns 0.0.

## Running Tests

To run the tests for the Binary Search Tree package, use the following command:

```bash
go test -v
```

The tests cover various operations, including insertion, search, deletion, traversal, successor, and predecessor, ensuring the correctness of the Binary Search Tree implementation. Additional edge cases are also tested.

Feel free to explore the code and run the tests to see how the Binary Search Tree operations work!