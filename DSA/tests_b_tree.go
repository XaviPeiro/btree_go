package main

import "fmt"

func main() {
	fmt.Println("Starting BTree...")
	btree := NewBTree()
	btree.Search("asdf")

	return

}
