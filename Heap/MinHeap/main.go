package minheap

type MinHeap struct {
	Parent     *MinHeap
	Data       float64
	LeftChild  *MinHeap
	RightChild *MinHeap
}

func NewMinHeap(initialData float64) *MinHeap {
	return &MinHeap{
		Data: initialData,
	}
}

func (mh *MinHeap) MinHeapify() {
	reHeap := mh
	for {
		if reHeap.LeftChild == nil && reHeap.RightChild == nil {
			break
		}
		smallest := reHeap
		if reHeap.LeftChild != nil && reHeap.LeftChild.Data < smallest.Data {
			smallest = reHeap.LeftChild
		}
		if reHeap.RightChild != nil && reHeap.RightChild.Data < smallest.Data {
			smallest = reHeap.RightChild
		}
		if smallest == reHeap {
			break
		}
		reHeap.Data, smallest.Data = smallest.Data, reHeap.Data
		reHeap = smallest
	}
}

func (mh *MinHeap) Insert(newData float64) {
	newValue := &MinHeap{
		Data: newData,
	}
	if mh == nil {
		mh = newValue
		return
	}

	var checkChild []*MinHeap
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
		if newValue.Data < newValue.Parent.Data {
			newValue.Data, newValue.Parent.Data = newValue.Parent.Data, newValue.Data
			newValue = newValue.Parent
		} else {
			break
		}
	}
}

func (mh *MinHeap) Extract() (*MinHeap, float64) {
	minNode := mh.Data

	lastNode := getLastNode(mh)
	if mh != lastNode {
		mh.Data = lastNode.Data
		parentNode := lastNode.Parent
		if lastNode == parentNode.LeftChild {
			parentNode.LeftChild = nil
		} else {
			parentNode.RightChild = nil
		}
		mh.MinHeapify()
	}

	return mh, minNode
}

func (mh *MinHeap) Peek() float64 {
	return mh.Data
}

func (mh *MinHeap) Delete(data float64) {
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
		mh.MinHeapify()
	}
}

func findNodeToDelete(mh *MinHeap, data float64) *MinHeap {
	var nodesToVisit []*MinHeap
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

func getLastNode(mh *MinHeap) *MinHeap {
	var nodesToVisit []*MinHeap
	nodesToVisit = append(nodesToVisit, mh)

	var lastNode *MinHeap

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
