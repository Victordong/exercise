package main

func findMin(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	left := 0
	right := n - 1
	mid := (left + right) / 2
	all := false
	for left < mid {
		if nums[left] <= nums[right] {
			return nums[left]
		} else {
			if nums[left] > nums[mid] {
				right = mid
				mid = (left + right) / 2
			} else if nums[mid] > nums[right] {
				left = mid
				mid = (left + right) / 2
			} else {
				all = true
				break
			}
		}
	}
	maxNum := ^int(^uint(0) >> 1)
	if all {
		for i := 0; i < n; i++ {
			if maxNum < nums[i] {
				maxNum = nums[i]
			}
		}
	} else {
		maxNum = min(nums[left], nums[right])
	}
	return maxNum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pow(base int, times int) int {
	sum := 1
	for times > 0 {
		sum *= base
		times -= 1
	}
	return sum
}

func partCount(root *TreeNode, height int) int {
	if root.Left == nil {
		return 1
	}
	rightHeight := 0
	var leftNum int
	var rightNum int
	if root.Right != nil {
		rightHeight = 1
		cur := root.Right
		for cur.Left != nil {
			rightHeight += 1
			cur = cur.Left
		}
	}
	if rightHeight == height-1 {
		leftNum = pow(2, height-1) - 1
		rightNum = partCount(root.Right, rightHeight)
	} else {
		leftNum = partCount(root.Left, height-1)
		rightNum = pow(2, height-2) - 1
	}
	return leftNum + rightNum + 1
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	height := 0
	cur := root
	for cur != nil {
		cur = cur.Left
		height += 1
	}
	return partCount(root, height)
}
