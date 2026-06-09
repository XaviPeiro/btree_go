package main

// type Key = string
type Key = int // GEneralise later 
type Value = string
type Index = int
/*
T = Minimum degree (minimum number of children).
max children = 2*T

Minimum number of keys: T-1
Minimum number of keys: 2*T-1
*/
const T = 2 

// TODO: optimization -> keys array
type Node struct {
	keys []int `json:"keys,omitempty"`
	leaf bool  `json:"leaf,omitempty"`
	children []*Node
}

func NewNode() *Node{
	node := Node{
		keys: make([]int, 2*T-1),
		leaf: true,
		children: make([]*Node, 2*T),	
	}
	return &node
}

type BTree struct {
	root Node
}

func NewBTree() *BTree {
	// Q? How does struct exactly work? What are their point apart than encapsulation?
	/*
	keys: values in the node
	leaf: no children - end node
	children: 
	*/
	root := Node{
		keys: make([]int, 2*T-1),
		leaf: true,
		children: make([]*Node, 2*T), // Q? What is the point of pre-allocating? I do not remember the problem with dynamic allocation in GO
	}

	tree := BTree{
		root: root,
	}

	return &tree
}
func (b *BTree) Insert(key Key) bool {
	node, index := b.Search(key, &b.root)
	if node == nil {
		// is leaf and key not found
		// node.keys = [:index] + key + [index:]  
		node.keys = append(node.keys, )
	} else {

	}
	return false
}

func (b *BTree) Search(target_key Key, starting_node *Node) (*Node, Index) {
	node := *starting_node
	c := 0
	
	for c<len(node.keys) && node.keys[c] < target_key{
		c++
	}  

	if c < len(node.keys) && node.keys[c] == target_key{
		return &node, c
	}
	
	if node.leaf == true {
		return nil, c
	} else {
		return b.Search(target_key, node.children[c])
	}
}



//func main() {
//}
