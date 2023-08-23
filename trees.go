package main

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
