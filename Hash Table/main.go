package hashing

type LinkedList struct {
	Next *LinkedList
	Data interface{}
}

type Hash struct {
	HashTable []*LinkedList
	Size      int
}

func NewHash(size int) *Hash {
	hashTable := make([]*LinkedList, size)
	return &Hash{
		HashTable: hashTable,
		Size:      size,
	}
}

func (h *Hash) Insert(newData interface{}, hashKey int) {
	if h.HashTable[hashKey] == nil {
		h.HashTable[hashKey] = &LinkedList{
			Data: newData,
		}
		return
	}

	newElem := &LinkedList{
		Data: newData,
		Next: h.HashTable[hashKey],
	}
	h.HashTable[hashKey] = newElem
}

func (h *Hash) Delete(data interface{}, hashKey int) bool {
	if h.HashTable[hashKey] != nil && h.HashTable[hashKey].Data == data {
		h.HashTable[hashKey] = h.HashTable[hashKey].Next
		return true
	}

	currentNode := h.HashTable[hashKey]
	for currentNode != nil && currentNode.Next != nil {
		if currentNode.Next.Data == data {
			currentNode.Next = currentNode.Next.Next
			return true
		}
		currentNode = currentNode.Next
	}
	return false
}

func (h *Hash) Search(data interface{}, hashKey int) bool {
	if h.HashTable[hashKey] != nil && h.HashTable[hashKey].Data == data {
		return true
	}

	currentNode := h.HashTable[hashKey]
	for currentNode != nil && currentNode.Next != nil {
		if currentNode.Next.Data == data {
			return true
		}
		currentNode = currentNode.Next
	}

	return false
}

func (h *Hash) Clear() {
	*h = *NewHash(h.Size)
}

func (h *Hash) Count() int {
	count := 0
	for idx := 0; idx < h.Size; idx++ {
		currentNode := h.HashTable[idx]
		for currentNode != nil {
			count++
			currentNode = currentNode.Next
		}
	}
	return count
}
