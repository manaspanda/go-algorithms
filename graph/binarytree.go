package main

import "fmt"

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func Insert(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return &TreeNode{val: key}
	}
	if key < root.val {
		root.left = Insert(root.left, key)
	} else {
		root.right = Insert(root.right, key)
	}
	return root
}

func Find(root *TreeNode, key int) *TreeNode {
	if root == nil || root.val == key {
		return root
	}
	if key < root.val {
		return Find(root.left, key)
	}
	return Find(root.right, key)
}

func Next(node *TreeNode) *TreeNode {
	if node == nil {
		return node
	}
	// Find the left-most node on the right
	next := node.right
	for next != nil && next.left != nil {
		next = next.left
	}
	return next
}

func Delete(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if key == root.val {
		// If left is nil, replace the right
		if root.left == nil {
			return root.right
		}
		// If right is nil, replace the left
		if root.right == nil {
			return root.left
		}
		// Node has both left and right, find the successor
		// swap the value, and delete
		succ := Next(root)
		root.val = succ.val
		root.right = Delete(root.right, succ.val)
		return root
	}

	if key < root.val {
		root.left = Delete(root.left, key)
	} else {
		root.right = Delete(root.right, key)
	}
	return root
}

func Height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := Height(root.left)
	rh := Height(root.right)
	h := lh
	if rh > lh {
		h = rh
	}
	return h + 1
}

type BSTNodePlugin func(t *TreeNode)

func NodePrint(t *TreeNode) {
	if t != nil {
		fmt.Printf("%d ", t.val)
	}
}

func BSTWalkInOrder(t *TreeNode, pl BSTNodePlugin) {
	if t == nil {
		return
	}
	BSTWalkInOrder(t.left, pl)
	pl(t)
	BSTWalkInOrder(t.right, pl)
}

func BSTWalkPreOrder(t *TreeNode, pl BSTNodePlugin) {
	if t == nil {
		return
	}
	pl(t)
	BSTWalkPreOrder(t.left, pl)
	BSTWalkPreOrder(t.right, pl)
}

func BSTWalkPostOrder(t *TreeNode, pl BSTNodePlugin) {
	if t == nil {
		return
	}
	BSTWalkPostOrder(t.left, pl)
	BSTWalkPostOrder(t.right, pl)
	pl(t)
}

func test_bst() {
	bst1 := Insert(nil, 2)
	Insert(bst1, 1)
	Insert(bst1, 7)
	Insert(bst1, 4)
	Insert(bst1, 8)
	Insert(bst1, 3)
	Insert(bst1, 6)
	Insert(bst1, 5)
	fmt.Println("\n------ BST:walk-in -------")
	BSTWalkInOrder(bst1, NodePrint)
	fmt.Println("\nHeight =>", Height(bst1))
	fmt.Println("\n------ BST:walk-pre -------")
	BSTWalkPreOrder(bst1, NodePrint)
	fmt.Println("\n------ BST:walk-post -------")
	BSTWalkPostOrder(bst1, NodePrint)
	fmt.Println("\n------ BST:delete 3 (no-child)--------")
	Delete(bst1, 3)
	BSTWalkInOrder(bst1, NodePrint)
	fmt.Println("\n------ BST:delete 6 (one-child) --------")
	Insert(bst1, 3)
	Delete(bst1, 6)
	BSTWalkInOrder(bst1, NodePrint)
	fmt.Println("\n------ BST:delete 4 (two-child) --------")
	Delete(bst1, 5)
	Insert(bst1, 6)
	Insert(bst1, 5)
	Delete(bst1, 4)
	BSTWalkInOrder(bst1, NodePrint)
	fmt.Println("\nHeight =>", Height(bst1))
	fmt.Println("\n------------------")
}
