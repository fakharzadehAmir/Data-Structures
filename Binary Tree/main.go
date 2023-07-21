package binarytree

type Node struct {
	Data       interface{}
	LeftChild  *Node
	RightChild *Node
}

type BinaryTree struct {
	Root *Node
}

func (bt *BinaryTree) Insert(newData interface{}) {
	newNode := &Node{Data: newData}
	if bt.Root == nil {
		bt.Root = newNode
		return
	}
	var checkNodes []*Node
	checkNodes = append(checkNodes, bt.Root)
	for len(checkNodes) > 0 {
		currentNode := checkNodes[0]
		checkNodes = checkNodes[1:]

		if currentNode.LeftChild == nil {
			currentNode.LeftChild = newNode
			return
		} else {
			checkNodes = append(checkNodes, currentNode.LeftChild)
		}

		if currentNode.RightChild == nil {
			currentNode.RightChild = newNode
			return
		} else {
			checkNodes = append(checkNodes, currentNode.RightChild)
		}
	}
}

func (bt *BinaryTree) Search(root *Node, data interface{}) bool {
	if root == nil {
		return false
	}
	if root.Data == data {
		return true
	}
	if bt.Search(root.LeftChild, data) {
		return true
	}
	return bt.Search(root.RightChild, data)
}

func (bt *BinaryTree) Size(root *Node) int {
	if root == nil {
		return 0
	}

	checkNodes := []*Node{root}
	size := 0

	for len(checkNodes) > 0 {
		currentNode := checkNodes[0]
		checkNodes = checkNodes[1:]
		size++

		if currentNode.LeftChild != nil {
			checkNodes = append(checkNodes, currentNode.LeftChild)
		}

		if currentNode.RightChild != nil {
			checkNodes = append(checkNodes, currentNode.RightChild)
		}
	}

	return size
}

func (bt *BinaryTree) InOrderTraverse(root *Node, capture func(interface{})) {
	if root != nil {
		bt.InOrderTraverse(root.LeftChild, capture)
		capture(root.Data)
		bt.InOrderTraverse(root.RightChild, capture)
	}
}

func (bt *BinaryTree) PreOrderTraverse(root *Node, capture func(interface{})) {
	if root != nil {
		capture(root.Data)
		bt.PreOrderTraverse(root.LeftChild, capture)
		bt.PreOrderTraverse(root.RightChild, capture)
	}
}

func (bt *BinaryTree) PostOrderTraverse(root *Node, capture func(interface{})) {
	if root != nil {
		bt.PostOrderTraverse(root.LeftChild, capture)
		bt.PostOrderTraverse(root.RightChild, capture)
		capture(root.Data)
	}
}

func (bt *BinaryTree) LevelOrderTraverse(capture func(interface{})) {
	if bt.Root == nil {
		return
	}
	var checkNodes []Node
	idx := 0
	checkNodes = append(checkNodes, *bt.Root)
	for len(checkNodes) != idx {
		capture(checkNodes[idx].Data)
		if checkNodes[idx].LeftChild != nil {
			checkNodes = append(checkNodes, *checkNodes[idx].LeftChild)
		}
		if checkNodes[idx].RightChild != nil {
			checkNodes = append(checkNodes, *checkNodes[idx].RightChild)
		}
		idx++
	}
}

func (bt *BinaryTree) DiagonalTraverse(capture func(interface{})) {
	if bt.Root == nil {
		return
	}
	checkNodes := []*Node{bt.Root}
	for len(checkNodes) > 0 {
		currentNode := checkNodes[0]
		checkNodes = checkNodes[1:]

		for currentNode != nil {
			capture(currentNode.Data)

			if currentNode.LeftChild != nil {
				checkNodes = append(checkNodes, currentNode.LeftChild)
			}

			currentNode = currentNode.RightChild
		}
	}
}

func (bt *BinaryTree) BoundaryTraverse(capture func(interface{})) {
	if bt.Root == nil {
		return
	}

	leftBoundary := func(leftRoot *Node) {
		for leftRoot != nil {
			if leftRoot.LeftChild != nil || leftRoot.RightChild != nil {
				capture(leftRoot.Data)
			}
			if leftRoot.LeftChild != nil {
				leftRoot = leftRoot.LeftChild
			} else {
				leftRoot = leftRoot.RightChild
			}
		}
	}

	rightBoundary := func(rightRoot *Node) {
		if rightRoot == nil {
			return
		}
		var rightBoundaryNodes []Node
		currentNode := rightRoot.RightChild
		for currentNode != nil {
			if currentNode.LeftChild != nil || currentNode.RightChild != nil {
				rightBoundaryNodes = append(rightBoundaryNodes, *currentNode)
			}
			if currentNode.RightChild != nil {
				currentNode = currentNode.RightChild
			} else {
				currentNode = currentNode.LeftChild
			}
		}

		for i := len(rightBoundaryNodes) - 1; i >= 0; i-- {
			capture(rightBoundaryNodes[i].Data)
		}
	}
	var leafBoundary func(root *Node)
	leafBoundary = func(root *Node) {
		if root == nil {
			return
		}
		if root.LeftChild == nil && root.RightChild == nil {
			capture(root.Data)
		}
		leafBoundary(root.LeftChild)
		leafBoundary(root.RightChild)
	}

	capture(bt.Root.Data)
	leftBoundary(bt.Root.LeftChild)
	leafBoundary(bt.Root.LeftChild)
	leafBoundary(bt.Root.RightChild)
	rightBoundary(bt.Root)
}
