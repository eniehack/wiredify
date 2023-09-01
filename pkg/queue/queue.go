package queue

type Queue struct {
	data []rune
}

type Queuer interface {
	isEmpty() bool
	isFull() bool
	Enqueue(rune)
	Dequeue() rune
	Seek() rune
}

func New(d []rune) *Queue {
	return &Queue{
		data: d,
	}
}

func (q *Queue) Len() int {
	return len(q.data)
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) IsFull() bool {
	return false
}

func (q *Queue) Enqueue(elem rune) {
	q.data = append(q.data, elem)
	//q.lastElement += 1
}

func (q *Queue) Dequeue() rune {
	elem := q.data[0]
	if len(q.data[1:]) == 0 {
		q.data = []rune{}
		return elem
	}
	q.data = q.data[1:]
	//q.lastElement -= 1
	return elem
}

func (q *Queue) DequeueMany(length uint) []rune {
	var i uint = 0
	arr := []rune{}
	for ; i < length; i++ {
		if q.IsEmpty() {
			break
		}
		arr[i] = q.Dequeue()
	}
	return arr
}

func (q *Queue) Seek() rune {
	return q.data[0]
}

func (q *Queue) Data() []rune {
	return q.data
}
