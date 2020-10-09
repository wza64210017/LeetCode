package main

import "fmt"

// 双向链表节点
type LRUNode struct {
	key, value int
	prev, next *LRUNode
}

type LRUCache struct {
	m    map[int]*LRUNode // 节点map
	len  int              // 长度
	cap  int              // 容量
	head *LRUNode         // 头指针
	tail *LRUNode         // 尾指针
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		m:   make(map[int]*LRUNode, capacity),
		len: 0,
		cap: capacity,
	}

	lru.head = &LRUNode{}
	lru.tail = &LRUNode{}

	lru.head.next = lru.tail
	lru.head.prev = nil

	lru.tail.prev = lru.head
	lru.tail.next = nil

	return lru
}

func (this *LRUCache) Get(key int) int {
	// key不存在
	node, exists := this.m[key]
	if !exists {
		fmt.Println("key:", key, "value:", -1)
		return -1
	}

	// 移至头节点
	this.RemoveNode(node)
	this.AddHead(node)
	fmt.Println("key:", key, "value:", node.value)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, exists := this.m[key]
	// key存在,移动至头部
	if exists {
		node.value = value
		this.m[key] = node
		this.RemoveNode(node)
		this.AddHead(node)
		return
	}

	// 容量已满，淘汰末尾
	if this.len >= this.cap {
		delete(this.m, this.tail.prev.key)
		this.RemoveNode(this.tail.prev)
		this.len--
	}

	node = &LRUNode{
		key:   key,
		value: value,
	}

	this.AddHead(node)
	this.m[key] = node
	this.len++
}

// 添加至头节点
func (this *LRUCache) AddHead(node *LRUNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

// 移除节点
func (this *LRUCache) RemoveNode(node *LRUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev

	node.prev = nil
	node.next = nil
}

func main() {
	lru := Constructor(2)
	lru.Put(2, 1)
	lru.Put(2, 2)
	lru.Get(2)
	lru.Put(1, 1)
	lru.Put(4, 1)
	lru.Get(2)
	//lru.Put(4, 4)
	//lru.Get(1)
	//lru.Get(3)
	//lru.Get(4)
}
