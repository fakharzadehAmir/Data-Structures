package singlelinkedlist

type Node struct {
	Data int
	Next *Node
}

func CreateNewLinkedList() *Node {
	return &Node{}
}

func (n *Node) InsertNewData(newData int) {
	newNode := &Node{
		Data: newData,
	}

	if n == nil {
		*n = *newNode
		return
	}

	currentNode := n
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}
	currentNode.Next = newNode
}

func (n *Node) SearchNode(data int) bool {
	currentNode := n
	for currentNode != nil {
		if currentNode.Data == data {
			return true
		}
		currentNode = currentNode.Next
	}
	return false
}

func (n *Node) DeleteNode(data int) int {
	currentNode := n.Next
	previousNode := n
	for currentNode != nil {
		if currentNode.Data == data {
			previousNode.Next = currentNode.Next
			return data
		}
		currentNode = currentNode.Next
		previousNode = previousNode.Next
	}
	return -1
}

func (n *Node) TraverseLinkedList() []int {
	var traverseList []int
	currentNode := n.Next
	for currentNode != nil {
		traverseList = append(traverseList, currentNode.Data)
		currentNode = currentNode.Next
	}
	return traverseList
}
