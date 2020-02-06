package main

import "fmt"

func position(nums []int, left int, right int) int {
	temp := nums[left]
	for left < right {
		for left < right && nums[right] <= temp {
			right -= 1
		}
		nums[left] = nums[right]
		for left < right && nums[left] >= temp {
			left += 1
		}
		nums[right] = nums[left]
	}
	nums[left] = temp
	return left
}

func quickSort(nums []int, left int, right int) {
	if left >= right {
		return
	}
	pos := position(nums, left, right)
	fmt.Println(left, right, pos)
	quickSort(nums, left, pos-1)
	quickSort(nums, pos+1, right)
}

func findKthLargest(nums []int, k int) int {
	left, right := 0, len(nums)-1
	for {
		if left >= right {
			break
		}
		pos := position(nums, left, right)
		if pos < k-1 {
			left = pos + 1
		} else if pos > k-1 {
			right = pos - 1
		} else {
			return nums[pos]
		}
	}
	return nums[k-1]
}
