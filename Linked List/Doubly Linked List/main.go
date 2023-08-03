package doubly_linked_list

type DoublyNode struct {
	Data interface{}
	Next *DoublyNode
	Prev *DoublyNode
}

type DoublyLinkedList struct {
	Root *DoublyNode
	Tail *DoublyNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dll *DoublyLinkedList) InsertBegin(newData interface{}) {
	newNode := &DoublyNode{
		Data: newData,
	}
	if dll.Root == nil {
		dll.Root = newNode
		dll.Tail = newNode
		return
	}
	nextNode := dll.Root
	nextNode.Prev = newNode
	newNode.Next = nextNode
	dll.Root = newNode
}

func (dll *DoublyLinkedList) InsertEnd(newData interface{}) {
	newNode := &DoublyNode{
		Data: newData,
	}
	if dll.Root == nil {
		dll.Root = newNode
		dll.Tail = newNode
		return
	}
	prevNode := dll.Tail
	prevNode.Next = newNode
	newNode.Prev = prevNode
	dll.Tail = newNode
}

func (dll *DoublyLinkedList) InsertAfter(newData, data interface{}) {
	newNode := &DoublyNode{
		Data: newData,
	}
	_, prevNode := dll.Search(data)
	if prevNode == nil {
		dll.InsertEnd(newData)
		return
	}
	nextNode := prevNode.Next
	newNode.Prev = prevNode
	newNode.Next = nextNode
	if nextNode != nil {
		nextNode.Prev = newNode
	}
	prevNode.Next = newNode
	if prevNode == dll.Tail {
		dll.Tail = newNode
	}
}

func (dll *DoublyLinkedList) InsertBefore(newData, data interface{}) {
	newNode := &DoublyNode{
		Data: newData,
	}
	_, nextNode := dll.Search(data)
	if nextNode == nil {
		dll.InsertBegin(newData)
		return
	}
	prevNode := nextNode.Prev
	newNode.Prev = prevNode
	newNode.Next = nextNode
	if prevNode != nil {
		prevNode.Next = newNode
	}
	nextNode.Prev = newNode
	if nextNode == dll.Root {
		dll.Root = newNode
	}
}

func (dll *DoublyLinkedList) Search(data interface{}) (bool, *DoublyNode) {
	currentNode := dll.Root
	for currentNode != nil {
		if currentNode.Data == data {
			return true, currentNode
		}
		currentNode = currentNode.Next
	}
	return false, nil
}

func (dll *DoublyLinkedList) Size() int {
	currentNode := dll.Root
	var size int = 0
	for currentNode != nil {
		size++
		currentNode = currentNode.Next
	}
	return size
}

func (dll *DoublyLinkedList) Deletion(data interface{}) bool {
	found, node := dll.Search(data)
	if node != nil {
		prevNode := node.Prev
		nextNode := node.Next
		if nextNode != nil {
			nextNode.Prev = prevNode
		}
		if prevNode != nil {
			prevNode.Next = nextNode
		}
		if node == dll.Root {
			dll.Root = nextNode
		}
		if node == dll.Tail {
			dll.Tail = prevNode
		}
	}
	return found
}
