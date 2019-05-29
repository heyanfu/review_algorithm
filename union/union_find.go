package main

import "fmt"

func main() {
	uf := &UnionFind{[]int{}, []int{}, 0}
	uf.init()
	uf.unionElements(3, 5)
	fmt.Println(uf.parent)
	fmt.Println(uf.rank)
}

type UnionFind struct {
	rank   []int
	parent []int
	count  int
}

func (uf *UnionFind) init() {
	uf.count = 10
	for i := 0; i < 10; i++ {
		uf.parent = append(uf.parent, i)
		uf.rank = append(uf.rank, 1)
	}
}

func (uf *UnionFind) find(p int) int {
	for p != uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]]
		p = uf.parent[p]
	}

	return p
}

func (uf *UnionFind) isConnect(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *UnionFind) unionElements(p, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)
	if pRoot == qRoot {
		return
	}
	if uf.rank[pRoot] < uf.rank[qRoot] {
		uf.parent[pRoot] = uf.parent[qRoot]
	} else if uf.rank[pRoot] > uf.rank[qRoot] {
		uf.parent[qRoot] = uf.parent[pRoot]
	} else {
		uf.parent[pRoot] = uf.parent[qRoot]
		uf.rank[qRoot]++
	}
}
