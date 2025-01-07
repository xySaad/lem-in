package path_finder

type queued struct {
	room, parrent string
}

type dequeue[T any] struct {
	slice []T
}

func NewQueue[T any](slice []T) *dequeue[T] {
	return &dequeue[T]{
		slice: slice,
	}
}

func (q *dequeue[T]) Len() int {
	if q == nil {
		return 0
	}
	return len(q.slice)
}

func (q *dequeue[T]) Push(x T) {
	if q == nil {
		return
	}
	q.slice = append(q.slice, x)
}
func (q *dequeue[T]) Append(x T) {
	if q == nil {
		return
	}
	q.slice = append([]T{x}, q.slice...)
}
func (q *dequeue[T]) Shift() *T {
	bottom := &q.slice[0]
	q.slice = q.slice[1:]
	return bottom
}

func (q *dequeue[T]) Pop() (*T, bool) {
	if q == nil || len(q.slice) == 0 {
		return nil, false
	}
	top := &q.slice[len(q.slice)-1]
	q.slice = q.slice[:len(q.slice)-1]
	return top, true
}

func (q *dequeue[T]) Peek() *T {
	if q == nil || len(q.slice) == 0 {
		return nil
	}
	return &q.slice[len(q.slice)-1]
}

func (q *dequeue[T]) Bottom() *T {
	if len(q.slice) == 0 {
		return nil
	}

	return &q.slice[0]
}

// TODO: code range method
func (q *dequeue[T]) All() *[]T {
	return &q.slice
}

func (q *dequeue[T]) removeAt(index int) *T {
	item := &q.slice[index]
	q.slice = append(q.slice[:index], q.slice[index+1:]...)
	return item
}
