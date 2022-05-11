package main

import (
	"fmt"
	"github.com/uarebug/queue"
	"log"
	"sync"
)

type PriorityItem struct {
	Name     string
	Priority int
}

func (p PriorityItem) Value() int {
	return p.Priority
}

func main() {
	wg := sync.WaitGroup{}
	q := queue.NewQueue(100)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		t := &PriorityItem{
			Name:     fmt.Sprintf("item-%d", i),
			Priority: i,
		}
		err := q.Enqueue(t)
		if err != nil {
			log.Println("enqueue error", err.Error())
			continue
		}
		log.Println("enqueue", t)
	}

	go func() {
		for {
			t := q.Dequeue().(*PriorityItem)
			log.Println("dequeue: ", t)
			wg.Done()
		}
	}()

	wg.Wait()
}
