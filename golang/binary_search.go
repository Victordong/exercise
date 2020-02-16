package main

import "fmt"

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

func search1(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	var mid int = (left + right) / 2
	if left == mid {
		for i, value := range nums {
			if value == target {
				return i
			}
		}
		return -1
	}
	for left < mid {
		if nums[mid] == target {
			return mid
		}
		if nums[left] < nums[right] {
			if nums[mid] > target {
				right = mid
			} else if nums[mid] < target {
				left = mid
			}
		} else if nums[mid] < nums[left] {
			if nums[mid] > target {
				right = mid
			} else if nums[mid] < target {
				if nums[right] == target {
					return right
				} else if nums[right] < target {
					right = mid
				} else if nums[right] > target {
					left = mid
				}
			}
		} else if nums[mid] > nums[left] {
			if nums[mid] < target {
				left = mid
			} else if nums[mid] > target {
				if nums[left] == target {
					return left
				} else if nums[left] < target {
					right = mid
				} else if nums[left] > target {
					left = mid
				}
			}
		}
		mid = (left + right) / 2
	}
	return -1
}

func findPeakElement1(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n - 1
	}
	if n == 2 {
		if nums[0] < nums[1] {
			return 1
		} else {
			return 0
		}
	}
	left, right := 0, n-1
	if nums[left] > nums[left+1] {
		return left
	}
	if nums[right] > nums[right-1] {
		return right
	}
	var mid int = (left + right) / 2
	for left < mid && mid <= n-2 && mid >= 1 {
		if nums[mid] > nums[mid+1] && nums[mid] > nums[mid-1] {
			return mid
		}
		if nums[mid] < nums[right] {
			left = mid
		} else if nums[mid] < nums[left] {
			right = mid
		} else {
			if nums[mid] < nums[mid+1] {
				left = mid + 1
			} else if nums[mid] < nums[mid-1] {
				right = mid - 1
			}
		}
		mid = (left + right) / 2
	}
	return mid
}

func searchRange(nums []int, target int) []int {
	n := len(nums)
	left, right := 0, n
	valueLeft, valueRight := -1, -1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		} else {
			valueLeft, valueRight = mid, mid
			break
		}
	}
	if valueLeft != -1 {
		left, right = 0, valueLeft
		for left < right {
			mid := (left + right) / 2
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid
				valueLeft = mid
			}
		}
		left, right = valueRight, n-1
		for left < right {
			mid := right - (right-left)/2
			if nums[mid] > target {
				right = mid - 1
			} else {
				left = mid
				valueRight = mid
			}
		}
	}
	return []int{valueLeft, valueRight}
}

func searchMatrix1(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	a, b := 0, m-1
	for {
		if matrix[a][b] == target {
			return true
		}
		if a == n-1 || b == 0 {
			break
		}
		if matrix[a][b-1] >= target {
			b = b - 1
		} else if matrix[a+1][b] <= target {
			a = a + 1
		} else if matrix[a][b-1] < target {
			a = a + 1
			b = b - 1
		}
	}
	if a == n-1 {
		for j := b; j >= 0; j-- {
			if matrix[a][j] == target {
				return true
			}
		}
	}
	if b == 0 {
		for i := a; i < n; i++ {
			if matrix[i][b] == target {
				return true
			}
		}
	}
	return false
}

func search(nums []int, target int) bool {
	n := len(nums)
	left, right := 0, n
	equal := false
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > nums[left] {
			if nums[mid] > target {
				if nums[left] > target {
					left = mid + 1
				} else if nums[left] < target {
					right = mid
				} else {
					return true
				}
			} else if nums[mid] < target {
				left = mid + 1
			}
		} else if nums[mid] < nums[left] {
			if nums[mid] > target {
				right = mid
			} else if nums[mid] < target {
				if nums[left] > target {
					left = mid + 1
				} else if nums[left] < target {
					right = mid
				} else {
					return true
				}
			}
		} else {
			equal = true
			break
		}
	}
	if equal {
		for _, value := range nums {
			if value == target {
				return true
			}
		}
	}
	return false
}

func count(matrix [][]int, num int, length int) int {
	result := 0
	for i := 0; i < length; i++ {
		left, right := 0, length-1
		for left < right {
			mid := right - (right-left)/2
			if matrix[i][mid] > num {
				right = mid - 1
			} else {
				left = mid
			}
		}
		if left == 0 {
			if matrix[i][left] <= num {
				result += 1
			}
		} else {
			result = result + (left + 1)
		}
	}
	return result
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	min, max := matrix[0][0], matrix[n-1][n-1]+1
	for min < max {
		mid := min + (max-min)/2
		num := count(matrix, mid, n)
		fmt.Println(mid, num)
		if num < k {
			min = mid + 1
		} else {
			max = mid
		}
	}
	return min
}

//
//func partFindMedia(nums1 []int, nums2 []int, k int) int {
//	m, n := len(nums1), len(nums2)
//	i, j := 0, 0
//	mid := k / 2
//	for mid > 0 {
//		if mid+i >= m {
//			if nums1[m-1] < nums2[mid+j] {
//				return nums2[k-m-1]
//			} else {
//				j = mid + j + 1
//			}
//		} else if mid+j >= n {
//			if nums2[n-1] < nums1[mid+i] {
//				return nums1[k-n-1]
//			} else {
//				i = mid + i + 1
//			}
//		} else {
//			if nums1[mid+i] < nums2[mid+j] {
//				i = mid + i + 1
//			} else {
//				j = mid + j + 1
//			}
//		}
//		mid = (k - k/2) / 2
//	}
//	return
//}
//
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//
//}

func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n - 1
	}
	left, right := 0, n-1
	if nums[left] > nums[left+1] {
		return left
	}
	if nums[right] > nums[right-1] {
		return right
	}
	var mid int = (left + right) / 2
	for left < mid && mid <= n-2 && mid >= 1 {
		if nums[mid] > nums[mid+1] && nums[mid] > nums[mid-1] {
			return mid
		}
		if nums[mid] < nums[mid+1] {
			left = mid
		} else if nums[mid] < nums[mid-1] {
			right = mid
		}
		mid = (left + right) / 2
	}
	return mid
}
