package leetcode

import "fmt"

type LRUCache struct {
	cache    map[int]*Node
	head     *Node
	tail     *Node
	capacity int
}

type Node struct {
	prev  *Node
	next  *Node
	value int
	key   int
}

// func Constructor(capacity int) LRUCache {
// 	return LRUCache{cache: make(map[int]*Node, capacity),
// 		head:     nil,
// 		tail:     nil,
// 		capacity: capacity,
// 	}
// }

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		// fmt.Printf("Get Key %d, Value %d\n", key, v.value)
		this.top(v, false)
		return v.value
	}
	// fmt.Printf("Get Key %d, Value %d\n", key, -1)
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.cache[key]
	if ok {
		// if contains key, top it
		v.value = value
		this.top(v, false)
	} else {
		// fmt.Printf("len(this.cache): %d, this.capacity : %d\n", len(this.cache), this.capacity)
		if len(this.cache) == this.capacity {
			this.remove()
		}
		nn := &Node{prev: nil, next: nil, value: value, key: key}
		this.top(nn, true)
		this.cache[key] = nn
	}
}

func (this *LRUCache) remove() {
	if this.tail != nil {
		delete(this.cache, this.tail.key)
		// fmt.Printf("deleteing %d\n", this.tail.key)
		this.tail = this.tail.prev
		if this.tail != nil {
			this.tail.next = nil
		}
	}
}

func (this *LRUCache) top(n *Node, newNode bool) {
	if len(this.cache) == 0 {
		// new node
		this.head = n
		this.tail = n
	} else if newNode {
		n.next = this.head
		this.head.prev = n
		this.head = n
	} else if n != this.head {
		n.prev.next = n.next
		if n.next != nil {
			n.next.prev = n.prev
		} else {
			this.tail = n.prev
		}
		n.prev = nil
		n.next = this.head
		this.head.prev = n
		this.head = n
	}
	// this.Print()
}

func (this *LRUCache) Print() {
	fmt.Printf("head:%d, tail:%d,,,", this.head.key, this.tail.key)
	for p := this.head; p != nil; p = p.next {
		fmt.Printf("-> %d", p.value)
	}
	fmt.Println()
}
