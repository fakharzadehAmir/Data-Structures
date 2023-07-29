package bst

import (
	"fmt"
	"math"
)

type Node struct {
	Data       float64
	RightChild *Node
	LeftChild  *Node
}

type BinarySearchTree struct {
	Root *Node
}

func NewBST() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) Insertion(newData float64) {
	if bst.Root == nil {
		bst.Root = &Node{
			Data: newData,
		}
		return
	}
	currentNode := bst.Root
	for {
		if currentNode.Data <= newData {
			if currentNode.RightChild == nil {
				currentNode.RightChild = &Node{
					Data: newData,
				}
				return
			}
			currentNode = currentNode.RightChild
		} else {
			if currentNode.LeftChild == nil {
				currentNode.LeftChild = &Node{
					Data: newData,
				}
				return
			}
			currentNode = currentNode.LeftChild
		}
	}
}

func (bst *BinarySearchTree) Deletion(data float64) error {
	var deleteNode, parentNode *Node
	currentNode := bst.Root
	for currentNode != nil {
		if currentNode.Data == data {
			deleteNode = currentNode
			break
		} else if currentNode.Data > data {
			parentNode = currentNode
			currentNode = currentNode.LeftChild
		} else {
			parentNode = currentNode
			currentNode = currentNode.RightChild
		}
	}

	if deleteNode == nil {
		return fmt.Errorf("%f is not in the BST", data)
	}

	if deleteNode.LeftChild == nil && deleteNode.RightChild == nil {
		if parentNode == nil {
			bst.Root = nil
		} else {
			if parentNode.LeftChild == deleteNode {
				parentNode.LeftChild = nil
			} else {
				parentNode.RightChild = nil
			}
		}
	} else if deleteNode.LeftChild == nil {
		if parentNode == nil {
			bst.Root = nil
		} else {
			deleteNode.Data = deleteNode.RightChild.Data
			deleteNode.RightChild = nil
		}
	} else if deleteNode.RightChild == nil {
		if parentNode == nil {
			bst.Root = nil
		} else {
			deleteNode.Data = deleteNode.LeftChild.Data
			deleteNode.LeftChild = nil
		}
	} else {
		successor := bst.Successor(data)
		bst.Deletion(successor)
		deleteNode.Data = successor
	}
	return nil
}

func (bst *BinarySearchTree) Search(data float64) (bool, *Node) {
	currentNode := bst.Root
	for currentNode != nil && currentNode.Data != data {
		if currentNode.RightChild == nil && currentNode.LeftChild == nil {
			return false, nil
		}
		if currentNode.Data > data {
			currentNode = currentNode.LeftChild
		} else if currentNode.Data <= data {
			currentNode = currentNode.RightChild
		}
	}
	if currentNode == nil {
		return false, nil
	}
	return true, currentNode
}

func (bst *BinarySearchTree) Minimum() float64 {
	current := bst.Root
	for current.LeftChild != nil {
		current = current.LeftChild
	}
	return current.Data
}

func (bst *BinarySearchTree) Maximum() float64 {
	current := bst.Root
	for current.RightChild != nil {
		current = current.RightChild
	}
	return current.Data
}

func (bst *BinarySearchTree) Height() int {
	return int(findMaxHeight(bst.Root))
}

func findMaxHeight(tree *Node) float64 {
	if tree == nil {
		return 0
	}
	return math.Max(findMaxHeight(tree.LeftChild), findMaxHeight(tree.RightChild)) + 1
}

func (bst *BinarySearchTree) InOrderTraverse(captrue func(float64)) {

	var InOrder func(node *Node)
	InOrder = func(node *Node) {
		if node != nil {
			InOrder(node.LeftChild)
			captrue(node.Data)
			InOrder(node.RightChild)
		}
	}
	InOrder(bst.Root)
}

func (bst *BinarySearchTree) PreOrderTraverse(captrue func(float64)) {

	var PreOrder func(node *Node)
	PreOrder = func(node *Node) {
		if node != nil {
			captrue(node.Data)
			PreOrder(node.LeftChild)
			PreOrder(node.RightChild)
		}
	}
	PreOrder(bst.Root)
}

func (bst *BinarySearchTree) PostOrderTraverse(captrue func(float64)) {

	var PostOrder func(node *Node)
	PostOrder = func(node *Node) {
		if node != nil {
			PostOrder(node.LeftChild)
			PostOrder(node.RightChild)
			captrue(node.Data)
		}
	}
	PostOrder(bst.Root)
}

func (bst *BinarySearchTree) Successor(data float64) float64 {
	_, node := bst.Search(data)
	if node == nil {
		return 0.0
	}

	if node.RightChild != nil {
		currentNode := node.RightChild
		for currentNode.LeftChild != nil {
			currentNode = currentNode.LeftChild
		}
		return currentNode.Data
	}

	successor := 0.0
	parentNode := bst.Root

	for parentNode != nil {
		if data < parentNode.Data {
			successor = parentNode.Data
			parentNode = parentNode.LeftChild
		} else if data >= parentNode.Data {
			parentNode = parentNode.RightChild
		}
	}
	return successor
}

func (bst *BinarySearchTree) Predecessor(data float64) float64 {
	_, node := bst.Search(data)
	if node == nil {
		return 0.0
	}

	if node.RightChild != nil {
		currentNode := node.LeftChild
		for currentNode.RightChild != nil {
			currentNode = currentNode.RightChild
		}
		return currentNode.Data
	}

	predecessor := 0.0
	parentNode := bst.Root

	for parentNode != nil {
		if data > parentNode.Data {
			predecessor = parentNode.Data
			parentNode = parentNode.RightChild
		} else if data <= parentNode.Data {
			parentNode = parentNode.LeftChild
		}
	}
	return predecessor
}
