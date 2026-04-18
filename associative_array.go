package main

import "sync"

type AssociativeArray struct {
	data map[string]interface{}
	lock sync.RWMutex
}

func NewAssociativeArray() *AssociativeArray {
	return &AssociativeArray{
		data: make(map[string]interface{}),
	}
}

func (a *AssociativeArray) Get(key string) (interface{}, bool) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	value, exists := a.data[key]
	return value, exists
}

func (a *AssociativeArray) Set(key string, value interface{}) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.data[key] = value
}

func (a *AssociativeArray) Delete(key string) {
	a.lock.Lock()
	defer a.lock.Unlock()
	delete(a.data, key)
}

func (a *AssociativeArray) Keys() []string {
	a.lock.RLock()
	defer a.lock.RUnlock()
	keys := make([]string, 0, len(a.data))
	for key := range a.data {
		keys = append(keys, key)
	}
	return keys
}

func (a *AssociativeArray) Contains(key string) bool {
	a.lock.RLock()
	defer a.lock.RUnlock()
	_, exists := a.data[key]
	return exists
}