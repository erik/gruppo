package util

import (
	"strings"
	"sync"
	"time"
)

type JSONTime struct{ Time time.Time }

func (j JSONTime) MarshalJSON() ([]byte, error) {
	t := j.Time.Format(time.RFC3339)
	return []byte("\"" + t + "\""), nil
}

func (j *JSONTime) UnmarshalJSON(p []byte) error {
	t, err := time.Parse(
		time.RFC3339,
		strings.Replace(string(p), "\"", "", -1))

	if err != nil {
		return err
	}

	*j = JSONTime{t}

	return nil
}

type QueueItem struct {
	K string
	V interface{}
}

type UniqueQueue struct {
	ch    chan QueueItem
	m     sync.Mutex
	items map[string]bool
}

func NewUniqueQueue(maxLen int) UniqueQueue {
	return UniqueQueue{
		ch:    make(chan QueueItem, maxLen),
		items: make(map[string]bool),
	}
}

func (q *UniqueQueue) Push(item QueueItem) {
	q.m.Lock()
	defer q.m.Unlock()

	// Nothing to do if item is already in queue
	if _, ok := q.items[item.K]; ok {
		return
	}

	q.items[item.K] = true
	q.ch <- item
}

func (q *UniqueQueue) Pop() interface{} {
	item := <-q.ch

	q.m.Lock()
	defer q.m.Unlock()

	delete(q.items, item.K)
	return item.V
}
