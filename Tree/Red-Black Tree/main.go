package redblack

type Node struct {
	Data       float64
	LeftChild  *Node
	RightChild *Node
	Parent     *Node
	Color      bool
}

type RedBlackTree struct {
	Root *Node
}

func NewRBTree() *RedBlackTree {
	return &RedBlackTree{}
}

func (rbt *RedBlackTree) Insertion(newData float64) {
	newNode := &Node{
		Data:  newData,
		Color: true,
	}
	if rbt.Root == nil {
		rbt.Root = newNode
		rbt.Root.Color = false
	} else {
		insertNode(rbt.Root, newNode, rbt.fixedAfterInsertion)
	}
}

func insertNode(parentNode, newNode *Node, fixed func(*Node)) {
	if parentNode.Data > newNode.Data {
		if parentNode.LeftChild == nil {
			parentNode.LeftChild = newNode
			newNode.Parent = parentNode
			fixed(newNode)
			return
		}
		insertNode(parentNode.LeftChild, newNode, fixed)
	} else {
		if parentNode.RightChild == nil {
			parentNode.RightChild = newNode
			newNode.Parent = parentNode
			fixed(newNode)
			return
		}
		insertNode(parentNode.RightChild, newNode, fixed)
	}
}

func (rbt *RedBlackTree) fixedAfterInsertion(newNode *Node) {
	for newNode.Parent != nil && newNode.Parent.Color {
		grandparent := newNode.Parent.Parent
		if grandparent == nil {
			return
		}

		if newNode.Parent == grandparent.LeftChild {
			uncle := grandparent.RightChild
			if uncle != nil && uncle.Color {
				newNode.Parent.Color = false
				uncle.Color = false
				grandparent.Color = true
				newNode = grandparent
			} else {
				if newNode == newNode.Parent.RightChild {
					newNode = newNode.Parent
					rbt.leftRotate(newNode)
				}
				newNode.Parent.Color = false
				grandparent.Color = true
				rbt.rightRotate(grandparent)
			}
		} else {
			uncle := grandparent.LeftChild
			if uncle != nil && uncle.Color {
				newNode.Parent.Color = false
				uncle.Color = false
				grandparent.Color = true
				newNode = grandparent
			} else {
				if newNode == newNode.Parent.LeftChild {
					newNode = newNode.Parent
					rbt.rightRotate(newNode)
				}
				newNode.Parent.Color = false
				grandparent.Color = true
				rbt.leftRotate(grandparent)
			}
		}
	}

	rbt.Root.Color = false
}

func (rbt *RedBlackTree) Deletion(data float64) {
	find, node := rbt.Search(data)
	if !find {
		return
	}

	var x *Node
	y := node
	yOriginalColor := y.Color

	if node.LeftChild == nil {
		x = node.RightChild
		rbt.transplant(node, node.RightChild)
	} else if node.RightChild == nil {
		x = node.LeftChild
		rbt.transplant(node, node.LeftChild)
	} else {
		findMin := func(node *Node) *Node {
			current := node
			for current.LeftChild != nil {
				current = current.LeftChild
			}
			return current
		}
		y = findMin(node.RightChild)
		yOriginalColor = y.Color
		x = y.RightChild

		if y.Parent == node {
			if x != nil {
				x.Parent = y
			}
		} else {
			rbt.transplant(y, y.RightChild)
			y.RightChild = node.RightChild
			if y.RightChild != nil {
				y.RightChild.Parent = y
			}
		}

		rbt.transplant(node, y)
		y.LeftChild = node.LeftChild
		y.LeftChild.Parent = y
		y.Color = node.Color
	}

	if !yOriginalColor {
		rbt.fixedAfterDeletion(x)
	}
}

func (rbt *RedBlackTree) transplant(u, v *Node) {
	if u.Parent == nil {
		rbt.Root = v
	} else if u == u.Parent.LeftChild {
		u.Parent.LeftChild = v
	} else {
		u.Parent.RightChild = v
	}

	if v != nil {
		v.Parent = u.Parent
	}
}

func (rbt *RedBlackTree) fixedAfterDeletion(node *Node) {
	root := rbt.Root
	var sibling *Node
	if node != nil {

		for node != root && !isRed(node) {
			if node == node.Parent.LeftChild {
				sibling = node.Parent.RightChild

				if isRed(sibling) {
					sibling.Color = false
					node.Parent.Color = true
					rbt.leftRotate(node.Parent)
					sibling = node.Parent.RightChild
				}

				if !isRed(sibling.LeftChild) && !isRed(sibling.RightChild) {
					sibling.Color = true
					node = node.Parent
				} else {
					if !isRed(sibling.RightChild) {
						sibling.LeftChild.Color = false
						sibling.Color = true
						rbt.rightRotate(sibling)
						sibling = node.Parent.RightChild
					}

					sibling.Color = node.Parent.Color
					node.Parent.Color = false
					sibling.RightChild.Color = false
					rbt.leftRotate(node.Parent)
					node = root
				}
			} else {
				sibling = node.Parent.LeftChild

				if isRed(sibling) {
					sibling.Color = false
					node.Parent.Color = true
					rbt.rightRotate(node.Parent)
					sibling = node.Parent.LeftChild
				}

				if !isRed(sibling.LeftChild) && !isRed(sibling.RightChild) {
					sibling.Color = true
					node = node.Parent
				} else {
					if !isRed(sibling.LeftChild) {
						sibling.RightChild.Color = false
						sibling.Color = true
						rbt.leftRotate(sibling)
						sibling = node.Parent.LeftChild
					}

					sibling.Color = node.Parent.Color
					node.Parent.Color = false
					sibling.LeftChild.Color = false
					rbt.rightRotate(node.Parent)
					node = root
				}
			}
		}
		node.Color = false
	}
}

func isRed(node *Node) bool {
	return node != nil && node.Color
}

func (rbt *RedBlackTree) leftRotate(x *Node) {
	y := x.RightChild
	x.RightChild = y.LeftChild
	if y.LeftChild != nil {
		y.LeftChild.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == nil {
		rbt.Root = y
	} else if x == x.Parent.LeftChild {
		x.Parent.LeftChild = y
	} else {
		x.Parent.RightChild = y
	}
	y.LeftChild = x
	x.Parent = y
}

func (rbt *RedBlackTree) rightRotate(y *Node) {
	x := y.LeftChild
	y.LeftChild = x.RightChild
	if x.RightChild != nil {
		x.RightChild.Parent = y
	}
	x.Parent = y.Parent
	if y.Parent == nil {
		rbt.Root = x
	} else if y == y.Parent.LeftChild {
		y.Parent.LeftChild = x
	} else {
		y.Parent.RightChild = x
	}
	x.RightChild = y
	y.Parent = x
}

func (rbt *RedBlackTree) Search(data float64) (bool, *Node) {
	currentNode := rbt.Root
	for currentNode != nil && currentNode.Data != data {
		if currentNode.LeftChild == nil && currentNode.RightChild == nil {
			return false, nil
		}
		if currentNode.Data > data {
			currentNode = currentNode.LeftChild
		} else if currentNode.Data <= data {
			currentNode = currentNode.RightChild
		}
	}
	if currentNode == nil {
		return false, nil
	}
	return true, currentNode
}

func (rbt *RedBlackTree) Minimum() float64 {
	currentNode := rbt.Root
	for currentNode.LeftChild != nil {
		currentNode = currentNode.LeftChild
	}
	return currentNode.Data
}

func (rbt *RedBlackTree) Maximum() float64 {
	currentNode := rbt.Root
	for currentNode.RightChild != nil {
		currentNode = currentNode.RightChild
	}
	return currentNode.Data
}

func (rbt *RedBlackTree) InOrderTraverse(captrue func(float64)) {

	var InOrder func(node *Node)
	InOrder = func(node *Node) {
		if node != nil {
			InOrder(node.LeftChild)
			captrue(node.Data)
			InOrder(node.RightChild)
		}
	}
	InOrder(rbt.Root)
}

func (rbt *RedBlackTree) PreOrderTraverse(captrue func(float64)) {

	var PreOrder func(node *Node)
	PreOrder = func(node *Node) {
		if node != nil {
			captrue(node.Data)
			PreOrder(node.LeftChild)
			PreOrder(node.RightChild)
		}
	}
	PreOrder(rbt.Root)
}

func (rbt *RedBlackTree) PostOrderTraverse(captrue func(float64)) {

	var PostOrder func(node *Node)
	PostOrder = func(node *Node) {
		if node != nil {
			PostOrder(node.LeftChild)
			PostOrder(node.RightChild)
			captrue(node.Data)
		}
	}
	PostOrder(rbt.Root)
}

func (rbt *RedBlackTree) Successor(data float64) float64 {
	_, node := rbt.Search(data)
	if node == nil {
		return 0.0
	}

	if node.RightChild != nil {
		currentNode := node.RightChild
		for currentNode.LeftChild != nil {
			currentNode = currentNode.LeftChild
		}
		return currentNode.Data
	}

	successor := 0.0
	parentNode := rbt.Root

	for parentNode != nil {
		if data < parentNode.Data {
			successor = parentNode.Data
			parentNode = parentNode.LeftChild
		} else if data >= parentNode.Data {
			parentNode = parentNode.RightChild
		}
	}
	return successor
}

func (rbt *RedBlackTree) Predecessor(data float64) float64 {
	_, node := rbt.Search(data)
	if node == nil {
		return 0.0
	}

	if node.RightChild != nil {
		currentNode := node.LeftChild
		for currentNode.RightChild != nil {
			currentNode = currentNode.RightChild
		}
		return currentNode.Data
	}

	predecessor := 0.0
	parentNode := rbt.Root

	for parentNode != nil {
		if data > parentNode.Data {
			predecessor = parentNode.Data
			parentNode = parentNode.RightChild
		} else if data <= parentNode.Data {
			parentNode = parentNode.LeftChild
		}
	}
	return predecessor
}
