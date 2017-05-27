package data

import (
	"sort"
)

// IntSet is a sorted immutable integer set
type IntSet struct {
	data []int
}

// NewIntSet creates a new integer set
func NewIntSet() IntSet {
	return IntSet{[]int{}}
}

// Len returns with the length of the set
func (i IntSet) Len() int {
	return len(i.data)
}

// Insert adds a new item to the set
func (i IntSet) Insert(val int) IntSet {
	if len(i.data) == 0 {
		return IntSet{[]int{val}}
	}
	index := sort.SearchInts(i.data, val)
	if index < len(i.data) && i.data[index] == val {
		return i
	}

	i2 := IntSet{make([]int, len(i.data)+1)}
	copy(i2.data[:index], i.data[:index])
	copy(i2.data[index+1:], i.data[index:])
	i2.data[index] = val

	return i2
}

// Union returns with the union of the two set
func (i IntSet) Union(i2 IntSet) IntSet {
	i3 := IntSet{make([]int, 0, len(i.data)+len(i2.data))}
	var n1, n2 int
	for n1 < len(i.data) || n2 < len(i2.data) {
		if n2 >= len(i2.data) || n1 < len(i.data) && i.data[n1] < i2.data[n2] {
			i3.data = append(i3.data, i.data[n1])
			n1++
		} else if n1 >= len(i.data) || n2 < len(i2.data) && i2.data[n2] < i.data[n1] {
			i3.data = append(i3.data, i2.data[n2])
			n2++
		} else {
			i3.data = append(i3.data, i.data[n1])
			n1++
			n2++
		}
	}
	return i3
}

// Each runs the given function on all elements of the set
func (i IntSet) Each(f func(val int)) {
	for _, v := range i.data {
		f(v)
	}
}
