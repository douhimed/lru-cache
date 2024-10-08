package main

import (
	"fmt"
)

/*
*	Node & Linked List
 */

type Node struct {
	next  *Node
	prev  *Node
	value interface{}
}

type LinkeList struct {
	head *Node
	tail *Node
	size int
}

func NewLInkedList() *LinkeList {
	return &LinkeList{}
}

func (ll *LinkeList) Add(value interface{}) *Node {

	newNode := &Node{
		value: value,
		prev:  nil,
		next:  nil,
	}

	ll.size++

	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.prev = ll.tail
		ll.tail.next = newNode
		ll.tail = newNode
	}

	return newNode
}

func (ll *LinkeList) Print() string {

	if ll.IsEmpty() {
		return ""
	}

	var s string = fmt.Sprintf("%v", ll.head.value)

	current := ll.head.next
	for current != nil {
		s += fmt.Sprintf(" -> %v", current.value)
		current = current.next
	}

	return s
}

func (ll *LinkeList) IsEmpty() bool {
	return ll.head == nil
}

func (ll *LinkeList) FindIndex(v interface{}) int {
	if ll.IsEmpty() {
		return -1
	}

	current := ll.head
	currentIndex := 0

	for current != nil {
		if current.value == v {
			return currentIndex
		}
		current = current.next
		currentIndex++
	}

	return -1
}

func (ll *LinkeList) Find(v interface{}) (*Node, bool) {
	if ll.IsEmpty() {
		return nil, false
	}

	current := ll.head
	for current != nil {
		if current.value == v {
			return current, true
		}
		current = current.next
	}

	return nil, false
}

func (ll *LinkeList) Check(v interface{}) bool {

	node, ok := ll.Find(v)

	if ok {

		prev := node.prev
		next := node.next

		if prev != nil {
			prev.next = next
		}

		if next != nil {
			next.prev = prev
		}

		node.prev = nil
		node.next = ll.head
		ll.head.prev = node
		ll.head = node

		return false
	}

	ll.Add(v)
	return true
}

/*
* Hash
 */
type Hash map[string]*Node

/*
*	Cache
 */

const SIZE = 5

type Cache struct {
	Queue LinkeList
	Hash  Hash
	size  int
}

func NewCache() *Cache {
	return &Cache{
		Queue: *NewLInkedList(),
		Hash:  Hash{},
	}
}

func (c *Cache) Remove(n *Node) *Node {
	prev := n.prev
	next := n.next

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}

	c.size -= 1
	delete(c.Hash, fmt.Sprintf("%v", n.value))

	return n
}

func (c *Cache) Add(v interface{}) *Node {
	c.size += 1
	n := c.Queue.Add(v)
	c.Hash[fmt.Sprintf("%v", v)] = n

	if c.size > SIZE {
		c.Remove(c.Queue.tail)
	}

	return n
}

func (c *Cache) Check(value interface{}) {
	if val, ok := c.Hash[fmt.Sprintf("%v", value)]; ok {
		c.Remove(val)
	}
	c.Hash[fmt.Sprintf("%v", value)] = c.Add(value)
}

func (c *Cache) Print() string {
	if c.Queue.IsEmpty() {
		return ""
	}

	linkedList := c.Queue
	return linkedList.Print()
}
