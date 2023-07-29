# Binary Tree Implementation in Go  

This repository contains a simple `Binary Tree` implementation in Go. The BinaryTree struct is designed to store a binary tree of nodes, and the `Node` struct represents individual nodes with data of type `interface{}` to allow for flexibility in the type of data stored.  

## Binary Tree Operations  

### 1. Insert

```go
func (bt *BinaryTree) Insert(newData interface{})
```  

The `Insert` function adds a new node with the specified data to the binary tree. It follows the rules of a binary tree to determine the appropriate position for the new node.  

### 2. Search

```go
func (bt *BinaryTree) Search(root *Node, data interface{}) bool
```  

The `Search` function searches for a node with the specified data in the binary tree. It starts the search from the given `root` node and recursively traverses the binary tree to find the data.  

### 3. Delete

```go
func (bt *BinaryTree) Delete(data interface{}) bool
```  

The `Delete` function removes a node with the specified data from the binary tree. If the node to be deleted has both left and right children, it finds the rightmost node in the left subtree (or the leftmost node in the right subtree) to replace the deleted node and maintains the binary tree structure.  

### 4. Height

```go
func (bt *BinaryTree) Height(root *Node) int
```  

The `Height` function calculates and returns the height of the binary tree. The height is the number of edges in the longest path from the root node to a leaf node.  

### 5. Size

```go
func (bt *BinaryTree) Size() int
```  

The `Size` function returns the number of nodes in the binary tree. It counts all the nodes in the binary tree, including the root.  

### 6. IsBalanced

```go
func (bt *BinaryTree) IsBalanced(root *Node) bool
```  

The `IsBalanced` function checks if the binary tree is balanced. A balanced binary tree is a tree in which the difference in heights between the left and right subtrees of every node is not more than one.  

### 7. InOrderTraverse

```go
func (bt *BinaryTree) InOrderTraverse(root *Node, capture func(interface{}))
```  

The `InOrderTraverse` function performs an in-order traversal of the binary tree and calls the `capture` function on each node's data. In-order traversal visits the left subtree, then the current node, and finally the right subtree.  

### 8. PreOrderTraverse

```go
func (bt *BinaryTree) PreOrderTraverse(root *Node, capture func(interface{}))
```  

The `PreOrderTraverse` function performs a pre-order traversal of the binary tree and calls the `capture` function on each node's data. Pre-order traversal visits the current node, then the left subtree, and finally the right subtree.

  

### 9. PostOrderTraverse

```go
func (bt *BinaryTree) PostOrderTraverse(root *Node, capture func(interface{}))
```  

The `PostOrderTraverse` function performs a post-order traversal of the binary tree and calls the `capture` function on each node's data. Post-order traversal visits the left subtree, then the right subtree, and finally the current node.  

### 10. LevelOrderTraverse

```go
func (bt *BinaryTree) LevelOrderTraverse(capture func(interface{}))
```  

The `LevelOrderTraverse` function performs a level-order traversal of the binary tree and calls the `capture` function on each node's data. Level-order traversal visits all the nodes at the same level before moving to the next level.  

### 11. DiagonalTraverse

```go
func (bt *BinaryTree) DiagonalTraverse(capture func(interface{}))
```  

The `DiagonalTraverse` function performs a diagonal traversal of the binary tree and calls the `capture` function on each node's data. Diagonal traversal visits all the nodes in the same diagonal before moving to the next diagonal.  

### 12. BoundaryTraverse

```go
func (bt *BinaryTree) BoundaryTraverse(capture func(interface{}))
```  

The `BoundaryTraverse` function performs a boundary traversal of the binary tree and calls the `capture` function on each node's data. Boundary traversal visits the left boundary, all leaf nodes, and the right boundary in a clockwise manner. 



## Running Test  

To run the tests for the Binary Tree implementation, use the following command:  
```bash
go test -v
```  

The tests cover various scenarios, including inserting nodes, searching for data, deleting nodes, and performing different types of traversals.  

Feel free to explore the code and run the tests to see how the Binary Tree implementation works!  