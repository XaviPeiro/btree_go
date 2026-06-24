package main

type Node struct {
	keys []int `json:"keys,omitempty"`
	leaf bool  `json:"leaf,omitempty"`
	children []*Node
	parent *Node
	// is_root bool
}


func NewNode(leaf bool, keys []int, children []*Node, parent *Node) *Node{
	if len(children) == 0 {
		children = make([]*Node, 0, MaxChildren)
	}

	if  len(keys) == 0 {
		keys = make([]int, 0, MaxKeys)
	}

	node := Node{
		keys: keys,
		leaf: leaf,
		children: children,
		parent: parent,
	}
	return &node
}
// copy()


func (node *Node) insertInNodeIndex(key Key) (bool, string) {
	// to_insert := key
	// // slice len 0, cap set -> should not loop if no item
	// for i, k := range node.keys {
	// 	if k <= key {
	// 		continue
	// 	}
	// 	tmp := k
	// 	node.keys[i] = to_insert
	// 	to_insert = tmp
	// }
	// node.keys = append(node.keys, to_insert)
	sortedInsert(node.keys, key)
	return true, ""
}


// Node version 2

type Node2 struct {
	leaf bool
	keys []*Element
}

type Element struct {
	left *Node2
	right *Node2
	key Key
}

func NewNode2(leaf bool) *Node2 {
	node := Node2{
		leaf: leaf,
		keys: make([]*Element, 0, MaxKeys),
	}

	return &node	
}

func NewElement(left *Node2, right *Node2, key Key) *Element{
	element := Element{
		left: left,
		right: right,
		key: key,
	}
	return &element
}

func (node *Node2) insertInNode(key Key) (bool, string) {
	to_insert := NewElement(nil,nil,key)

	// slice len 0, cap set -> should not loop if no item
	for i, k := range node.keys {
		if k.key <= key {
			continue
		}
		tmp := k
		node.keys[i] = to_insert
		to_insert = tmp
	}
	node.keys = append(node.keys, to_insert)
	return true, ""
}