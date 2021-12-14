package types

type MatrixElementQueue struct {
	elements []MatrixElement
}

func NewMatrixElementQueue() *MatrixElementQueue {
	return &MatrixElementQueue{}
}

func (q *MatrixElementQueue) PushMultiple(elements []MatrixElement) {
	for _, e := range elements {
		q.Push(e)
	}
}

func (q *MatrixElementQueue) Push(e MatrixElement) {
	q.elements = append(q.elements, e)
}

func (q *MatrixElementQueue) Pop() MatrixElement {
	e := q.elements[0]

	if len(q.elements) == 1 {
		q.elements = []MatrixElement{}
	} else {
		q.elements = q.elements[1:]
	}

	return e
}

func (q *MatrixElementQueue) Size() int {
	return len(q.elements)
}
