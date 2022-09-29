package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.Mutex
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	listItem, ok := l.items[key]
	newCacheItem := cacheItem{key, value}
	if ok {
		listItem.Value = newCacheItem
		l.queue.MoveToFront(listItem)
		return true
	}
	l.queue.PushFront(newCacheItem)
	l.items[key] = l.queue.Front()
	if l.queue.Len() > l.capacity {
		tailListItem := l.queue.Back()
		l.queue.Remove(tailListItem)
		delete(l.items, tailListItem.Value.(cacheItem).key)
	}
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	item, ok := l.items[key]
	if !ok {
		return nil, false
	}
	l.queue.MoveToFront(item)
	return item.Value.(cacheItem).value, ok
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
