package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	cur := root
	trees := make([]*TreeNode, 0)
	trees = append(trees, cur)
	for len(trees) != 0 {
		for cur.Left != nil {
			trees = append(trees, cur.Left)
			cur = cur.Left
		}
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		result = append(result, cur.Val)
		for cur.Right == nil && len(trees) != 0 {
			cur = trees[len(trees)-1]
			trees = trees[:len(trees)-1]
			result = append(result, cur.Val)
		}
		if cur.Right != nil {
			cur = cur.Right
			trees = append(trees, cur)
		}

	}
	return result
}

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	cur := root
	trees := make([]*TreeNode, 0)
	trees = append(trees, cur)
	for len(trees) != 0 {
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		if cur.Right != nil {
			trees = append(trees, cur.Right)
		}
		if cur.Left != nil {
			trees = append(trees, cur.Left)
		}
		result = append(result, cur.Val)
	}
	return result
}

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	cur := root
	treeFirst := make([]*TreeNode, 0)
	treeSecond := make([]*TreeNode, 0)
	treeFirst = append(treeFirst, cur)
	for len(treeFirst) != 0 {
		cur = treeFirst[len(treeFirst)-1]
		treeFirst = treeFirst[:len(treeFirst)-1]
		treeSecond = append(treeSecond, cur)
		if cur.Left != nil {
			treeFirst = append(treeFirst, cur.Left)
		}
		if cur.Right != nil {
			treeFirst = append(treeFirst, cur.Right)
		}
	}
	for len(treeSecond) != 0 {
		cur = treeSecond[len(treeSecond)-1]
		treeSecond = treeSecond[:len(treeSecond)-1]
		result = append(result, cur.Val)
	}
	return result
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	cur := root
	end := root
	nextEnd := root
	trees := make([]*TreeNode, 0)
	partResult := make([]int, 0)
	trees = append(trees, cur)
	for len(trees) != 0 {
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		partResult = append(partResult, cur.Val)
		if cur.Left != nil {
			nextEnd = cur.Left
			trees = append([]*TreeNode{cur.Left}, trees...)
		}
		if cur.Right != nil {
			nextEnd = cur.Right
			trees = append([]*TreeNode{cur.Right}, trees...)
		}
		if cur == end {
			end = nextEnd
			result = append(result, partResult)
			partResult = partResult[0:0]
		}
	}
	return result
}

func partBalance(root *TreeNode, height int) (int, bool) {
	if root == nil {
		return height, true
	}
	leftHeight, leftBalance := partBalance(root.Left, height+1)
	rightHeight, rightBalance := partBalance(root.Right, height+1)
	ifBalance := leftBalance && rightBalance
	if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
		ifBalance = false
	}
	return max(rightHeight, leftHeight), ifBalance
}

func isBalanced(root *TreeNode) bool {
	_, ifBalance := partBalance(root, 0)
	return ifBalance
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	cur := root
	end := root
	nextEnd := root
	trees := make([]*TreeNode, 0)
	partResult := make([]int, 0)
	trees = append(trees, cur)
	for len(trees) != 0 {
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		partResult = append(partResult, cur.Val)
		if cur.Left != nil {
			nextEnd = cur.Left
			trees = append([]*TreeNode{cur.Left}, trees...)
		}
		if cur.Right != nil {
			nextEnd = cur.Right
			trees = append([]*TreeNode{cur.Right}, trees...)
		}
		if cur == end {
			end = nextEnd
			result = append([][]int{partResult}, result...)
			partResult = make([]int, 0)
		}
	}
	return result
}

func isValidBST(root *TreeNode) bool {
	result := make([]int, 0)
	if root == nil {
		return true
	}
	trees := make([]*TreeNode, 0)
	cur := root
	trees = append(trees, cur)
	for len(trees) != 0 {
		for cur.Left != nil {
			trees = append(trees, cur.Left)
			cur = cur.Left
		}
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		result = append(result, cur.Val)
		for cur.Right != nil && len(trees) != 0 {
			cur = trees[len(trees)-1]
			trees = trees[:len(trees)-1]
			result = append(result, cur.Val)
		}
		if cur.Right != nil {
			trees = append(trees, cur.Right)
			cur = cur.Right
		}
	}
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}
