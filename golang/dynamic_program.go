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
	for i:=0;i<n;i++ {
		for j:=i;j<n;j++ {
			if a[i][j] &&j-i>max {
				max = j-i
				begin = i
				end = j
			}
		}
	}
	fmt.Println(begin, end)
	return s[begin:end+1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	var a [101][101]int
	can := true
	for i:=0;i<m;i++ {
		if obstacleGrid[i][0]==1 {
			can =false
		}
		if can {
			a[i][0] = 1
		} else {
			a[i][0] = 0
		}
	}
	can = true
	for j:= 0;j<n;j++ {
		if obstacleGrid[0][j] ==1 {
			can = false
		}
		if can {
			a[0][j] = 1
		} else {
			a[0][j] = 0
		}
	}
	for i:=1;i<m;i++ {
		for j:=1;j<n;j++ {
			if obstacleGrid[i][j] ==1 {
				a[i][j] = 0
			}else {
				a[i][j] = a[i-1][j] + a[i][j-1]
			}
		}
	}
	return a[m-1][n-1]
}

func wordBreak(s string, wordDict []string) bool {
	return false
}

func min(first int, second int) int {
	if first<second{
		return first
	} else {
		return second
	}
}

func minimumTotal(triangle [][]int) int {
	for i:=1;i<len(triangle);i++ {
		triangle[i][0] += triangle[i][0]
		for j:=1;j<len(triangle[i]);j++ {
			triangle[i][j] += min(triangle[i-1][j-1], triangle[i-1][j])
		}
	}
	min := 0
	for j:=0;j<len(triangle[len(triangle)-1]);j++ {
		if min> triangle[len(triangle)-1][j] {
			min = triangle[len(triangle)-1][j]
		}
	}
	return min
}