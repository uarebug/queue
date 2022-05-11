package queue

import (
	"container/heap"
	"fmt"
)

type Queue interface {
	Enqueue(i Item) error

	Dequeue() (i Item)
}

type queue struct {
	items     []Item
	c         chan Item
	maxLength int
}

func NewQueue(maxQueueLength int) Queue {
	return &queue{
		c:         make(chan Item, 1),
		maxLength: maxQueueLength,
	}
}

// implement heap.Interface

func (q queue) Len() int {
	if q.items == nil {
		return 0
	}

	return len(q.items)
}

func (q queue) Swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
}

func (q queue) Less(i, j int) bool {
	return q.items[i].Value() >= q.items[j].Value()
}

func (q *queue) Push(h interface{}) {
	q.items = append(q.items, h.(Item))
}

func (q *queue) Pop() (x interface{}) {
	n := len(q.items)
	x = (q.items)[n-1]
	q.items = q.items[:n-1]

	return x
}

// implement Queue interface

// Enqueue enqueue
func (q *queue) Enqueue(i Item) error {
	if q.Len() >= q.maxLength {
		return fmt.Errorf("队列已满，请稍后再试")
	}

	if q.Len() == 0 && len(q.c) == 0 {
		q.c <- i
		return nil
	}

	heap.Push(q, i)
	return nil
}

// Dequeue dequeue
func (q *queue) Dequeue() (i Item) {
	if q.Len() > 0 {
		return heap.Pop(q).(Item)
	}

	return <-q.c
}
