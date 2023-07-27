package queue

type Queue struct {
	Items []interface{}
	Size  int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(newData interface{}) {
	q.Items = append(q.Items, newData)
	q.Size++
}

func (q *Queue) Dequeue() interface{} {
	if q.Size > 0 {
		deqeueuedItem := q.Items[0]
		q.Items = q.Items[1:]
		q.Size--
		return deqeueuedItem
	}
	return nil
}

func (q *Queue) FrontValue() interface{} {
	if q.Size > 0 {
		return q.Items[0]
	}
	return nil
}

func (q *Queue) RearValue() interface{} {
	if q.Size > 0 {
		return q.Items[q.Size-1]
	}
	return nil
}

func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}

func (q *Queue) Count() int {
	return q.Size
}

func (q *Queue) Clear() {
	q.Items = nil
	q.Size = 0
}

func (q *Queue) Contains(data interface{}) bool {
	for _, value := range q.Items {
		if value == data {
			return true
		}
	}
	return false
}

func (q *Queue) IndexOf(data interface{}) int {
	for index, value := range q.Items {
		if value == data {
			return index
		}
	}
	return -1
}

func (q *Queue) LastIndexOf(data interface{}) int {
	for index := q.Size - 1; index >= 0; index-- {
		if q.Items[index] == data {
			return index
		}
	}
	return -1
}
