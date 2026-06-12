package main

/*
BTree
---

# Invariants

Most of them, if not all, are oriented to keep the tree balanced.

## Children and keys
All nodes have a mD (minimum degree, in other words, minimum number of children) named T.
A MD (maximum degree) of 2T.

By extension, mK (minimum keys) of T-1 and MK of 2T-1.

Root node is the exception, it may has less than T.

# Leafs
A leaf is a node with no children.

All leaves shall be at the very same level.


# Insert
Always at the bottom of the tree.
When a node is full (2T-1 K) is divided by 2 and the remaining one is promoted to the parent (depth-1).
This is done recursively.
=> 
- Ensures all leafs are at the very same level (so same depth).
- *The filters adapt to the number of keys*; in the case of a monotonic one the keys promoted will the lastly entered,
increasing the gap btw 'em (greater filter) and promoting them in concordance with the inserts.*
I SHOULD ANALYZE THIS BETTER, the only clear thing I see is that it adapt to the 'hottest trend' by promoting them up. 
For random keys as UUID it may be problematic.

*/
// type Key = string
type Key = int // Generalize later 
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
	// This probably better be a Node method
    if len(node.keys) >= MaxKeys {

		/*
		When full rebalance in 2 nodes: maxk=2T-1 mink=T-1 -> solo una posibilidad, 2 nodos	

		*/
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
	node.keys = append(node.keys, to_insert)
	return true, ""
}

// func rebalance

func (b *BTree) InsertInNode(node *Node, key Key) (bool, string) {
	var index Index = 0
	for index < len(node.keys) && node.keys[index]<=key {
		index++
	}

	return b.InsertInNodeIndex(node, index, key)
}


func (b *BTree) Insert(key Key) (bool, string) {
	node, _, found := b.Search(key, &b.root)
	if found == false {
		b.InsertInNode(node, key)
		return true, ""
	}

	return false, "Key must be unique"
}

func (btree *BTree) SearchForInsert() {

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
