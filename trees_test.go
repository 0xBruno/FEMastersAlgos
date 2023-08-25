package main

import (
	"reflect"
	"sort"
	"testing"
)

func newBinaryTrees() (*TreeNode, *TreeNode) {
	tree1 := &TreeNode{
		value: 20,
		left: &TreeNode{
			value: 10,
			right: &TreeNode{
				value: 15,
				right: nil,
				left:  nil,
			},
			left: &TreeNode{
				value: 5,
				right: &TreeNode{
					value: 7,
					right: nil,
					left:  nil,
				},
				left: nil,
			},
		},
		right: &TreeNode{
			value: 50,
			right: &TreeNode{
				value: 100,
				left:  nil,
				right: nil,
			},
			left: &TreeNode{
				value: 30,
				right: &TreeNode{
					value: 45,
					left:  nil,
					right: nil,
				},
				left: &TreeNode{
					value: 29,
					left:  nil,
					right: nil,
				},
			},
		},
	}

	tree2 := &TreeNode{
		value: 20,
		left: &TreeNode{
			value: 10,
			left: &TreeNode{
				value: 29,
				left: &TreeNode{
					value: 21,
					left:  nil,
					right: nil,
				},
				right: nil,
			},
			right: &TreeNode{
				value: 15,
				right: nil,
				left:  nil,
			},
		},
		right: &TreeNode{
			value: 50,
			right: nil,
			left: &TreeNode{
				value: 30,
				right: &TreeNode{
					value: 45,
					right: &TreeNode{
						value: 49,
						left:  nil,
						right: nil,
					},
					left: nil,
				},
				left: &TreeNode{
					value: 29,
					left: &TreeNode{
						value: 21,
						left:  nil,
						right: nil,
					},
					right: nil,
				},
			},
		},
	}
	return tree1, tree2
}
func TestPreOrderSearch(t *testing.T) {
	tree1, tree2 := newBinaryTrees()
	result1 := PreOrderSearch(tree1)
	test1 := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	sort.Ints(test1)
	sort.Ints(result1)

	if !reflect.DeepEqual(result1, test1) {
		t.Fatal(result1, test1)
	}

	result2 := PreOrderSearch(tree2)
	test2 := []int{20, 10, 29, 21, 15, 50, 30, 29, 21, 45, 49}
	sort.Ints(test2)
	sort.Ints(result2)

	if !reflect.DeepEqual(result2, test2) {
		t.Fatal(result2, test2)
	}

}

func TestInOrderSearch(t *testing.T) {
	tree1, _ := newBinaryTrees()
	result1 := InOrderSearch(tree1)
	test1 := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	sort.Ints(test1)
	sort.Ints(result1)

	if !reflect.DeepEqual(result1, test1) {
		t.Fatal(result1, test1)
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	tree1, _ := newBinaryTrees()

	got := BreadthFirstSearch(tree1, 45)
	got2 := BreadthFirstSearch(tree1, 7)
	got3 := BreadthFirstSearch(tree1, 69)

	if got != true {
		t.Fatal(got)
	}

	if got2 != true {
		t.Fatal(got2)
	}

	if got3 != false {
		t.Fatal(got3)
	}

}

func TestCompareBinaryTrees(t *testing.T) {
	tree1, tree2 := newBinaryTrees()

	if !CompareBinaryTrees(tree1, tree1) {
		t.Fatal("something wrong in comparison, comparing same tree")
	}

	if CompareBinaryTrees(tree1, tree2) {
		t.Fatal("something wrong in comparison, comparing different tree but results is equal")
	}
}
