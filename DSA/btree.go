package main

type Key = string
type Value = string

const T = 2

// TODO: optimization -> keys array
type Node struct {
	keys []int `json:"keys,omitempty"`
	leaf bool  `json:"leaf,omitempty"`
}

type BTree struct {
	root Node
}

func (b BTree) NewBTree() *BTree {

	root := Node{
		keys: make([]int, 2*T-1),
		leaf: true,
	}
	tree := BTree{
		root: root,
	}

	return &tree
}

func (b BTree) Search() (Key, Value) {
	return "", ""
}
func (b BTree) Insert() bool {}

func main() {
}
