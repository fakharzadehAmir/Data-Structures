package minheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {

	data := []float64{10, 5, 15, 20, 8, 3, 25}
	heap := NewMinHeap(data[0])
	for i := 1; i < len(data); i++ {
		heap.Insert(data[i])
	}

	assertHeapProperty(t, heap)
}

func assertHeapProperty(t *testing.T, heap *MinHeap) {
	assert := assert.New(t)

	if heap == nil {
		return
	}

	if heap.LeftChild != nil {
		assert.True(heap.Data <= heap.LeftChild.Data, "Left child is smaller than parent")
		assertHeapProperty(t, heap.LeftChild)
	}

	if heap.RightChild != nil {
		assert.True(heap.Data <= heap.RightChild.Data, "Right child is smaller than parent")
		assertHeapProperty(t, heap.RightChild)
	}
}

func TestExtract(t *testing.T) {
	mh := NewMinHeap(5)
	mh, minValue := mh.Extract()
	assert.Equal(t, float64(5), minValue, "Test 1 failed")
	mh = NewMinHeap(5)
	mh.Insert(8)
	mh.Insert(10)
	_, minValue = mh.Extract()
	assert.Equal(t, float64(5), minValue, "Test 2 failed")
}
