package main

import (
	"log"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func preOrderWalk(curr *TreeNode, path *[]int) *[]int {
	// base case
	if curr == nil {
		return path
	}

	// recurse
	// pre
	*path = append(*path, curr.value)

	// recurse
	preOrderWalk(curr.left, path)
	preOrderWalk(curr.right, path)

	// post
	return path
}

func PreOrderSearch(root *TreeNode) []int {
	return *preOrderWalk(root, &[]int{})
}

func inOrderWalk(curr *TreeNode, path *[]int) *[]int {
	// base case
	if curr == nil {
		return path
	}

	// recurse
	// pre

	// recurse
	inOrderWalk(curr.left, path)
	*path = append(*path, curr.value)
	inOrderWalk(curr.right, path)

	// post
	return path
}

func InOrderSearch(root *TreeNode) []int {
	return *inOrderWalk(root, &[]int{})
}

// not doing post order, you just move the pushing of the path stack after the recursion

func BreadthFirstSearch(root *TreeNode, needle int) bool {
	q := Queue{}
	q.Enqueue(root)

	for {
		if q.Length == 0 {
			break
		}

		curr, err := q.Dequeue()
		if err != nil {
			log.Fatalln(err)
		}

		if curr.(*TreeNode).value == needle {
			return true
		}

		if curr.(*TreeNode).left != nil {
			q.Enqueue(curr.(*TreeNode).left)
		}

		if curr.(*TreeNode).right != nil {
			q.Enqueue(curr.(*TreeNode).right)
		}

	}

	return false
}

func CompareBinaryTrees(a *TreeNode, b *TreeNode) bool {
	// this is actually so simple it makes me want to cry happy tears
	// base case
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.value != b.value {
		return false
	}

	return CompareBinaryTrees(a.left, b.left) && CompareBinaryTrees(a.right, b.right)
}

func DepthFirstSearchBST(current *TreeNode, value int) bool {
	// base case
	if current == nil {
		return false
	}

	if current.value == value {
		return true
	}

	if current.value < value {
		return DepthFirstSearchBST(current.right, value)
	}

	return DepthFirstSearchBST(current.left, value)

}
