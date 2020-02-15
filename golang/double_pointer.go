package main

import (
	"fmt"
	"sort"
)

func maxArea(height []int) int {
	n := len(height)
	maxNum := 0
	left, right := 0, n-1
	for left < right {
		maxNum = max(maxNum, (right-left)*max(height[left], height[right])-min(height[left], height[right]))
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return maxNum
}

func trap(height []int) int {
	n := len(height)
	leftHeight, rightHeight := 0, 0
	i, j := 0, n-1
	total := 0
	for i < j {
		if leftHeight <= rightHeight {
			if height[i] <= leftHeight {
				total += leftHeight - height[i]
			} else {
				leftHeight = height[i]
			}
			i += 1
		} else {
			if height[j] <= rightHeight {
				total += rightHeight - height[j]
			} else {
				rightHeight = height[j]
			}
			j -= 1
		}
	}
	return total
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	sort.Ints(nums)
	var left, right, sum int
	result := make([][]int, 0)
	for i := 0; i < n-2; i++ {
		if nums[i] >= 0 {
			return result
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left = i + 1
		right = n - 1
		for left < right {
			sum = nums[i] + nums[left] + nums[right]
			if sum == 0 {
				fmt.Println(left, right, i)
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[right-1] == nums[right] {
					right -= 1
				}
				for left < right && nums[left+1] == nums[left] {
					left += 1
				}
				right -= 1
				left += 1
			} else if sum < 0 {
				left += 1
			} else {
				right -= 1
			}
		}
	}
	return result
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	cur1, cur2 := headA, headB
	a, b := false, false
	for {
		if cur1 == nil && a {
			break
		} else if cur1 == nil && !a {
			a = true
			cur1 = headB
		}
		if cur2 == nil && b {
			break
		} else if cur2 == nil && !b {
			b = true
			cur2 = headA
		}
		if cur2 == cur1 {
			return cur1
		}
		cur1, cur2 = cur1.Next, cur2.Next
	}
	return nil
}
