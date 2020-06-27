package lru

import (
	"github.com/a1ekaeyVorobyev/otus_golang_hw/hw4/list"
)

type Item struct {
	Key   string
	Value interface{}
}

type LRU struct {
	capacity int
	items    map[string]*list.Item
	queue    *List
}

func NewLru(capacity int) *LRU {
	return &LRU{
		capacity: capacity,
		items:    make(map[string]*list.Element),
		queue:    list.New(),
	}
}
