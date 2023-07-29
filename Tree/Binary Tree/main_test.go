package binarytree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	var binaryTree BinaryTree
	binaryTree.Insert(1)
	binaryTree.Insert(2)
	binaryTree.Insert(3)
	binaryTree.Insert(4)
	binaryTree.Insert(5)
	binaryTree.Insert(6)
	var result []interface{}
	captureOutput := func(data interface{}) {
		result = append(result, data)
	}
	binaryTree.LevelOrderTraverse(captureOutput)
	expected := []interface{}{1, 2, 3, 4, 5, 6}
	assert.Equal(t, result,
		expected,
		fmt.Sprintf("Insert level-by-level it must print %v\n but got %v", expected, result))

}

func TestSearch(t *testing.T) {
	binaryTree := &BinaryTree{
		Root: &Node{
			Data: 1,
		},
	}
	binaryTree.Root.LeftChild = &Node{Data: 2}
	binaryTree.Root.RightChild = &Node{Data: 3}
	binaryTree.Root.LeftChild.LeftChild = &Node{Data: 4}
	binaryTree.Root.LeftChild.RightChild = &Node{Data: 5}
	binaryTree.Root.RightChild.LeftChild = &Node{Data: 6}

	// Test 1
	result := binaryTree.Search(binaryTree.Root, 3)
	expected := true
	assert.Equal(t, result,
		expected,
		fmt.Sprintf("For searching 3 must print %v\n but got %v", expected, result))

	// Test 2
	result = binaryTree.Search(binaryTree.Root, 5)
	expected = true
	assert.Equal(t, result,
		expected,
		fmt.Sprintf("For searching 5 must print %v\n but got %v", expected, result))

	// Test 3
	result = binaryTree.Search(binaryTree.Root, 7)
	expected = false
	assert.Equal(t, result,
		expected,
		fmt.Sprintf("For searching 7 must print %v\n but got %v", expected, result))

	// Test 4
	result = binaryTree.Search(binaryTree.Root, 8)
	expected = false
	assert.Equal(t, result,
		expected,
		fmt.Sprintf("For searching 8 must print %v\n but got %v", expected, result))
}

func TestDelete(t *testing.T) {
	var binaryTree BinaryTree
	binaryTree.Insert(1)
	binaryTree.Insert(2)
	binaryTree.Insert(3)
	binaryTree.Insert(4)
	binaryTree.Insert(5)
	binaryTree.Insert(6)
	binaryTree.Insert(7)

	// Test 1
	binaryTree.Delete(2)
	var result []interface{}
	captureOutput := func(data interface{}) {
		result = append(result, data)
	}
	binaryTree.LevelOrderTraverse(captureOutput)
	expected := []interface{}{1, 7, 3, 4, 5, 6}
	size := binaryTree.Size()
	expectedSize := 6
	assert.Equal(t, size,
		expectedSize,
		fmt.Sprintf("Size of this binary tree must be %v\n but got %v", expectedSize, size))

	assert.Equal(t, result,
		expected,
		fmt.Sprintf("After deletion of this binary tree must be %v\n but got %v", expected, result))

	// Test 2
	binaryTree.Delete(3)
	result = nil
	binaryTree.LevelOrderTraverse(captureOutput)
	expected = []interface{}{1, 7, 6, 4, 5}
	size = binaryTree.Size()
	expectedSize = 5
	assert.Equal(t, size,
		expectedSize,
		fmt.Sprintf("Size of this binary tree must be %v\n but got %v", expectedSize, size))

	assert.Equal(t, result,
		expected,
		fmt.Sprintf("After deletion of this binary tree must be %v\n but got %v", expected, result))

	// Test 3
	binaryTree.Delete(6)
	result = nil
	binaryTree.LevelOrderTraverse(captureOutput)
	expected = []interface{}{1, 7, 4, 5}
	size = binaryTree.Size()
	expectedSize = 4
	assert.Equal(t, size,
		expectedSize,
		fmt.Sprintf("Size of this binary tree must be %v\n but got %v", expectedSize, size))

	assert.Equal(t, result,
		expected,
		fmt.Sprintf("After deletion of this binary tree must be %v\n but got %v", expected, result))

}

func TestHeight(t *testing.T) {
	var binaryTree BinaryTree
	binaryTree.Insert(1)
	binaryTree.Insert(2)
	binaryTree.Insert(3)
	binaryTree.Insert(4)
	binaryTree.Insert(5)
	binaryTree.Insert(6)
	result := binaryTree.Height(binaryTree.Root)
	expected := 3
	assert.Equal(t, result,
		expected,
		fmt.Sprintf("Height of this binary tree must be %v\n but got %v", expected, result))
}

func TestSize(t *testing.T) {
	var binaryTree BinaryTree
	binaryTree.Insert(1)
	binaryTree.Insert(2)
	binaryTree.Insert(3)
	binaryTree.Insert(4)
	binaryTree.Insert(5)
	binaryTree.Insert(6)
	// Test 1
	size := binaryTree.Size()
	expected := 6
	assert.Equal(t, size,
		expected,
		fmt.Sprintf("Size of this binary tree must be %v\n but got %v", expected, size))

	// Test 2
	for i := 0; i < 30000; i++ {
		binaryTree.Insert(i + 7)
	}

	size = binaryTree.Size()
	expected = 30006
	assert.Equal(t, size,
		expected,
		fmt.Sprintf("Size of this binary tree must be %v\n but got %v", expected, size))

}

func TestIsBalanced(t *testing.T) {
	var binaryTree BinaryTree
	binaryTree.Insert(1)
	binaryTree.Insert(2)
	binaryTree.Insert(3)
	binaryTree.Insert(4)
	binaryTree.Insert(5)
	binaryTree.Insert(6)

	// Test 1
	result := binaryTree.IsBalanced(binaryTree.Root)
	expected := true
	assert.Equal(t, result,
		expected,
		fmt.Sprintf("Being balanced of this binary tree must be %v\n but got %v", expected, result))

	// Test 2
	binaryTree.Root.RightChild.LeftChild.LeftChild = &Node{Data: 7}
	result = binaryTree.IsBalanced(binaryTree.Root)
	expected = false
	assert.Equal(t, result,
		expected,
		fmt.Sprintf("Being balanced of this binary tree must be %v\n but got %v", expected, result))

}

func TestTraverse(t *testing.T) {
	binaryTree := &BinaryTree{
		Root: &Node{
			Data: 1,
		},
	}
	binaryTree.Root.LeftChild = &Node{Data: 2}
	binaryTree.Root.RightChild = &Node{Data: 3}
	binaryTree.Root.LeftChild.LeftChild = &Node{Data: 4}
	binaryTree.Root.LeftChild.RightChild = &Node{Data: 5}
	binaryTree.Root.RightChild.LeftChild = &Node{Data: 6}

	var result []interface{}
	captureOutput := func(data interface{}) {
		result = append(result, data)
	}
	// Test InOrder
	binaryTree.InOrderTraverse(binaryTree.Root, captureOutput)
	expected := []interface{}{4, 2, 5, 1, 6, 3}
	assert.Equal(t, result,
		expected,
		fmt.Sprintf("InOrder traverse must print %v\n but got %v", expected, result))

	// Test PreOrder
	result = nil
	binaryTree.PreOrderTraverse(binaryTree.Root, captureOutput)
	expected = []interface{}{1, 2, 4, 5, 3, 6}
	assert.Equal(t, result,
		expected,
		fmt.Sprintf("PreOrder traverse must print %v\n but got %v", expected, result))

	// Test PostOrder
	result = nil
	binaryTree.PostOrderTraverse(binaryTree.Root, captureOutput)
	expected = []interface{}{4, 5, 2, 6, 3, 1}
	assert.Equal(t, result,
		expected,
		fmt.Sprintf("PostOrder traverse must print %v\n but got %v", expected, result))

	// Test LevelOrder
	result = nil
	binaryTree.LevelOrderTraverse(captureOutput)
	expected = []interface{}{1, 2, 3, 4, 5, 6}
	assert.Equal(t, result,
		expected,
		fmt.Sprintf("LevelOrder traverse must print %v\n but got %v", expected, result))

	// Test Diagonal
	result = nil
	binaryTree.DiagonalTraverse(captureOutput)
	expected = []interface{}{1, 3, 2, 5, 6, 4}
	assert.Equal(t, result,
		expected,
		fmt.Sprintf("Diagonal traverse must print %v\n but got %v", expected, result))

	// Test Boundary
	binaryTree.Root.RightChild.RightChild = &Node{Data: 7}
	binaryTree.Root.LeftChild.LeftChild.LeftChild = &Node{Data: 8}
	binaryTree.Root.LeftChild.RightChild.LeftChild = &Node{Data: 9}
	binaryTree.Root.LeftChild.RightChild.RightChild = &Node{Data: 10}
	// Test 1
	result = nil
	binaryTree.BoundaryTraverse(captureOutput)
	expected = []interface{}{1, 2, 4, 8, 9, 10, 6, 7, 3}
	assert.Equal(t, result,
		expected,
		fmt.Sprintf("Boundary traverse must print %v\n but got %v", expected, result))

	// Test 2
	binaryTree.Root.RightChild.RightChild.LeftChild = &Node{Data: 11}
	binaryTree.Root.RightChild.RightChild.RightChild = &Node{Data: 12}
	result = nil
	binaryTree.BoundaryTraverse(captureOutput)
	expected = []interface{}{1, 2, 4, 8, 9, 10, 6, 11, 12, 7, 3}
	assert.Equal(t, result,
		expected,
		fmt.Sprintf("Boundary traverse must print %v\n but got %v", expected, result))
}
