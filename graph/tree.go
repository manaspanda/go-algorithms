package main

import (
	"fmt"
)

type Tree struct {
	val      int
	children map[int]*Tree
}

type TreePlugin func(node *Tree)

func AddChild(node *Tree, key int) *Tree {
	child := &Tree{
		val:      key,
		children: make(map[int]*Tree),
	}
	if node != nil {
		node.children[key] = child
	}
	return child
}

func RemoveChild(node *Tree, key int) {
	if node != nil {
		delete(node.children, key)
	}
}

func MaxDepth(node *Tree) int {
	if node == nil {
		return 0
	}
	numChildren := len(node.children)
	if numChildren == 0 {
		return 1
	}
	childDepth := make([]int, numChildren)
	i := 0
	for _, v := range node.children {
		childDepth[i] = MaxDepth(v)
		i++
	}
	maxChildDepth := 0
	for i := 0; i < numChildren; i++ {
		if childDepth[i] > maxChildDepth {
			maxChildDepth = childDepth[i]
		}
	}
	return maxChildDepth + 1
}

func TreeNodePrint(node *Tree) {
	fmt.Printf("%d ", node.val)
}

func TreeWalkPreOrder(root *Tree, pl TreePlugin) {
	pl(root)
	for _, v := range root.children {
		TreeWalkPreOrder(v, pl)
	}
}

func TreeWalkPostOrder(root *Tree, pl TreePlugin) {
	for _, v := range root.children {
		TreeWalkPostOrder(v, pl)
	}
	pl(root)
}

func test_tree() {
	t1 := AddChild(nil, 1)
	t1_c1 := AddChild(t1, 3)
	AddChild(t1, 2)
	AddChild(t1, 4)
	AddChild(t1_c1, 5)
	AddChild(t1_c1, 6)

	fmt.Println("\n------ Tree:walk-preorder -------")
	TreeWalkPreOrder(t1, TreeNodePrint)
	fmt.Println("\n------ Tree:walk-postorder -------")
	TreeWalkPostOrder(t1, TreeNodePrint)
	fmt.Printf("\nMax-Depth=%d", MaxDepth(t1))
	fmt.Println("\n------------------")
}
