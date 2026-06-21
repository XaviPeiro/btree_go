package main

import (
	"container/list"
	"fmt"
	"image/color/palette"
	"math"
)

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

// func rebalance


func (b *BTree) Insert(key Key) (bool, string) {
	node, _, found := b.Search(key, &b.root)
	if found == true {
		return false, "Key must be unique"
	}

	// At this point we that it does not exist and where it should be.
	// Due to search, it should be always a leaf, however double-check just in case.
	if node.leaf == false {
		panic("ERRROR: BITCH: This it not a leaf, we cannot insert in here!")
	}


	// We set as an invariant that there is always place, so we split the node after insert if it is full.
	has_inserted, err := node.insertInNodeIndex(key)
    if len(node.keys) < MaxKeys {
		// Inserted, everything is cool, nothing to do.
		return has_inserted, err
	}

	// We have inserted and now the node is full (of keys)
	// La chicha; we need to rebalance the tree
	/*
		When full rebalance in 2 nodes: maxk=2T-1 mink=T-1 -> solo una posibilidad, 2 nodos	
	*/

	// Divide keys in 2: 1/2 current, 1/2 new sibling, remaining (always 1) to the parent.
	// TODO: Instead of having ekeys and node apart I could have just nodes with left and right.!
	left_ks, parentk, right_ks := node.keys[:(T-1)], node.keys[T-1], node.keys[T:(2*T)-1] 
	panic("Insert, node full should be split. Not implemented yet")

}

func (b *BTree) Search(target_key Key, starting_node *Node) (*Node, Index, Found) {
	/*
		:return:
		- node: in which the target_key should be (found or not)
		- index: where it should be place in the node
		- found: it is already present or not
		
	*/
	node := starting_node
	current_index := 0
	
	for current_index<len(node.keys) && node.keys[current_index] < target_key{
		current_index++
	}  

	if current_index < len(node.keys) && node.keys[current_index] == target_key{
		return node, current_index, true
	}
	
	if node.leaf == true {
		return node, current_index, false
	} else {
		return b.Search(target_key, node.children[current_index])
	}
}


func (b *BTree) String() string {
	/*
	This is just a representation to easily check if our btree is correct.

	The aim is BFS + controlling each level; simple.
	*/
	var result [][]int = [][]int{}
	// result = append(result, 1)

	q := list.New()
	q.PushBack([]*Node{&b.root})

	lvl := 0.0
	// BFS
	for q.Len() > 0 {
		lvl++

		var next_lvl_values []*Node = make([]*Node, 0, int(math.Pow(T, lvl)))
		var lvl_res []int = []int{}
		// var next_lvl_values []*Node = make([]*Node, 0)

		front := q.Front()
		nodes_in_lvl := front.Value.([]*Node)
		q.Remove(front)
		
		for _, node := range nodes_in_lvl {
			for _, child_node := range node.children {
				// nodes
				next_lvl_values = append(next_lvl_values, child_node)
			}

			// keys
			lvl_res = append(lvl_res, node.keys...)
		}

		if len(next_lvl_values) > 0{
			q.PushBack(next_lvl_values)
		}
		result = append(result, lvl_res)
	} 

	// stringify
	return fmt.Sprint(result)
}
//func main() {
//}
