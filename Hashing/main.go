package hashing

// LinkedList represents a node in the linked list.
type LinkedList struct {
	Next *LinkedList
	Data interface{}
}

// Hash represents the hashtable data structure.
type Hash struct {
	HashTable []*LinkedList // Array of linked lists to store the elements.
	Size      int           // Size of the hashtable (number of linked lists in the array).
}

// NewHash creates a new hashtable with the specified size.
func NewHash(size int) *Hash {
	hashTable := make([]*LinkedList, size) // Initialize the array of linked lists with the given size.
	return &Hash{
		HashTable: hashTable, // Set the array as the underlying storage for the hashtable.
		Size:      size,      // Store the size of the hashtable.
	}
}

// Insert adds a new element with the given data to the hashtable.
func (h *Hash) Insert(newData interface{}, hashKey int) {
	// if not closure
	if h.HashTable[hashKey] == nil {
		h.HashTable[hashKey] = &LinkedList{ // Create a new linked list node at the specified index.
			Data: newData,
		}
		return
	}

	// if closure
	newElem := &LinkedList{ // Create a new linked list node.
		Data: newData,
		Next: h.HashTable[hashKey], // Set the new node's next pointer to the current node at the same index.
	}
	h.HashTable[hashKey] = newElem // Update the linked list head to point to the new node.
}

// Delete removes the node containing the given data from the hashtable.
func (h *Hash) Delete(data interface{}, hashKey int) bool {
	// Handle the case when the first node matches the data
	if h.HashTable[hashKey] != nil && h.HashTable[hashKey].Data == data {
		h.HashTable[hashKey] = h.HashTable[hashKey].Next // Move the head of the linked list to the next node.
		return true
	}

	// Traverse the linked list to find the node to delete
	currentNode := h.HashTable[hashKey]
	for currentNode != nil && currentNode.Next != nil {
		if currentNode.Next.Data == data {
			currentNode.Next = currentNode.Next.Next // Remove the node by bypassing it.
			return true
		}
		currentNode = currentNode.Next
	}
	return false // The data was not found in the linked list.
}

// Search checks if the given data exists in the hashtable.
func (h *Hash) Search(data interface{}, hashKey int) bool {
	// Handle the case when the first node matches the data
	if h.HashTable[hashKey] != nil && h.HashTable[hashKey].Data == data {
		return true
	}

	// Traverse the linked list to find the node to find
	currentNode := h.HashTable[hashKey]
	for currentNode != nil && currentNode.Next != nil {
		if currentNode.Next.Data == data {
			return true
		}
		currentNode = currentNode.Next
	}

	return false // The data was not found in the linked list.
}

// Clear removes all elements from the hashtable, effectively resetting it to an empty state.
func (h *Hash) Clear() {
	*h = *NewHash(h.Size) // Create a new empty hashtable and assign it to the current hashtable.
}

// Count returns the total number of elements in the hashtable.
func (h *Hash) Count() int {
	count := 0
	for idx := 0; idx < h.Size; idx++ {
		currentNode := h.HashTable[idx]
		for currentNode != nil {
			count++
			currentNode = currentNode.Next
		}
	}
	return count // Return the total count of elements.
}
