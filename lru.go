package learn

import "sync"

type listNode struct {
	key  interface{}
	val  interface{}
	pre  *listNode
	next *listNode
}

type Lru struct {
	node  listNode
	items map[interface{}]*listNode
	mu    sync.RWMutex
	size  int
}

func (lru *Lru) remove(node *listNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
	node.pre = nil
	node.next = nil
}

func (lru *Lru) addToHead(node *listNode) {
	node.next = lru.node.next
	lru.node.next.pre = node
	lru.node.next = node
	node.pre = &lru.node
}

func LRU(size int) *Lru {
	lru := &Lru{}
	lru.node.pre = &lru.node
	lru.node.next = &lru.node
	lru.items = make(map[interface{}]*listNode, size)
	lru.size = size
	return lru
}

func (lru *Lru) Set(key interface{}, val interface{}) {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	if node, ok := lru.items[key]; ok {
		lru.remove(node)
		lru.addToHead(node)
		return
	}

	if len(lru.items) >= lru.size {
		node := lru.node.pre
		lru.remove(node)
		delete(lru.items, node.key)
	}

	lru.items[key] = &listNode{key: key, val: val}
	lru.addToHead(lru.items[key])
}

func (lru *Lru) Get(key interface{}) interface{} {
	lru.mu.Lock()
	defer lru.mu.Unlock()

	if node, ok := lru.items[key]; ok {
		lru.remove(node)
		lru.addToHead(node)
		return node.val
	}

	return nil
}

func (lru *Lru) Len() int {
	lru.mu.RLock()
	defer lru.mu.RUnlock()
	return len(lru.items)
}

func (lru *Lru) GetAll() map[interface{}]interface{} {
	lru.mu.RLock()
	defer lru.mu.RUnlock()

	all := make(map[interface{}]interface{}, len(lru.items))

	for k, v := range lru.items {
		all[k] = v.val
	}

	return all
}

func (lru *Lru) GetLinkKey() []interface{} {
	lru.mu.RLock()
	defer lru.mu.RUnlock()

	l := make([]interface{}, 0, len(lru.items))

	cur := lru.node.next

	for cur != &lru.node {
		l = append(l, cur.key)
		cur = cur.next
	}

	return l
}
