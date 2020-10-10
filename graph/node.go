package main

import (
	"fmt"
	"sync"
)

//
// Node list
//
type NodeList struct {
	items []*Node
	sync.RWMutex
}

func NewNodeList() *NodeList {
	q := &NodeList{
		items: []*Node{},
	}
	return q
}

func (q *NodeList) Size() int {
	return len(q.items)
}

func (q *NodeList) Enqueue(n *Node) {
	q.Lock()
	q.items = append(q.items, n)
	q.Unlock()
}

func (q *NodeList) Dequeue() *Node {
	var n *Node = nil
	q.Lock()
	ql := len(q.items)
	if ql > 0 {
		n = q.items[0]
		q.items = q.items[1:ql]
	}
	q.Unlock()
	return n
}

func (q *NodeList) Print() {
	fmt.Printf("[%d]: ", len(q.items))
	for _, i := range q.items {
		fmt.Printf("%d, ", i.item)
	}
	fmt.Printf("\n")
}

//
// Node Stack
//
type NodeStack struct {
	items []*Node
	sync.RWMutex
}

func NewNodeStack() *NodeStack {
	q := &NodeStack{
		items: []*Node{},
	}
	return q
}

func (q *NodeStack) Size() int {
	return len(q.items)
}

func (q *NodeStack) Push(n *Node) {
	q.Lock()
	q.items = append(q.items, n)
	q.Unlock()
}

func (q *NodeStack) Pop() *Node {
	var n *Node = nil
	q.Lock()
	ql := len(q.items)
	if ql > 0 {
		n = q.items[ql-1]
		q.items = q.items[0 : ql-1]
	}
	q.Unlock()
	return n
}

func (q *NodeStack) Print() {
	fmt.Printf("[%d]: ", len(q.items))
	for i := len(q.items); i > 0; i-- {
		fmt.Printf("%d, ", q.items[i-1].item)
	}
	fmt.Printf("\n")
}

func test_node_queue() {
	n1 := &Node{item: 56}
	n2 := &Node{item: 64}
	n3 := &Node{item: 19}

	// Check FIFO
	fmt.Println("------ Queue FIFO ------")
	fifoQ := NewNodeList()
	fifoQ.Print()
	fifoQ.Enqueue(n1)
	fifoQ.Enqueue(n2)
	fifoQ.Enqueue(n3)
	fifoQ.Print()
	n := fifoQ.Dequeue()
	fmt.Printf("Dequeue: %d\n", n.item)
	fifoQ.Print()
	n = fifoQ.Dequeue()
	fmt.Printf("Dequeue: %d\n", n.item)
	fifoQ.Print()

	// Check LIFO
	fmt.Println("------ Queue LIFO (Stack) ------")
	stack := NewNodeStack()
	stack.Push(n1)
	stack.Push(n2)
	stack.Push(n3)
	stack.Print()
	n = stack.Pop()
	fmt.Printf("Stack pop: %d\n", n.item)
	stack.Print()
	n = stack.Pop()
	fmt.Printf("Stack pop: %d\n", n.item)
	stack.Print()
}
