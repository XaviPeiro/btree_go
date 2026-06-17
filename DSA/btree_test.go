package main

import (
	"slices"
	"testing"
	"fmt"
)

// func TestInsertNodes(t *testing.T) {
// 	btree := NewBTree()
// 	btree.Insert()
// }

func TestBtreeInsertKey1NodeRebalance(t *testing.T) {
	tests := []struct {
		name string
		keys []int
		key  int
		expected_root_keys []int
	}{
		{
			name: "insert first",
			keys: []int{11,13,14},
			key:  10,
			expected_root_keys: []int{13},
		},
		{
			name: "insert last",
			keys: []int{20, 30, 40},
			key:  50,
			expected_root_keys: []int{40},
		},
		{
			name: "insert in middle",
			keys: []int{10, 20, 30},
			key:  15,
			expected_root_keys: []int{20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()

			btree := NewBTree()
			for _, k := range tt.keys {
				btree.InsertInNode(&btree.root, k)
				fmt.Printf("key loop %d \n", k)
			}
			fmt.Println(btree.root.keys)
			btree.InsertInNode(&btree.root, tt.key)

			if !slices.Equal(btree.root.keys, tt.expected_root_keys) {
				t.Fatalf("InsertInNode(%v, %d) = %v, want %v",
					tt.keys, tt.key, &btree.root.keys, tt.expected_root_keys)
			}
		})
	}
}

func TestBtreeInsertKeyNoRebalance(t *testing.T) {
	tests := []struct {
		name string
		keys []int
		key  int
		want []int
	}{
		{
			name: "insert into empty slice",
			keys: []int{},
			key:  10,
			want: []int{10},
		},
		{
			name: "insert at beginning",
			keys: []int{20, 30},
			key:  10,
			want: []int{10, 20, 30},
		},
		{
			name: "insert in middle",
			keys: []int{10, 30},
			key:  20,
			want: []int{10, 20, 30},
		},
		{
			name: "insert at end",
			keys: []int{10, 20},
			key:  30,
			want: []int{10, 20, 30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()

			btree := NewBTree()
			for _, k := range tt.keys {
				btree.InsertInNode(&btree.root, k)
				fmt.Printf("key loop %d \n", k)
			}
			// fmt.Println(btree.root.keys)
			btree.InsertInNode(&btree.root, tt.key)

			if !slices.Equal(btree.root.keys, tt.want) {
				t.Fatalf("InsertInNode(%v, %d) = %v, want %v",
					tt.keys, tt.key, &btree.root.keys, tt.want)
			}
		})
	}
}

// func TestBTreeInsertAndSearch(t *testing.T) {
// 	tree := NewBTree()

// 	pairs := []struct {
// 		key   Key
// 		value Value
// 	}{
// 		{"d", "delta"},
// 		{"b", "bravo"},
// 		{"a", "alpha"},
// 		{"c", "charlie"},
// 		{"e", "echo"},
// 		{"f", "foxtrot"},
// 		{"g", "golf"},
// 		{"h", "hotel"},
// 		{"i", "india"},
// 	}

// 	for _, pair := range pairs {
// 		if inserted := tree.Insert(pair.key, pair.value); !inserted {
// 			t.Fatalf("Insert(%q) reported an update, want new insertion", pair.key)
// 		}
// 		assertBTreeInvariants(t, tree)
// 	}

// 	for _, pair := range pairs {
// 		value, found := tree.Search(pair.key)
// 		if !found {
// 			t.Fatalf("Search(%q) was not found", pair.key)
// 		}
// 		if value != pair.value {
// 			t.Fatalf("Search(%q) = %q, want %q", pair.key, value, pair.value)
// 		}
// 	}

// 	if value, found := tree.Search("z"); found {
// 		t.Fatalf("Search(%q) = %q, want missing key", "z", value)
// 	}
// }

// func TestBTreeUpdateExistingKey(t *testing.T) {
// 	tree := NewBTree()

// 	if inserted := tree.Insert("c", "old"); !inserted {
// 		t.Fatal("first insert reported an update")
// 	}
// 	if inserted := tree.Insert("c", "new"); inserted {
// 		t.Fatal("second insert reported a new key, want update")
// 	}

// 	value, found := tree.Search("c")
// 	if !found {
// 		t.Fatal("updated key was not found")
// 	}
// 	if value != "new" {
// 		t.Fatalf("Search(%q) = %q, want %q", "c", value, "new")
// 	}

// 	assertBTreeInvariants(t, tree)
// }

// func TestBTreeRootSplit(t *testing.T) {
// 	tree := NewBTree()

// 	for _, pair := range []struct {
// 		key   Key
// 		value Value
// 	}{
// 		{"a", "alpha"},
// 		{"b", "bravo"},
// 		{"c", "charlie"},
// 		{"d", "delta"},
// 	} {
// 		tree.Insert(pair.key, pair.value)
// 	}

// 	if tree.root.leaf {
// 		t.Fatal("root is still a leaf after enough insertions to split")
// 	}
// 	if len(tree.root.keys) != 1 || tree.root.keys[0] != "b" {
// 		t.Fatalf("root keys = %v, want [b]", tree.root.keys)
// 	}
// 	if len(tree.root.children) != 2 {
// 		t.Fatalf("root has %d children, want 2", len(tree.root.children))
// 	}

// 	assertBTreeInvariants(t, tree)
// }

// func assertBTreeInvariants(t *testing.T, tree *BTree) {
// 	t.Helper()

// 	if tree == nil || tree.root == nil {
// 		t.Fatal("tree root is nil")
// 	}

// 	leafDepth := -1
// 	validateNode(t, tree.root, true, "", false, "", false, 0, &leafDepth)
// }

// func validateNode(
// 	t *testing.T,
// 	node *Node,
// 	isRoot bool,
// 	lower Key,
// 	hasLower bool,
// 	upper Key,
// 	hasUpper bool,
// 	depth int,
// 	leafDepth *int,
// ) {
// 	t.Helper()

// 	if len(node.keys) != len(node.values) {
// 		t.Fatalf("keys length %d differs from values length %d", len(node.keys), len(node.values))
// 	}
// 	if len(node.keys) > maxKeys {
// 		t.Fatalf("node keys %v exceed max key count %d", node.keys, maxKeys)
// 	}
// 	if !isRoot && len(node.keys) < T-1 {
// 		t.Fatalf("non-root node keys %v have fewer than %d keys", node.keys, T-1)
// 	}

// 	for index, key := range node.keys {
// 		if index > 0 && node.keys[index-1] >= key {
// 			t.Fatalf("node keys are not strictly sorted: %v", node.keys)
// 		}
// 		if hasLower && key <= lower {
// 			t.Fatalf("key %q violates lower bound %q", key, lower)
// 		}
// 		if hasUpper && key >= upper {
// 			t.Fatalf("key %q violates upper bound %q", key, upper)
// 		}
// 	}

// 	if node.leaf {
// 		if len(node.children) != 0 {
// 			t.Fatalf("leaf node %v has %d children", node.keys, len(node.children))
// 		}
// 		if *leafDepth == -1 {
// 			*leafDepth = depth
// 		}
// 		if depth != *leafDepth {
// 			t.Fatalf("leaf depth = %d, want %d", depth, *leafDepth)
// 		}
// 		return
// 	}

// 	if len(node.children) != len(node.keys)+1 {
// 		t.Fatalf("internal node %v has %d children, want %d", node.keys, len(node.children), len(node.keys)+1)
// 	}

// 	for index, child := range node.children {
// 		childLower := lower
// 		childHasLower := hasLower
// 		childUpper := upper
// 		childHasUpper := hasUpper

// 		if index > 0 {
// 			childLower = node.keys[index-1]
// 			childHasLower = true
// 		}
// 		if index < len(node.keys) {
// 			childUpper = node.keys[index]
// 			childHasUpper = true
// 		}

// 		validateNode(t, child, false, childLower, childHasLower, childUpper, childHasUpper, depth+1, leafDepth)
// 	}
// }
