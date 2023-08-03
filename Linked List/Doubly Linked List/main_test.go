package doubly_linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertBegin(t *testing.T) {
	dll := NewDoublyLinkedList()

	dll.InsertBegin(1)
	assert.Equal(t, 1, dll.Root.Data)
	assert.Equal(t, 1, dll.Tail.Data)

	dll.InsertBegin(2)
	assert.Equal(t, 2, dll.Root.Data)
	assert.Equal(t, 1, dll.Tail.Data)
	assert.Equal(t, 1, dll.Root.Next.Data)
	assert.Equal(t, 2, dll.Tail.Prev.Data)
}

func TestInsertEnd(t *testing.T) {
	dll := NewDoublyLinkedList()

	dll.InsertEnd(3)
	assert.Equal(t, 3, dll.Root.Data)
	assert.Equal(t, 3, dll.Tail.Data)

	dll.InsertEnd(4)
	assert.Equal(t, 3, dll.Root.Data)
	assert.Equal(t, 4, dll.Tail.Data)
	assert.Equal(t, 4, dll.Root.Next.Data)
	assert.Equal(t, 3, dll.Tail.Prev.Data)
}

func TestInsertAfter(t *testing.T) {
	dll := NewDoublyLinkedList()

	dll.InsertEnd(1)
	dll.InsertAfter(2, 1)
	assert.Equal(t, 2, dll.Root.Next.Data)
	assert.Equal(t, 1, dll.Tail.Prev.Data)

	dll.InsertAfter(3, 2)
	assert.Equal(t, 3, dll.Root.Next.Next.Data)
	assert.Equal(t, 2, dll.Tail.Prev.Data)
}

func TestInsertBefore(t *testing.T) {
	dll := NewDoublyLinkedList()

	dll.InsertBegin(1)
	dll.InsertBefore(2, 1)
	assert.Equal(t, 2, dll.Root.Data)
	assert.Equal(t, 2, dll.Tail.Prev.Data)

	dll.InsertBefore(3, 2)
	assert.Equal(t, 2, dll.Root.Next.Data)
	assert.Equal(t, 3, dll.Tail.Prev.Prev.Data)
}

func TestSearch(t *testing.T) {
	dll := NewDoublyLinkedList()

	dll.InsertBegin(1)
	dll.InsertEnd(3)

	found, node := dll.Search(2)
	assert.False(t, found)
	assert.Nil(t, node)

	found, node = dll.Search(1)
	assert.True(t, found)
	assert.Equal(t, 1, node.Data)
}

func TestSize(t *testing.T) {
	dll := NewDoublyLinkedList()

	dll.InsertBegin(1)
	dll.InsertEnd(3)

	size := dll.Size()
	assert.Equal(t, 2, size)

	dll.InsertBefore(2, 3)
	dll.InsertAfter(4, 1)
	size = dll.Size()
	assert.Equal(t, 4, size)
}

func TestDeletion(t *testing.T) {
	dll := NewDoublyLinkedList()

	dll.InsertBegin(1)
	dll.InsertAfter(2, 1)

	deleted := dll.Deletion(2)
	assert.True(t, deleted)
	assert.Equal(t, 1, dll.Size())
	assert.Nil(t, dll.Root.Next)

	dll.InsertEnd(3)
	dll.InsertBefore(2, 3)
	deleted = dll.Deletion(2)
	assert.True(t, deleted)
	assert.Equal(t, 2, dll.Size())
	assert.Equal(t, 3, dll.Root.Next.Data)
	assert.Equal(t, 1, dll.Tail.Prev.Data)
}
