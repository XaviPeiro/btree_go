package main

import "cmp"

func sortedInsert[T cmp.Ordered](collection *[]T, value T) Index {
	var i Index
	var v T

	*collection = append(*collection, v)
	s := *collection

	for i, v = range s{
		if v > value {
			break
		}
	}
	copy(s[i+1:], s[i:])
	s[i] = value

	return i
}

func insertAtIndex[T any](collection *[]T, value T, i Index){
	var v T
	*collection = append(*collection, v)
	s := *collection
	
	copy(s[i+1:], s[i:])
	s[i] = value
}