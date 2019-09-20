package cmap

import (
	"sync"
)

type CMap struct {
	v map[string]interface{}
	l sync.RWMutex
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

func (m *CMap) Has(key string) bool {
	m.l.RLock()
	defer m.l.RUnlock()
	_, exists := m.v[key]
	return exists
}
