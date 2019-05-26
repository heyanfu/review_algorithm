package main

import "fmt"

func preOrder(root *Node) {
	if root != nil {
		fmt.Printf("key = %d, val = %d", root.key, root.val)
		fmt.Println()
		preOrder(root.left)
		preOrder(root.right)
	}
}

func middleOrder(root *Node) {
	if root != nil {
		preOrder(root.left)
		fmt.Printf("key = %d, val = %d", root.key, root.val)
		fmt.Println()
		preOrder(root.right)
	}
}

func postOrder(root *Node) {
	if root != nil {
		preOrder(root.left)
		preOrder(root.right)
		fmt.Printf("key = %d, val = %d", root.key, root.val)
		fmt.Println()
	}
}
