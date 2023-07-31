# Red-Black Tree in Go

This repository contains a simple implementation of a Red-Black Tree in Go. A Red-Black Tree is a self-balancing binary search tree that ensures that the tree remains balanced, resulting in efficient search, insertion, and deletion operations.

## Red-Black Tree Operations

### 1. Insertion

```go
func (rbt *RedBlackTree) Insertion(newData float64)
```

The `Insertion` method inserts a new element with the given value into the Red-Black Tree while maintaining the Red-Black Tree properties. After insertion, the tree is balanced to ensure that it conforms to the properties of a Red-Black Tree.

### 2. Search

```go
func (rbt *RedBlackTree) Search(data float64) (bool, *Node)
```

The `Search` method searches for a given value in the Red-Black Tree. If the value is found, it returns `true` and a pointer to the node containing the value. If the value is not present in the tree, it returns `false` and a `nil` node pointer.

### 3. Deletion

```go
func (rbt *RedBlackTree) Deletion(data float64)
```

The `Deletion` method removes a node with the specified value from the Red-Black Tree while maintaining the Red-Black Tree properties. After deletion, the tree is rebalanced to ensure that it conforms to the properties of a Red-Black Tree.

### 4. Minimum and Maximum

```go
func (rbt *RedBlackTree) Minimum() float64
func (rbt *RedBlackTree) Maximum() float64
```

The `Minimum` method returns the minimum value in the Red-Black Tree, which is the leftmost node in the tree. The `Maximum` method returns the maximum value in the Red-Black Tree, which is the rightmost node in the tree.

### 5. InOrder Traversal

```go
func (rbt *RedBlackTree) InOrderTraverse(capture func(float64))
```

The `InOrderTraverse` method performs an in-order traversal of the Red-Black Tree and calls the `capture` function on each node's value, resulting in values sorted in ascending order.

### 6. PreOrder Traversal

```go
func (rbt *RedBlackTree) PreOrderTraverse(capture func(float64))
```

The `PreOrderTraverse` method performs a pre-order traversal of the Red-Black Tree and calls the `capture` function on each node's value in the order: node, left subtree, right subtree.

### 7. PostOrder Traversal

```go
func (rbt *RedBlackTree) PostOrderTraverse(capture func(float64))
```

The `PostOrderTraverse` method performs a post-order traversal of the Red-Black Tree and calls the `capture` function on each node's value in the order: left subtree, right subtree, node.

### 8. Successor

```go
func (rbt *RedBlackTree) Successor(data float64) float64
```

The `Successor` method finds the next greater value in the Red-Black Tree after the given value. If no such value exists (e.g., the given value is the maximum), it returns 0.0.

### 9. Predecessor

```go
func (rbt *RedBlackTree) Predecessor(data float64) float64
```

The `Predecessor` method finds the next smaller value in the Red-Black Tree before the given value. If no such value exists (e.g., the given value is the minimum), it returns 0.0.

## Running Tests

To run the tests for the Red-Black Tree package, use the following command:

```bash
go test -v
```

The tests cover various operations, including insertion, search, deletion, traversal, successor, and predecessor, ensuring the correctness of the Red-Black Tree implementation. Additional edge cases are also tested.

Feel free to explore the code and run the tests to see how the Red-Black Tree operations work!