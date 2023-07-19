# Single Linked List in Go

This repository contains a basic implementation of a single linked list in Go. A linked list is a data structure where elements, called nodes, are connected through pointers. Each node contains data and a pointer to the next node in the list.

## Code Structure

The main code for the single linked list is located in the `singlelinkedlist` package and consists of the following:

- `Node`: A struct representing a single linked list node. Each node contains an integer data field and a pointer to the next node in the list.
- `CreateNewLinkedList()`: A function to create a new empty linked list with a default node.
- `InsertNewData(newData int)`: A method to insert a new node with the given data at the end of the linked list.
- `SearchNode(data int) bool`: A method to search for a node with the specified data in the linked list and return true if found, false otherwise.
- `DeleteNode(data int) int`: A method to delete the first occurrence of a node with the specified data from the linked list and return the deleted data value. If the node is not found, it returns -1.
- `TraverseLinkedList() []int`: A method to traverse the linked list and return a slice containing the data values of all nodes in the list.

## Running Tests

The test file `main_test.go` contains unit tests for the single linked list implementation. The tests cover the various functionalities of the linked list, including creating a new linked list, inserting new nodes, searching for nodes, deleting nodes, and traversing the linked list.

To run the tests, use the following command:

```shell
go test -v
```  



## Test Cases

The test cases cover the following scenarios:

- TestCreateNewLinkedList: Ensures that a new linked list is created correctly with a default node and its attributes are set accordingly.
- TestInsertNewData: Tests the insertion of new nodes into the linked list and verifies the correct ordering of nodes.
- TestSearchData: Checks the search functionality to find nodes with specific data values in the linked list.
- TestDelete: Tests the deletion of nodes from the linked list and verifies the correct deletion behavior.
- TestTraverseLinkedList: Ensures that the linked list is traversed correctly, and the data values of all nodes are collected in the correct order.
