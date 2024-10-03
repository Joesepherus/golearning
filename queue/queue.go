package queue

import "sync"

type Queue struct {
	items []any
	sync.Mutex
}

func (q *Queue) Enqueue(item any) {
	q.Lock()
	defer q.Unlock()
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (any, bool) {
	q.Lock()
	defer q.Unlock()
	if len(q.items) == 0 {
		return nil, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue) IsEmpty() bool {
	q.Lock()
	defer q.Unlock()
	return len(q.items) == 0
}
