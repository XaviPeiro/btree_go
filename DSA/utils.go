package main

import "cmp"
// type Index int
func sortedInsert[T cmp.Ordered](collection []T, value T) Index {
	var i Index
	var v T

	collection = append(collection, v)
	for i, v = range collection{
		if v > value {
			break
		}
	}

	copy(collection[i+1:], collection[i:])
	collection[i] = value

	return i
}

func insertAtIndex[T any](collection []T, value T, i Index){
	var v T
	collection = append(collection, v)
	
	copy(collection[i+1:], collection[i:])
	collection[i] = value
}