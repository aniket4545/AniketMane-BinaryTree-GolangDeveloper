package main

import (
	"fmt"
	"sync"
)

type Tree struct {
	sync.RWMutex
	root *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func main() {

	arr := []int{10, 20, 30, 4, 5, 6}
	binaryTree := new(Tree)

	for _, val := range arr {
		binaryTree.InsertNode(val)
	}

	binaryTree.Print()

	fmt.Println("----PREORDER------>>", binaryTree.PreOrder())
	fmt.Println("----POSTORDER----->>", binaryTree.PostOrder())
	fmt.Println("-----INORDER------>>", binaryTree.InOrder())
}

func newNode(value int) *Node {

	return &Node{
		value: value,
	}
}

// insert value inside the tree
func (t *Tree) InsertNode(value int) {
	t.RWMutex.Lock()
	defer t.RWMutex.Unlock()

	if t.root == nil {
		t.root = newNode(value)
		return
	}
	insert(value, t.root)
}

// if new value is less then the current node then we will have it as new left child
// if new value is greater then current node then we will have it as new right child
func insert(value int, node *Node) {

	if value < node.value {

		if node.left == nil {
			node.left = newNode(value)
		} else {
			insert(value, node.left)
		}
	} else {

		if node.right == nil {
			node.right = newNode(value)
		} else {
			insert(value, node.right)
		}
	}
}

// left -> root -> right
func (t *Tree) InOrder() []int {
	t.RWMutex.RLock()
	defer t.RWMutex.RUnlock()

	var inOrderArray []int
	inOrder(t.root, &inOrderArray)
	return inOrderArray
}

func inOrder(node *Node, array *[]int) {

	if node.left != nil {
		inOrder(node.left, array)
	}

	*array = append(*array, node.value)

	if node.right != nil {
		inOrder(node.right, array)
	}
}

// root -> left -> right
func (t *Tree) PreOrder() []int {
	t.RWMutex.RLock()
	defer t.RWMutex.RUnlock()

	var preOrderArray []int
	preOrder(t.root, &preOrderArray)
	return preOrderArray
}

func preOrder(node *Node, array *[]int) {

	*array = append(*array, node.value)

	if node.left != nil {
		preOrder(node.left, array)
	}
	if node.right != nil {
		preOrder(node.right, array)
	}
}

// left -> right -> root
func (t *Tree) PostOrder() []int {
	t.RWMutex.RLock()
	defer t.RWMutex.RUnlock()

	var postOrderArray []int
	postOrder(t.root, &postOrderArray)
	return postOrderArray
}

func postOrder(node *Node, array *[]int) {

	if node.left != nil {
		postOrder(node.left, array)
	}
	if node.right != nil {
		postOrder(node.right, array)
	}
	*array = append(*array, node.value)
}

func (t *Tree) Print() {
	t.RWMutex.RLock()
	defer t.RWMutex.RUnlock()

	fmt.Printf("\n ---------- YOUR BINARY TREE LOOKS LIKE --------------\n  %s\n\n", " M-> ROOT-NODE; L-> LEVEL; Ln-> LEFT-NODE; Rn-> RIGHT-NODE")
	print(t.root, 0, 0, "M")
	fmt.Println("\n            ---------------------------              ")
}

func print(node *Node, space, level int, nodeType string) {
	if node == nil {
		return
	}

	fmt.Printf(" L:%d", level)
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("   - %s: %v\n", nodeType, node.value)
	level++
	print(node.left, space+4, level, "Ln")
	print(node.right, space+4, level, "Rn")
}
