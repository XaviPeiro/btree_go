package main

type Key = string
type Value = string

// type I=
const T = 2

//type Entry struct {
//	Key Key
//	//Val Value
//	LeftNode  Node
//	RightNode Node
//}

// TODO: optimization -> keys array
type Node struct {
	Keys     []Key //`json:"keys,omitempty"`
	Children []*Node
	Leaf     bool //`json:"leaf,omitempty"
}

type BTree struct {
	root Node
}

func (b BTree) NewBTree() *BTree {
	root := Node{
		Keys:     make([]Key, 2*T-1),
		Leaf:     true,
		Children: make([]*Node, 2*T-1),
	}
	tree := BTree{
		root: root,
	}

	return &tree
}

func (b BTree) Search(key Key, node *Node) (int, *Node) {
	current_node := b.root
	var i int = 0
	// this On I guess it could be logn
	for i < len(current_node.Keys) && key > current_node.Keys[i] {
		i++
	}

	if key == current_node.Keys[i] {
		return i, node
	}
	if current_node.Leaf == true {
		return i, current_node
	} else {
		return b.Search(key, current_node.Children[i])
	}

	return ""
}
func (b BTree) InsertNode(key Key) bool {
	// if key does not exist
	b.Search(key)
}

func main() {
}
