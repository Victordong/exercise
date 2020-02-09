package main

import "sort"

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	partPermute(nums, []int{}, &result)

	return result
}

func partPermute(nums []int, values []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, values)
		return
	}
	for i := 0; i < len(nums); i++ {
		newNums := make([]int, len(nums))
		copy(newNums, nums)
		newValues := make([]int, len(values))
		copy(newValues, values)
		partPermute(append(newNums[:i], newNums[i+1:]...), append(newValues, nums[i]), result)
	}
}

func partExist(board [][]byte, used [][]bool, X int, Y int, word string, index int, maxX int, maxY int) bool {
	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}
	if board[X][Y] == word[index] {
		if index == len(word)-1 {
			return true
		}
		for i := 0; i < len(dx); i++ {
			newX := X + dx[i]
			newY := X + dy[i]
			if newX <= maxX && newY <= maxY && newX >= 0 && newY >= 0 && !used[newX][newY] {
				used[newX][newY] = true
				if partExist(board, used, newX, newY, word, index+1, maxX, maxY) {
					return true
				}
				used[newX][newY] = false
			}
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])
	used := make([][]bool, 0)
	for i := 0; i < n; i++ {
		used = append(used, make([]bool, m))
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == word[0] {
				if partExist(board, used, i, j, word, 0, n-1, m-1) {
					return true
				}
			}
		}
	}
	return false
}

func partIsland(grid [][]byte, used [][]bool, X int, Y int, maxX int, maxY int) {
	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}
	if grid[X][Y] == '1' {
		for i := 0; i < 4; i++ {
			newX, newY := X+dx[i], Y+dy[i]
			if newY >= 0 && newY <= maxY && newX >= 0 && newX <= maxX && !used[newX][newY] {
				used[newX][newY] = true
				partIsland(grid, used, newX, newY, maxX, maxY)
			}
		}
	} else {
		used[X][Y] = false
	}
}

func numIslands(grid [][]byte) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	used := make([][]bool, 0)
	for i := 0; i < n; i++ {
		used = append(used, make([]bool, m))
	}
	total := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if !used[i][j] && grid[i][j] == '1' {
				total += 1
				used[i][j] = true
				partIsland(grid, used, i, j, n-1, m-1)
			}
		}
	}
	return total
}

func partGenerateParenthesis(left int, right int, number int, cur string, result *[]string) {
	if left == 0 {
		for i := 0; i < right; i++ {
			cur = cur + ")"
		}
		*result = append(*result, cur)
		return
	}
	if number > 0 {
		partGenerateParenthesis(left, right-1, number-1, cur+")", result)
	}
	partGenerateParenthesis(left-1, right, number+1, cur+"(", result)
}

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	partGenerateParenthesis(n-1, n, 1, "(", &result)
	return result
}

func partSub(nums []int, begin int, now []int, result *[][]int) {
	*result = append(*result, now)
	for i := begin; i < len(nums); i++ {
		curNow := make([]int, len(now))
		copy(curNow, now)
		curNow = append(curNow, nums[i])
		partSub(nums, i+1, curNow, result)
	}
}

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	partSub(nums, 0, []int{}, &result)
	return result
}

func partSubDup(nums []int, begin int, now []int, result *[][]int) {
	*result = append(*result, now)
	for i := begin; i < len(nums); i++ {
		if i == begin || (i > begin && nums[i] != nums[i-1]) {
			curNow := make([]int, len(now))
			copy(curNow, now)
			curNow = append(curNow, nums[i])
			partSub(nums, i+1, curNow, result)
		}
	}
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	partSub(nums, 0, []int{}, &result)
	return result
}

func partPermuteUnique(nums []int, now []int, result *[][]int) {
	if len(nums) == 0 {
		*result = append(*result, now)
	}
	for i, value := range nums {
		if i == 0 || (i > 0 && nums[i-1] != nums[i]) {
			newNums := make([]int, len(nums))
			copy(newNums, nums)
			newNow := make([]int, len(now))
			copy(newNow, now)
			newNow = append(newNow, value)
			partPermuteUnique(append(newNums[:i], newNums[i+1:]...), newNow, result)
		}
	}
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	partPermuteUnique(nums, []int{}, &result)
	return result
}