package main

type Key = string
type Value = string

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
	children []Node 
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
		children: make([]Node, 2*T), // Q? What is the point of pre-allocating? I do not remember the problem with dynamic allocation in GO
	}

	tree := BTree{
		root: root,
	}

	return &tree
}

func (b *BTree) Search(target_key Key, starting_node *Node) (Key, Value) {

	// Missing default value for starting_node
	if starting_node == nil {
		starting_node = &b.root
	}

	// if starting_node == nil {
	// 	node := b.root
	// }

	/*
	logic:
		node = root
		node_key = nil
		c = 0
		while c<len(node.keys) and node_key < target_key:
			node_key = keys[c]
			c++

		if node_key == target_key:
			return node_key
		else:
			if node_key.leaf is false:
				search(target_key, node_key.child)
			else:
				return nil

	*/
	return "", ""
}

func (b BTree) Insert(key Key, value Value) bool {
	// Spot the position in which should be inserted
	/*
		If keys < 2*T-1 ->  insert normally
		else
			(recursive) split node in 2, update parents.
	*/

	return false
}

//func main() {
//}
