package main

// type Key = string
type Key = int // GEneralise later 
type Value = string
type Index = int
type Found = bool

/*
T = Minimum degree (minimum number of children).
max children = 2*T

Minimum number of keys: T-1
Minimum number of keys: 2*T-1
*/
const T = 2 
const MaxChildren = 2*T
const MaxKeys = 2*T-1
const MinKeys = T-1

// TODO: optimization -> keys array
type Node struct {
	keys []int `json:"keys,omitempty"`
	leaf bool  `json:"leaf,omitempty"`
	children []*Node
	// is_root bool
}

// func (n *Node) Insert(key Key) 

func NewNode() *Node{
	node := Node{
		keys: make([]int, 0, MaxKeys),
		leaf: true,
		children: make([]*Node, 0, MaxChildren),

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
		keys: make([]int, 0, MaxKeys),
		leaf: true,
		children: make([]*Node, 0, MaxChildren), // Q? What is the point of pre-allocating? I do not remember the problem with dynamic allocation in GO
	}

	tree := BTree{
		root: root,
	}

	return &tree
}
func (b *BTree) InsertInNodeIndex(node *Node, index Index, key Key) (bool, string) {
    if len(node.keys) >= MaxKeys {
		// We should rebalance, but isn't implemented yet
		panic("Insert, node full should be split. Not implemented yet")
    }

	if len(node.keys) == 0 {
		node.keys = append(node.keys, key)
		return true, ""
	}

	to_insert := key
	for i, k := range node.keys {
		if i < index{
			continue
		}
		tmp := k
		node.keys[i] = to_insert
		to_insert = tmp
	}
	return true, ""
}

func (b *BTree) InsertInNode(node *Node, key Key) (bool, string) {
    if len(node.keys) >= MaxKeys {
		// We should rebalance, but isn't implemented yet
		panic("Insert, node full should be split. Not implemented yet")
    }

	if len(node.keys) == 0 {
		node.keys = append(node.keys, key)
		return true, ""
	}

	to_insert := key
	for i, k :=  range node.keys {
		if k>to_insert {
			node.keys[i] = to_insert
			to_insert = k
		} 
	}
	return true, ""
}


func (b *BTree) Insert(key Key) (bool, string) {
	node, _, found := b.Search(key, &b.root)
	if found == false {
		b.InsertInNode(node, key)
		return true, ""
	}

	return false, "Key must be unique"
}

func (b *BTree) Search(target_key Key, starting_node *Node) (*Node, Index, Found) {
	/*
		:return:
		- node: in which the target_key should be (found or not)
		- index: where it should be place in the node
		- found: it is already present or not
		
	*/
	node := *starting_node
	c := 0
	
	for c<len(node.keys) && node.keys[c] < target_key{
		c++
	}  

	if c < len(node.keys) && node.keys[c] == target_key{
		return &node, c, true
	}
	
	if node.leaf == true {
		return &node, c, false
	} else {
		return b.Search(target_key, node.children[c])
	}
}



//func main() {
//}
