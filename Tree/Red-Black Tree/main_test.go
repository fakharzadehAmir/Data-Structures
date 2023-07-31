package redblack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	rbt := NewRBTree()
	rbt.Root = &Node{Data: 5, Color: false}
	rbt.Root.LeftChild = &Node{Data: 3, Color: true}
	rbt.Root.RightChild = &Node{Data: 8, Color: true}
	rbt.Root.LeftChild.LeftChild = &Node{Data: 1, Color: false}
	rbt.Root.LeftChild.RightChild = &Node{Data: 4, Color: false}
	rbt.Root.RightChild.LeftChild = &Node{Data: 7, Color: false}
	rbt.Root.RightChild.RightChild = &Node{Data: 10, Color: false}

	tests := []struct {
		data   float64
		exists bool
		node   *Node
	}{
		{data: 5, exists: true, node: rbt.Root},
		{data: 4, exists: true, node: rbt.Root.LeftChild.RightChild},
		{data: 9, exists: false, node: nil},
		{data: 7, exists: true, node: rbt.Root.RightChild.LeftChild},
		{data: 12, exists: false, node: nil},
	}

	for _, test := range tests {
		exists, node := rbt.Search(test.data)
		assert.Equal(t, test.exists, exists)
		assert.Equal(t, test.node, node)
	}
}

func TestInsertion(t *testing.T) {
	rbt := NewRBTree()
	assert.Nil(t, rbt.Root)
	rbt.Insertion(50)
	rbt.Insertion(30)
	rbt.Insertion(70)
	rbt.Insertion(20)
	rbt.Insertion(40)
	rbt.Insertion(60)
	rbt.Insertion(80)
	rbt.Insertion(10)
	rbt.Insertion(90)
	rbt.Insertion(55)
	rbt.Insertion(65)

	assert.Equal(t, float64(50), rbt.Root.Data)
	assert.False(t, rbt.Root.Color)

	left := rbt.Root.LeftChild
	assert.Equal(t, float64(30), left.Data)
	assert.True(t, left.Color)

	leftLeft := left.LeftChild
	assert.Equal(t, float64(20), leftLeft.Data)
	assert.False(t, leftLeft.Color)

	leftRight := left.RightChild
	assert.Equal(t, float64(40), leftRight.Data)
	assert.False(t, leftRight.Color)

	right := rbt.Root.RightChild
	assert.Equal(t, float64(70), right.Data)
	assert.True(t, right.Color)
	rightLeft := right.LeftChild
	assert.Equal(t, float64(60), rightLeft.Data)
	assert.False(t, rightLeft.Color)

	rightRight := right.RightChild
	assert.Equal(t, float64(80), rightRight.Data)
	assert.False(t, rightRight.Color)

	assert.Equal(t, float64(10), leftLeft.LeftChild.Data)
	assert.True(t, leftLeft.LeftChild.Color)

	assert.Equal(t, float64(90), rightRight.RightChild.Data)
	assert.True(t, rightRight.RightChild.Color)

	assert.Equal(t, float64(55), rightLeft.LeftChild.Data)
	assert.True(t, rightLeft.LeftChild.Color)

	assert.Equal(t, float64(65), rightLeft.RightChild.Data)
	assert.True(t, rightLeft.RightChild.Color)

	assert.True(t, isRedBlackTree(rbt.Root))

}

func isRedBlackTree(node *Node) bool {
	if node == nil {
		return true
	}

	if node.Color == true {
		if node.LeftChild != nil && node.LeftChild.Color == true {
			return false
		}
		if node.RightChild != nil && node.RightChild.Color == true {
			return false
		}
	}

	return isRedBlackTree(node.LeftChild) && isRedBlackTree(node.RightChild)
}

func TestDeletion(t *testing.T) {
	rbt := NewRBTree()
	assert.Nil(t, rbt.Root)

	numbers := []float64{50, 30, 70, 20, 40, 60, 80, 10, 90, 55, 65}
	for _, num := range numbers {
		rbt.Insertion(num)
	}

	rbt.Deletion(90)
	rbt.Deletion(10)
	rbt.Deletion(70)

	_, node90 := rbt.Search(90)
	_, node10 := rbt.Search(10)
	_, node70 := rbt.Search(70)

	assert.Nil(t, node90)
	assert.Nil(t, node10)
	assert.Nil(t, node70)

	assert.True(t, isRedBlackTree(rbt.Root))
}

func TestMinMax(t *testing.T) {
	rbt := NewRBTree()
	assert.Nil(t, rbt.Root)
	rbt.Insertion(50)
	rbt.Insertion(30)
	rbt.Insertion(70)
	rbt.Insertion(20)
	rbt.Insertion(40)
	rbt.Insertion(60)
	rbt.Insertion(80)
	rbt.Insertion(10)
	rbt.Insertion(90)
	rbt.Insertion(55)
	rbt.Insertion(65)
	assert.Equal(t, float64(90), rbt.Maximum())
	assert.Equal(t, float64(10), rbt.Minimum())

}

func TestSuccessorAndPredecessor(t *testing.T) {
	rbt := NewRBTree()
	rbt.Insertion(50)
	rbt.Insertion(30)
	rbt.Insertion(70)
	rbt.Insertion(20)
	rbt.Insertion(40)
	rbt.Insertion(60)
	rbt.Insertion(80)

	assert.Equal(t, float64(60), rbt.Successor(50))
	assert.Equal(t, float64(50), rbt.Successor(40))
	assert.Equal(t, float64(80), rbt.Successor(70))
	assert.Equal(t, float64(0), rbt.Successor(80))

	assert.Equal(t, float64(40), rbt.Predecessor(50))
	assert.Equal(t, float64(30), rbt.Predecessor(40))
	assert.Equal(t, float64(50), rbt.Predecessor(60))
	assert.Equal(t, float64(0), rbt.Predecessor(20))
}

func TestTraversals(t *testing.T) {
	rbt := NewRBTree()
	rbt.Insertion(50)
	rbt.Insertion(30)
	rbt.Insertion(70)
	rbt.Insertion(20)
	rbt.Insertion(40)
	rbt.Insertion(60)
	rbt.Insertion(80)

	var inorderResult []float64
	rbt.InOrderTraverse(func(data float64) {
		inorderResult = append(inorderResult, data)
	})

	var preorderResult []float64
	rbt.PreOrderTraverse(func(data float64) {
		preorderResult = append(preorderResult, data)
	})

	var postorderResult []float64
	rbt.PostOrderTraverse(func(data float64) {
		postorderResult = append(postorderResult, data)
	})

	expectedInorder := []float64{20, 30, 40, 50, 60, 70, 80}
	expectedPreorder := []float64{50, 30, 20, 40, 70, 60, 80}
	expectedPostorder := []float64{20, 40, 30, 60, 80, 70, 50}

	assert.Equal(t, expectedInorder, inorderResult)
	assert.Equal(t, expectedPreorder, preorderResult)
	assert.Equal(t, expectedPostorder, postorderResult)
}
