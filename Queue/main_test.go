package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	assert.Equal(t, 1, queue.FrontValue(), "FrontValue should return 1 after Enqueue")
	assert.Equal(t, 3, queue.RearValue(), "RearValue should return 3 after Enqueue")
}

func TestDequeue(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	dequeueVal := queue.Dequeue()

	assert.Equal(t, 1, dequeueVal, "Dequeue should return 1")
	assert.Equal(t, 2, queue.FrontValue(), "FrontValue should return 2 after Dequeue")
	assert.Equal(t, 3, queue.RearValue(), "RearValue should return 3 after Dequeue")

	dequeueVal = queue.Dequeue()
	assert.Equal(t, 2, dequeueVal, "Dequeue should return 2")
	assert.Equal(t, 3, queue.FrontValue(), "FrontValue should return 3 after second Dequeue")
	assert.Equal(t, 3, queue.RearValue(), "RearValue should return 3 after second Dequeue")

	dequeueVal = queue.Dequeue()
	assert.Equal(t, 3, dequeueVal, "Dequeue should return 3")
	assert.True(t, queue.IsEmpty(), "Queue should be empty after third Dequeue")
}

func TestIsEmpty(t *testing.T) {
	queue := NewQueue()

	assert.True(t, queue.IsEmpty(), "Newly created queue should be empty")

	queue.Enqueue(1)
	assert.False(t, queue.IsEmpty(), "Queue should not be empty after Enqueue")
}

func TestCount(t *testing.T) {
	queue := NewQueue()

	assert.Equal(t, 0, queue.Count(), "Newly created queue should have zero elements")

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	assert.Equal(t, 3, queue.Count(), "Queue should have 3 elements after 3 Enqueue operations")
}

func TestClear(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	queue.Clear()

	assert.True(t, queue.IsEmpty(), "Queue should be empty after Clear")
	assert.Equal(t, 0, queue.Count(), "Count should be 0 after Clear")
	assert.Nil(t, queue.Dequeue(), "Dequeue on an empty queue after Clear should return nil")

	queue.Enqueue(4)
	assert.Equal(t, 4, queue.FrontValue(), "FrontValue after Enqueue after Clear should return 4")
	assert.Equal(t, 4, queue.RearValue(), "RearValue after Enqueue after Clear should return 4")
}

func TestContains(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	assert.True(t, queue.Contains(2), "Contains should return true for an element present in the queue")
	assert.False(t, queue.Contains(4), "Contains should return false for an element not present in the queue")
}

func TestIndexOf(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	assert.Equal(t, 1, queue.IndexOf(2), "IndexOf should return 1 for the element 2")
	assert.Equal(t, -1, queue.IndexOf(4), "IndexOf should return -1 for an element not present in the queue")
}

func TestLastIndexOf(t *testing.T) {
	queue := NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(2)

	assert.Equal(t, 3, queue.LastIndexOf(2), "LastIndexOf should return 3 for the last occurrence of element 2")
	assert.Equal(t, -1, queue.LastIndexOf(4), "LastIndexOf should return -1 for an element not present in the queue")
}

func TestQueueEdgeCases(t *testing.T) {
	queue := NewQueue()

	assert.Nil(t, queue.Dequeue(), "Dequeue on an empty queue should return nil")

	assert.Nil(t, queue.FrontValue(), "FrontValue on an empty queue should return nil")
	assert.Nil(t, queue.RearValue(), "RearValue on an empty queue should return nil")

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Clear()
	assert.True(t, queue.IsEmpty(), "Queue should be empty after Clear")
	assert.Equal(t, 0, queue.Count(), "Count should be 0 after Clear")
	assert.Nil(t, queue.Dequeue(), "Dequeue on an empty queue after Clear should return nil")
}
