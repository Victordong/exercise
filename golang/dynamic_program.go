package main

import "fmt"

func longestPalindrome(s string) string {
	var a [1001][1001]bool
	n := len(s)
	if n == 0 {
		return s
	}
	for i := 0; i < n; i++ {
		a[i][i] = true
	}
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			a[i-1][i] = true
		}
	}
	for m := 2; m < n; m++ {
		for i := 0; i < n; i++ {
			j := m + i
			if j >= n {
				break
			}
			if s[i] == s[j] && a[i+1][j-1] {
				a[i][j] = true
			} else {
				a[i][j] = false
			}
		}
	}
	max := -1
	begin := 0
	end := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if a[i][j] && j-i > max {
				max = j - i
				begin = i
				end = j
			}
		}
	}
	fmt.Println(begin, end)
	return s[begin : end+1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	var a [101][101]int
	can := true
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			can = false
		}
		if can {
			a[i][0] = 1
		} else {
			a[i][0] = 0
		}
	}
	can = true
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			can = false
		}
		if can {
			a[0][j] = 1
		} else {
			a[0][j] = 0
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				a[i][j] = 0
			} else {
				a[i][j] = a[i-1][j] + a[i][j-1]
			}
		}
	}
	return a[m-1][n-1]
}

func wordBreak(s string, wordDict []string) bool {
	return false
}

func minimumTotal(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		triangle[i][0] += triangle[i-1][0]
		triangle[i][len(triangle[i])-1] += triangle[i-1][len(triangle[i])-2]
		for j := 1; j < len(triangle[i])-1; j++ {
			triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
		}
	}
	min := int(^uint(0) >> 1)
	for j := 0; j < len(triangle[len(triangle)-1]); j++ {
		if min > triangle[len(triangle)-1][j] {
			min = triangle[len(triangle)-1][j]
		}
	}
	return min
}

func min(first int, second int) int {
	if first < second {
		return first
	} else {
		return second
	}
}

func max(first int, second int) int {
	if first < second {
		return second
	} else {
		return first
	}
}

func maxProduct(nums []int) int {
	var imax = 1
	var imin = 1
	var max_num = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			temp := imax
			imax = imin
			imin = temp
		}
		imax = max(imax*nums[i], nums[i])
		imin = min(imin*nums[i], nums[i])
		max_num = max(max_num, imax)
	}
	return imax
}

func maxProfit1(prices []int) int {
	min := prices[0]
	max := ^min
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > max {
			max = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	if max < 0 {
		return 0
	} else {
		return max
	}
}

func maxProfit2(prices []int) int {
	n := len(prices)
	var b [100000]int
	b[0] = 0
	for i := 1; i < n; i++ {
		max := 0
		for j := 0; j < i; j++ {
			var val int
			if prices[i] > prices[j] {
				val = 0
			} else {
				val = prices[j] - prices[i]
			}
			if max < b[j]+val {
				max = b[j] + val
			}
		}
		b[i] = max
	}
	return b[n-1]
}

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	max := 0
	for i := 0; i < n-1; i++ {
		if prices[i] < prices[i+1] {
			max += prices[i+1] - prices[i]
		}
	}
	return max
}
