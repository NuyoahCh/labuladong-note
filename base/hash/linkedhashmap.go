package main

import "fmt"

type MyLinkedHashMap struct {
	head *Node
	tail *Node
	m    map[string]*Node
}

type Node struct {
	key  string
	val  int
	next *Node
	prev *Node
}

func Constructor1() MyLinkedHashMap {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return MyLinkedHashMap{
		head: head,
		tail: tail,
		m:    make(map[string]*Node),
	}
}

func (this *MyLinkedHashMap) Get(key string) int {
	if node, ok := this.m[key]; ok {
		return node.val
	}
	return -1
}

func (this *MyLinkedHashMap) Put(key string, val int) {
	if _, ok := this.m[key]; !ok {
		node := &Node{key: key, val: val}
		this.m[key] = node
		return
	}
	this.m[key].val = val
}

func (this *MyLinkedHashMap) Remove(key string) {
	if _, ok := this.m[key]; !ok {
		return
	}
	node := this.m[key]
	delete(this.m, key)
	this.removeNode(node)
}

func (this *MyLinkedHashMap) ContainsKey(key string) bool {
	_, ok := this.m[key]
	return ok
}

func (this *MyLinkedHashMap) Keys1() []string {
	keyList := make([]string, 0)
	for p := this.head.next; p != this.tail; p = p.next {
		keyList = append(keyList, p.key)
	}
	return keyList
}

func (this *MyLinkedHashMap) addLastNode(x *Node) {
	temp := this.tail.prev
	// temp <-> tail

	x.next = this.tail
	x.prev = temp
	// temp <-> x <-> tail

	temp.next = x
	this.tail.prev = x
	// temp <-> x <-> tail
}

func (this *MyLinkedHashMap) removeNode(x *Node) {
	prev := x.prev
	next := x.next
	// prev <-> x <-> next

	prev.next = next
	next.prev = prev

	x.next = nil
	x.prev = nil
}

func main() {
	myMap := Constructor1()
	myMap.Put("a", 1)
	myMap.Put("b", 2)
	myMap.Put("c", 3)
	myMap.Put("d", 4)
	myMap.Put("e", 5)

	// output: a b c d e
	fmt.Println(myMap.Keys1())

	myMap.Remove("c")

	// output: a b d e
	fmt.Println(myMap.Keys1())
}