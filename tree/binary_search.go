package main

func main() {
	arr := []int{6, 4, 7, 1, 3, 2, 5, 9, 8, 0}
	var root *Node
	for _, v := range arr {
		root = root.addNode(v, v)
	}
	//preOrder(root)
	//fmt.Println(root.hasKey(8))
	//fmt.Println(root.hasKey(10))
	//fmt.Println(root.getValByKey(1))
	//fmt.Println(root.getValByKey(11))
	levelOrder(root)
}

type Node struct {
	key   int
	val   int
	left  *Node
	right *Node
}

func (node *Node) addNode(key, val int) *Node {
	if node == nil {
		return &Node{key, val, nil, nil}
	}
	if key == node.key {
		node.val = val
	} else if key < node.key {
		node.left = node.left.addNode(key, val)
	} else {
		node.right = node.right.addNode(key, val)
	}

	return node
}

func (node *Node) hasKey(key int) bool {
	if node == nil {
		return false
	}
	if key == node.key {
		return true
	} else if key < node.key {
		return node.left.hasKey(key)
	} else {
		return node.right.hasKey(key)
	}
}

func (node *Node) getValByKey(key int) int {
	if node == nil {
		return -1
	}
	if key == node.key {
		return node.val
	} else if key < node.key {
		return node.left.getValByKey(key)
	} else {
		return node.right.getValByKey(key)
	}
}
