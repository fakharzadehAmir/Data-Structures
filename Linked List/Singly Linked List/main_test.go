package singlelinkedlist

import "testing"

var NODE Node

func CheckOrder(list1, list2 []int) bool {
	for index := range list1 {
		if list1[index] != list2[index] {
			return false
		}
	}
	return true
}

func TestCreateNewLinkedList(t *testing.T) {
	NODE := CreateNewLinkedList()
	if NODE == nil {
		t.Errorf("Expected node not to be null")
		return
	}
	if NODE.Data != 0 {
		t.Errorf("Expected NODE.Data to be 0, got %d", NODE.Data)
	}
	if NODE.Next != nil {
		t.Error("Expected NODE.Next to be nil")
	}
}

func TestInsertNewData(t *testing.T) {
	TestCreateNewLinkedList(t)
	NODE.InsertNewData(10)
	NODE.InsertNewData(20)
	NODE.InsertNewData(30)

	currentNode := NODE.Next
	expectedData := []int{10, 20, 30}

	for i, expected := range expectedData {
		if currentNode.Data != expected {
			t.Errorf("Expected node at index %d to have data %d, got %d", i, expected, currentNode.Data)
		}
		currentNode = currentNode.Next
	}

	if currentNode != nil {
		t.Error("Expected end of the linked list, got non-nil node")
	}
}

func TestSearchData(t *testing.T) {
	TestInsertNewData(t)
	exist := NODE.SearchNode(10)
	if !exist {
		t.Errorf("10 exists in Linked List")
	}
	exist = NODE.SearchNode(20)
	if !exist {
		t.Errorf("20 exists in Linked List")
	}
	exist = NODE.SearchNode(50)
	if exist {
		t.Errorf("50 doesn't exist in Linked List")
	}
}

func TestDelete(t *testing.T) {
	TestInsertNewData(t)
	if deleteNode := NODE.DeleteNode(10); deleteNode != 10 {
		t.Errorf("10 was deleted")
	}
	if deleteNode := NODE.DeleteNode(40); deleteNode != -1 {
		t.Errorf("40 doesn't exist in Linked List")
	}
}

func TestTraverseLinkedList(t *testing.T) {
	TestInsertNewData(t)
	traverseList := NODE.TraverseLinkedList()
	expected := []int{10, 20, 30}
	if !CheckOrder(traverseList, expected) {
		t.Errorf("There is a problem in traversing the linked list")
	}
}
