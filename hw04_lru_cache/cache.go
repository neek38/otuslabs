package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key string, value interface{}) bool // Добавить значение в кэш по ключу
	Get(key string) (interface{}, bool)     // Получить значение из кэша по ключу
	Clear()                                 // Очистить кэш
}

type lruCache struct {
	capacity int
	queue    *list
	items    map[string]*listItem
}

func (l *lruCache) Set(key string, value interface{}) bool {
	if element, exists := l.items[key]; exists {
		l.queue.MoveToFront(element)
		element.Value = &cacheItem{key: key, Value: value}
		return true
	}

	if l.queue.Len() == l.capacity {
		if element := l.queue.Back(); element != nil {
			l.queue.Remove(element)
			delete(l.items, element.Value.(*cacheItem).key)
		}
	}

	item := &cacheItem{key: key, Value: value}

	l.queue.PushFront(item)
	l.items[key] = l.queue.Front()
	return false
}

func (l *lruCache) Get(key string) (interface{}, bool) {
	if elem, exists := l.items[key]; exists {
		l.queue.MoveToFront(elem)
		return elem.Value.(*cacheItem).Value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.queue = &list{}
	l.items = make(map[string]*listItem, l.capacity)
}

type cacheItem struct {
	key   string
	Value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{capacity: capacity,
		queue: &list{},
		items: make(map[string]*listItem, capacity)}
}
