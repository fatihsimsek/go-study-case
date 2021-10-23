package concurrentmap

import "sync"

type ConcurrentMap struct {
	items map[interface{}]interface{}
	mu    sync.RWMutex
}

func New() *ConcurrentMap {
	return &ConcurrentMap{items: make(map[interface{}]interface{})}
}

func (m *ConcurrentMap) Get(key interface{}) (value interface{}, found bool) {
	m.mu.RLock()
	value, found = m.items[key]
	m.mu.RUnlock()
	return
}

func (m *ConcurrentMap) Contains(key interface{}) bool {
	_, found := m.Get(key)
	return found
}

func (m *ConcurrentMap) Put(key interface{}, value interface{}) {
	m.mu.Lock()
	m.items[key] = value
	m.mu.Unlock()
}

func (m *ConcurrentMap) Remove(key interface{}) (found bool) {
	_, found = m.Get(key)
	if found {
		m.mu.Lock()
		delete(m.items, key)
		defer m.mu.Unlock()
	}
	return
}

func (m *ConcurrentMap) CopyTo() map[interface{}]interface{} {
	newMap := make(map[interface{}]interface{})
	m.mu.RLock()
	for k, v := range m.items {
		newMap[k] = v
	}
	m.mu.RUnlock()
	return newMap
}

func (m *ConcurrentMap) Size() (size int) {
	m.mu.RLock()
	size = len(m.items)
	m.mu.RUnlock()
	return
}
