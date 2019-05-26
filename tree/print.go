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

func levelOrder(root *Node) {
	if root == nil {
		return
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		fmt.Println(node.key)
		queue = queue[1:]
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
}
