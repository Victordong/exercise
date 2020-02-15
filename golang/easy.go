package main

import (
	"sort"
	"strconv"
)

func kthSmallest1(matrix [][]int, k int) int {
	n := len(matrix)
	var i int
	var j int
	if k%n == 0 {
		i = k/n - 1
		j = n
	} else {
		i = k/n - 1
		j = k - i*n - 1
	}
	return matrix[i][j]
}

func twoSum(nums []int, target int) []int {
	n := len(nums)
	hashMap := make(map[int]int)
	for i := 0; i < n; i++ {
		hashMap[nums[i]] = i
	}
	for i := 0; i < n; i++ {
		if _, ok := hashMap[target-nums[i]]; ok {
			if i != hashMap[target-nums[i]] {
				return []int{i, hashMap[target-nums[i]]}
			}
		}
	}
	return nil
}

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	temp := x
	total := 0
	for {
		total = 10*total + temp%10
		temp = temp / 10
		if temp == 0 {
			break
		}
	}
	if total == x {
		return true
	}
	return false
}

func reverse1(x int) int {
	total := 0
	re := false
	if x < 0 {
		re = true
		x = x * -1
	}
	for {
		total = total*10 + x%10
		x = x / 10
		if x == 0 {
			break
		}
	}
	if re {
		total = total * -1
	}
	return total
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	i := 0
	j := n - 1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			i += 1
		}
		if matrix[i][j] > target {
			j -= 1
		}
	}
	return false
}

type largestNumberList []int

func (l largestNumberList) Len() int {
	return len(l)
}

func (l largestNumberList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l largestNumberList) Less(i, j int) bool {
	return strconv.Itoa(l[i])+strconv.Itoa(l[j]) > strconv.Itoa(l[j])+strconv.Itoa(l[i])
}

func largestNumber(nums []int) string {
	l := make(largestNumberList, len(nums))
	copy(l, nums)
	sort.Sort(l)
	result := ""
	for _, value := range l {
		result += strconv.Itoa(value)
	}
	part, err := strconv.Atoi(result)
	if err != nil {
		return result
	} else {
		return strconv.Itoa(part)
	}
}
