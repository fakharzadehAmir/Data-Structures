package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertion(t *testing.T) {

	bst := &BinarySearchTree{}

	bst.Insertion(5)
	bst.Insertion(3)
	bst.Insertion(7)
	bst.Insertion(2)
	bst.Insertion(4)

	assert.NotNil(t, bst.Root)

	assert.Equal(t, 5.0, bst.Root.Data)

	assert.NotNil(t, bst.Root.LeftChild)
	assert.Equal(t, 3.0, bst.Root.LeftChild.Data)

	assert.NotNil(t, bst.Root.RightChild)
	assert.Equal(t, 7.0, bst.Root.RightChild.Data)

	assert.NotNil(t, bst.Root.LeftChild.LeftChild)
	assert.Equal(t, 2.0, bst.Root.LeftChild.LeftChild.Data)

	assert.NotNil(t, bst.Root.LeftChild.RightChild)
	assert.Equal(t, 4.0, bst.Root.LeftChild.RightChild.Data)
}

func TestSearch(t *testing.T) {
	bst := NewBST()
	bst.Insertion(10)
	bst.Insertion(5)
	bst.Insertion(15)
	bst.Insertion(2)
	bst.Insertion(7)
	bst.Insertion(12)
	bst.Insertion(20)

	found, _ := bst.Search(10)
	assert.True(t, found)

	found, _ = bst.Search(5)
	assert.True(t, found)

	found, _ = bst.Search(15)
	assert.True(t, found)

	found, _ = bst.Search(2)
	assert.True(t, found)

	found, _ = bst.Search(7)
	assert.True(t, found)

	found, _ = bst.Search(12)
	assert.True(t, found)

	found, _ = bst.Search(20)
	assert.True(t, found)

	found, _ = bst.Search(1)
	assert.False(t, found)

	found, _ = bst.Search(8)
	assert.False(t, found)

	found, _ = bst.Search(11)
	assert.False(t, found)

	found, _ = bst.Search(25)
	assert.False(t, found)
}

func TestDeletion(t *testing.T) {
	bst := NewBST()

	bst.Insertion(10)
	bst.Insertion(5)
	bst.Insertion(15)
	bst.Insertion(2)
	bst.Insertion(7)
	bst.Insertion(12)
	bst.Insertion(20)

	err := bst.Deletion(2)
	assert.NoError(t, err)
	found, _ := bst.Search(2)
	assert.False(t, found)

	err = bst.Deletion(15)
	assert.NoError(t, err)
	found, _ = bst.Search(15)
	assert.False(t, found)

	err = bst.Deletion(10)
	assert.NoError(t, err)
	found, _ = bst.Search(10)
	assert.False(t, found)

	err = bst.Deletion(5)
	assert.NoError(t, err)
	found, _ = bst.Search(5)
	assert.False(t, found)
}

func TestMinMax(t *testing.T) {
	bst := NewBST()
	bst.Insertion(10)
	bst.Insertion(5)
	bst.Insertion(15)
	bst.Insertion(2)
	bst.Insertion(7)
	bst.Insertion(12)
	bst.Insertion(20)

	assert.Equal(t, 20.0, bst.Maximum())
	assert.Equal(t, 2.0, bst.Minimum())
}

func TestHeight(t *testing.T) {
	bst := NewBST()
	assert.Equal(t, 0, bst.Height())

	bst.Insertion(10)
	assert.Equal(t, 1, bst.Height())

	bst.Insertion(5)
	bst.Insertion(15)
	assert.Equal(t, 2, bst.Height())

	bst.Insertion(2)
	bst.Insertion(7)
	bst.Insertion(12)
	bst.Insertion(20)
	assert.Equal(t, 3, bst.Height())
}

func TestTraverse(t *testing.T) {
	bst := NewBST()

	bst.Insertion(5)
	bst.Insertion(3)
	bst.Insertion(7)
	bst.Insertion(2)
	bst.Insertion(4)
	bst.Insertion(6)
	bst.Insertion(8)

	// InOrder Traverse: [2, 3, 4, 5, 6, 7, 8]
	var inOrderResult []float64
	bst.InOrderTraverse(func(data float64) {
		inOrderResult = append(inOrderResult, data)
	})
	assert.Equal(t, []float64{2, 3, 4, 5, 6, 7, 8}, inOrderResult)

	// PreOrder Traverse: [5, 3, 2, 4, 7, 6, 8]
	var preOrderResult []float64
	bst.PreOrderTraverse(func(data float64) {
		preOrderResult = append(preOrderResult, data)
	})
	assert.Equal(t, []float64{5, 3, 2, 4, 7, 6, 8}, preOrderResult)

	// PostOrder Traverse: [2, 4, 3, 6, 8, 7, 5]
	var postOrderResult []float64
	bst.PostOrderTraverse(func(data float64) {
		postOrderResult = append(postOrderResult, data)
	})
	assert.Equal(t, []float64{2, 4, 3, 6, 8, 7, 5}, postOrderResult)
}

func TestSuccessorAndPredecessor(t *testing.T) {
	bst := NewBST()
	bst.Insertion(10)
	bst.Insertion(5)
	bst.Insertion(15)
	bst.Insertion(2)
	bst.Insertion(7)
	bst.Insertion(12)
	bst.Insertion(20)

	successorTests := []struct {
		input    float64
		expected float64
	}{
		{input: 2, expected: 5},
		{input: 5, expected: 7},
		{input: 7, expected: 10},
		{input: 10, expected: 12},
		{input: 12, expected: 15},
		{input: 15, expected: 20},
		{input: 20, expected: 0},
		{input: 25, expected: 0},
	}

	predecessorTests := []struct {
		input    float64
		expected float64
	}{
		{input: 2, expected: 0},
		{input: 5, expected: 2},
		{input: 7, expected: 5},
		{input: 10, expected: 7},
		{input: 12, expected: 10},
		{input: 15, expected: 12},
		{input: 20, expected: 15},
		{input: 25, expected: 0},
	}

	for _, test := range successorTests {
		successor := bst.Successor(test.input)
		assert.Equal(t, test.expected, successor, "Successor of %v should be %v", test.input, test.expected)
	}

	for _, test := range predecessorTests {
		predecessor := bst.Predecessor(test.input)
		assert.Equal(t, test.expected, predecessor, "Predecessor of %v should be %v", test.input, test.expected)
	}
}
