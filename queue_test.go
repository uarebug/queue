package queue

import (
	"fmt"
	"math/rand"
	"testing"
)

type TestItem struct {
	name  string
	value int
}

func (t *TestItem) Value() int {
	return t.value
}

func TestQueue_Length(t *testing.T) {
	q := NewQueue(1000)
	for i := 0; i < 100; i++ {
		err := q.Enqueue(&TestItem{name: fmt.Sprintf("test%d", i), value: rand.Intn(100)})
		if err != nil {
			fmt.Println(err)
			t.FailNow()
		}
	}

	if q.Length() != 100 {
		t.Fail()
	}
}
