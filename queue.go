package queue

import (
	"container/heap"
	"errors"
	"sort"
	"sync"
)

type Queue interface {
	Enqueue(i Item) error

	Dequeue() Item

	Length() int

	Peek() Item

	Range(f func(key, value interface{}) bool)
}

type queue struct {
	items     []Item
	c         chan Item
	maxLength int
	lock      sync.Mutex
}

func NewQueue(maxQueueLength int) Queue {
	return &queue{
		c:         make(chan Item, 1),
		maxLength: maxQueueLength,
	}
}

// implement heap.Interface

func (q *queue) Len() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.items == nil {
		return 0
	}

	return len(q.items)
}

func (q *queue) Swap(i, j int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

func (q *queue) Less(i, j int) bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.items[i].Value() >= q.items[j].Value()
}

func (q *queue) Push(h interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, h.(Item))
}

func (q *queue) Pop() (x interface{}) {
	n := q.Len()
	q.lock.Lock()
	defer q.lock.Unlock()

	x = (q.items)[n-1]
	q.items = q.items[:n-1]

	return x
}

// implement Queue interface

// Enqueue enqueue
func (q *queue) Enqueue(i Item) error {
	if q.Len() >= q.maxLength {
		return errors.New("the queue is full")
	}

	if q.Len() == 0 && len(q.c) == 0 {
		q.c <- i
		return nil
	}

	heap.Push(q, i)
	return nil
}

// Dequeue dequeue
func (q *queue) Dequeue() Item {
	if q.Len() > 0 {
		return heap.Pop(q).(Item)
	}

	return <-q.c
}

func (q *queue) Length() int {
	return q.Len() + len(q.c)
}

func (q *queue) Peek() Item {
	if q.Len() > 0 {
		return q.items[0]
	}

	return <-q.c
}

func (q *queue) Range(f func(key, value interface{}) bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	items := make([]Item, len(q.items))
	copy(items, q.items)

	sort.Slice(items, func(i, j int) bool {
		return items[i].Value() > items[j].Value()
	})

	for k, v := range items {
		if !f(k, v) {
			break
		}
	}
}
