package main

type Node struct {
	keys []int `json:"keys,omitempty"`
	leaf bool  `json:"leaf,omitempty"`
	children []*Node
	// is_root bool
}

func NewNode() *Node{
	node := Node{
		keys: make([]int, 0, MaxKeys),
		leaf: true,
		children: make([]*Node, 0, MaxChildren),

	}
	return &node
}