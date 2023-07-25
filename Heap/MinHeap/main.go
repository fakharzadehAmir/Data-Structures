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
		if reHeap.LeftChild != nil && reHeap.Data < reHeap.LeftChild.Data {
			break
		}
		if reHeap.RightChild != nil && reHeap.Data < reHeap.RightChild.Data {
			break
		}
		if reHeap.Data >= reHeap.LeftChild.Data {
			temp := reHeap.LeftChild.Data
			reHeap.LeftChild.Data = reHeap.Data
			reHeap.Data = temp
			reHeap = reHeap.LeftChild
		} else if reHeap.Data >= reHeap.RightChild.Data {
			temp := reHeap.RightChild.Data
			reHeap.RightChild.Data = reHeap.Data
			reHeap.Data = temp
			reHeap = reHeap.RightChild
		}
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

	InsertMinHeapify := func(newValue *MinHeap) {
		for newValue.Parent != nil {
			if newValue.Data < newValue.Parent.Data {
				temp := newValue.Parent.Data
				newValue.Parent.Data = newValue.Data
				newValue.Data = temp
				newValue = newValue.Parent
			} else {
				break
			}
		}
	}
	InsertMinHeapify(newValue)
}

func (mh *MinHeap) Extract() (*MinHeap, float64) {
	minNode := mh.Data

	var findLastNode []*MinHeap
	findLastNode = append(findLastNode, mh)
	var lastNode *MinHeap
	for len(findLastNode) > 0 {
		lastNode = findLastNode[0]
		findLastNode = findLastNode[1:]
		if lastNode.LeftChild != nil {
			findLastNode = append(findLastNode, lastNode.LeftChild)
		}
		if lastNode.RightChild != nil {
			findLastNode = append(findLastNode, lastNode.RightChild)
		}
	}

	if mh == lastNode {
		return nil, minNode
	} else {
		mh.Data = lastNode.Data
		parentNode := lastNode.Parent
		if lastNode == parentNode.LeftChild {
			parentNode.LeftChild = nil
		}
		if lastNode == parentNode.RightChild {
			parentNode.RightChild = nil
		}

		mh.MinHeapify()
		return mh, minNode
	}

}

func (mh *MinHeap) Peek() float64 {
	return mh.Data
}
