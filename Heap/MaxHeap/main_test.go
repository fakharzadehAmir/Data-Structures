package maxheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	data := []float64{10, 5, 15, 20, 8, 3, 25}
	heap := NewMaxHeap(data[0])
	for i := 1; i < len(data); i++ {
		heap.Insert(data[i])
	}

	assertHeapProperty(t, heap)
}

func assertHeapProperty(t *testing.T, heap *MaxHeap) {
	assert := assert.New(t)

	if heap == nil {
		return
	}

	if heap.LeftChild != nil {
		assert.True(heap.Data >= heap.LeftChild.Data, "Left child is greater than parent")
		assertHeapProperty(t, heap.LeftChild)
	}

	if heap.RightChild != nil {
		assert.True(heap.Data >= heap.RightChild.Data, "Right child is greater than parent")
		assertHeapProperty(t, heap.RightChild)
	}
}

func TestExtract(t *testing.T) {
	mh := NewMaxHeap(5)
	mh, maxValue := mh.Extract()
	assert.Equal(t, float64(5), maxValue, "Test 1 failed")
	mh = NewMaxHeap(5)
	mh.Insert(8)
	mh.Insert(10)
	_, maxValue = mh.Extract()
	assert.Equal(t, float64(10), maxValue, "Test 2 failed")
}

func TestPeek(t *testing.T) {
	mh := NewMaxHeap(5)
	maxValue := mh.Peek()
	assert.Equal(t, float64(5), maxValue, "Test 1 failed")
	mh = NewMaxHeap(2)
	mh.Insert(3)
	mh.Insert(4)
	maxValue = mh.Peek()
	assert.Equal(t, float64(4), maxValue, "Test 2 failed")
}

func TestDelete(t *testing.T) {
	data := []float64{10, 5, 15, 20, 8, 3, 25}
	heap := NewMaxHeap(data[0])
	for i := 1; i < len(data); i++ {
		heap.Insert(data[i])
	}

	heap.Delete(15)
	assertHeapProperty(t, heap)

	heap.Delete(3)
	assertHeapProperty(t, heap)

	heap.Delete(10)
	assertHeapProperty(t, heap)
}
