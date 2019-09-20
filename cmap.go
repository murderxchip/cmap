package cmap

import (
	"sync"
)

type CMap struct {
	v map[string]interface{}
	l sync.RWMutex
}

type MapItem struct {
	Key   string
	Value interface{}
}

func NewCMap() *CMap {
	return &CMap{v: make(map[string]interface{})}
}

func (m *CMap) Size() int {
	return len(m.v)
}

func (m *CMap) Set(key string, value interface{}) {
	m.l.Lock()
	defer m.l.Unlock()
	m.v[key] = value
}

func (m *CMap) Get(key string) (value interface{}, exists bool) {
	m.l.RLock()
	defer m.l.RUnlock()
	value, exists = m.v[key]
	return
}

func (m *CMap) Exist(key string) bool {
	m.l.RLock()
	defer m.l.RUnlock()
	_, exists := m.v[key]
	return exists
}

func (m *CMap) Dump() <-chan MapItem {
	outChan := make(chan MapItem, m.Size())
	for k, v := range m.v {
		outChan <- MapItem{
			Key:   k,
			Value: v,
		}
	}
	close(outChan)
	return outChan
}
