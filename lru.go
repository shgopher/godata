package godata

import "container/list"

type LRUCache struct {
	Cache map[int]*list.Element
	Cap   int
	List  *list.List
}

type LRUNode struct {
	Key, Value int
}
//new a LRU
func NewLRU(capacity int) LRUCache {
	return LRUCache{
		Cache: make(map[int]*list.Element),
		Cap:   capacity,
		List:  list.New(),
	}
}
// Rebuild the LRU.
func(l *LRUCache)ReBuild(){
	l.List.Init()
	l.Cache = make(map[int]*list.Element)
}

// exist:value, no:-1
func (l *LRUCache) Get(key int) int {
	v, ok := l.Cache[key]
	if ok {
		l.List.MoveToFront(v)
		return v.Value.(*LRUNode).Value
	}
	// top the LRUNode
	return -1
}
// put a new data
func (l *LRUCache) Put(key int, value int) {
	if l.List.Len() < l.Cap {
		if v, ok := l.Cache[key]; ok {
			// if has key and match value.
			if v.Value.(*LRUNode).Value != value {
				v.Value.(*LRUNode).Value  = value
			}
			l.List.MoveToFront(v)
			// the key is not match value

		} else {
			n := &LRUNode{
				Key:   key,
				Value: value,
			}
			el := l.List.PushFront(n)
			l.Cache[key] = el
		}
	} else {
		if v, ok := l.Cache[key]; ok {
			// if has key and match value.
			if v.Value.(*LRUNode).Value != value {
				v.Value.(*LRUNode).Value  = value
			}
			l.List.MoveToFront(v)
		} else {
			tti := l.List.Back()
			delete(l.Cache, tti.Value.(*LRUNode).Key)
			l.List.Remove(tti)
			n := &LRUNode{
				Key:   key,
				Value: value,
			}
			el := l.List.PushFront(n)
			l.Cache[key] = el
		}

	}
}