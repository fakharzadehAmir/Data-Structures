package maxheap

type MaxHeap struct {
	Parent     *MaxHeap
	Data       float64
	LeftChild  *MaxHeap
	RightChild *MaxHeap
}

func NewMaxHeap(initialData float64) *MaxHeap {
	return &MaxHeap{
		Data: initialData,
	}
}

func (mh *MaxHeap) MaxHeapify() {
	reHeap := mh
	for {
		if reHeap.LeftChild == nil && reHeap.RightChild == nil {
			break
		}
		largest := reHeap
		if reHeap.LeftChild != nil && reHeap.LeftChild.Data > largest.Data {
			largest = reHeap.LeftChild
		}
		if reHeap.RightChild != nil && reHeap.RightChild.Data > largest.Data {
			largest = reHeap.RightChild
		}
		if largest == reHeap {
			break
		}
		reHeap.Data, largest.Data = largest.Data, reHeap.Data
		reHeap = largest
	}
}

func (mh *MaxHeap) Insert(newData float64) {
	newValue := &MaxHeap{
		Data: newData,
	}
	if mh == nil {
		mh = newValue
		return
	}

	var checkChild []*MaxHeap
	checkChild = append(checkChild, mh)
	for len(checkChild) > 0 {
		currentNode := checkChild[0]
		checkChild = checkChild[1:]
		if currentNode.LeftChild == nil {
			currentNode.LeftChild = newValue
			newValue.Parent = currentNode
			break
		} else {
			checkChild = append(checkChild, currentNode.LeftChild)
		}

		if currentNode.RightChild == nil {
			currentNode.RightChild = newValue
			newValue.Parent = currentNode
			break
		} else {
			checkChild = append(checkChild, currentNode.RightChild)
		}
	}

	for newValue.Parent != nil {
		if newValue.Data > newValue.Parent.Data {
			newValue.Data, newValue.Parent.Data = newValue.Parent.Data, newValue.Data
			newValue = newValue.Parent
		} else {
			break
		}
	}
}

func (mh *MaxHeap) Extract() (*MaxHeap, float64) {
	maxNode := mh.Data

	lastNode := getLastNode(mh)
	if mh != lastNode {
		mh.Data = lastNode.Data
		parentNode := lastNode.Parent
		if lastNode == parentNode.LeftChild {
			parentNode.LeftChild = nil
		} else {
			parentNode.RightChild = nil
		}
		mh.MaxHeapify()
	}

	return mh, maxNode
}

func (mh *MaxHeap) Peek() float64 {
	return mh.Data
}

func (mh *MaxHeap) Delete(data float64) {
	nodeToDelete := findNodeToDelete(mh, data)

	if nodeToDelete == nil {
		return
	}

	lastNode := getLastNode(mh)
	if nodeToDelete != lastNode {
		nodeToDelete.Data = lastNode.Data
		parentNode := lastNode.Parent
		if lastNode == parentNode.LeftChild {
			parentNode.LeftChild = nil
		} else {
			parentNode.RightChild = nil
		}
		mh.MaxHeapify()
	} else {
		parentNode := lastNode.Parent
		if lastNode == parentNode.LeftChild {
			parentNode.LeftChild = nil
		} else {
			parentNode.RightChild = nil
		}
	}
}

func findNodeToDelete(mh *MaxHeap, data float64) *MaxHeap {
	var nodesToVisit []*MaxHeap
	nodesToVisit = append(nodesToVisit, mh)

	for len(nodesToVisit) > 0 {
		currentNode := nodesToVisit[0]
		nodesToVisit = nodesToVisit[1:]

		if currentNode.Data == data {
			return currentNode
		}

		if currentNode.LeftChild != nil {
			nodesToVisit = append(nodesToVisit, currentNode.LeftChild)
		}

		if currentNode.RightChild != nil {
			nodesToVisit = append(nodesToVisit, currentNode.RightChild)
		}
	}

	return nil
}

func getLastNode(mh *MaxHeap) *MaxHeap {
	var nodesToVisit []*MaxHeap
	nodesToVisit = append(nodesToVisit, mh)

	var lastNode *MaxHeap

	for len(nodesToVisit) > 0 {
		lastNode = nodesToVisit[0]
		nodesToVisit = nodesToVisit[1:]

		if lastNode.LeftChild != nil {
			nodesToVisit = append(nodesToVisit, lastNode.LeftChild)
		}

		if lastNode.RightChild != nil {
			nodesToVisit = append(nodesToVisit, lastNode.RightChild)
		}
	}

	return lastNode
}
