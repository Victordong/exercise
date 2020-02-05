package main

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
