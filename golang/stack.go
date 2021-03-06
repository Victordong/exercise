package golang

import "fmt"

func dailyTemperatures(T []int) []int {
	n := len(T)
	if n == 0 {
		return nil
	}
	result := make([]int, n)
	stack := make([]int, 0)
	stack = append(stack, 0)
	var cur int
	for i := 1; i < n; i++ {
		for len(stack) > 0 {
			cur = stack[len(stack)-1]
			if T[i] <= T[cur] {
				break
			}
			stack = stack[:len(stack)-1]
			result[cur] = i - cur
		}
		stack = append(stack, i)
	}
	for _, value := range stack {
		result[value] = 0
	}
	return result
}

func isValid(s string) bool {
	wordStack := make([]byte, 0)
	var temp byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			wordStack = append(wordStack, s[i])
		} else {
			if len(wordStack) == 0 {
				return false
			}
			temp = wordStack[len(wordStack)-1]
			wordStack = wordStack[:len(wordStack)-1]
			if (s[i] == ')' && temp != '(') || (s[i] == '}' && temp != '{') || (s[i] == ']' && temp != '[') {
				return false
			}
		}
	}
	if len(wordStack) == 0 {
		return true
	} else {
		return false
	}
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	heightStack := make([]int, 0)
	heightStack = append(heightStack, -1)
	maxNumber := 0
	var curI int
	var currentArea int
	for i := 0; i < len(heights); i++ {
		for len(heightStack) != 1 && heights[heightStack[len(heightStack)-1]] >= heights[i] {
			curI = heightStack[len(heightStack)-1]
			heightStack = heightStack[:len(heightStack)-1]
			currentArea = (i - heightStack[len(heightStack)-1] - 1) * heights[curI]
			maxNumber = max(maxNumber, currentArea)
		}
		heightStack = append(heightStack, i)
	}
	for len(heightStack) != 1 {
		curI = heightStack[len(heightStack)-1]
		heightStack = heightStack[:len(heightStack)-1]
		currentArea = (n - heightStack[len(heightStack)-1] - 1) * heights[curI]
		maxNumber = max(maxNumber, currentArea)
	}
	return maxNumber
}
func partMaximalRectangle(heights [][]int, x int, length int) int {
	heightStack := make([]int, 0)
	heightStack = append(heightStack, -1)
	maxNumber := 0
	var curI int
	var currentArea int
	for i := 0; i < length; i++ {
		for len(heightStack) != 1 && heights[x][heightStack[len(heightStack)-1]] >= heights[x][i] {
			curI = heightStack[len(heightStack)-1]
			heightStack = heightStack[:len(heightStack)-1]
			currentArea = (i - heightStack[len(heightStack)-1] - 1) * heights[x][curI]
			maxNumber = max(maxNumber, currentArea)
		}
		heightStack = append(heightStack, i)
	}
	for len(heightStack) != 1 {
		curI = heightStack[len(heightStack)-1]
		heightStack = heightStack[:len(heightStack)-1]
		currentArea = (length - heightStack[len(heightStack)-1] - 1) * heights[x][curI]
		maxNumber = max(maxNumber, currentArea)
	}
	return maxNumber
}

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, 0)
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}
	for i := 0; i < n; i++ {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
		} else {
			dp[0][i] = 0
		}
	}
	maxNumber := partMaximalRectangle(dp, 0, n)
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = dp[i-1][j] + 1
			} else {
				dp[i][j] = 0
			}
		}
		for j := 0; j < n; j++ {
			fmt.Print(dp[i][j], " ")
		}
		partResult := partMaximalRectangle(dp, i, n)
		fmt.Println("     ", i, "    ", partResult)
		maxNumber = max(maxNumber, partResult)
	}
	return maxNumber
}

func scoreOfParentheses(S string) int {
	scoreStack := make([]int, 0)
	scoreStack = append(scoreStack, 0)
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			scoreStack = append(scoreStack, 0)
		} else {
			currentScore := scoreStack[len(scoreStack)-1]
			lastScore := scoreStack[len(scoreStack)-2]
			lastScore += max(currentScore*2, currentScore+1)
			scoreStack[len(scoreStack)-2] = lastScore
			scoreStack = scoreStack[:len(scoreStack)-1]
		}
	}
	return scoreStack[len(scoreStack)-1]
}

type MyQueue struct {
	main []int
	help []int
}

/** Initialize your data structure here. */
func Constructor3() MyQueue {
	return MyQueue{
		main: make([]int, 0),
		help: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.main = append(this.main, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	var cur int
	for len(this.main) != 0 {
		cur = this.main[len(this.main)-1]
		this.main = this.main[:len(this.main)-1]
		this.help = append(this.help, cur)
	}
	result := this.help[len(this.help)-1]
	this.help = this.help[:len(this.help)-1]
	for len(this.help) != 0 {
		cur = this.help[len(this.help)-1]
		this.help = this.help[:len(this.help)-1]
		this.main = append(this.main, cur)
	}
	return result
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	var cur int
	for len(this.main) != 0 {
		cur = this.main[len(this.main)-1]
		this.main = this.main[:len(this.main)-1]
		this.help = append(this.help, cur)
	}
	result := this.help[len(this.help)-1]
	for len(this.help) != 0 {
		cur = this.help[len(this.help)-1]
		this.help = this.help[:len(this.help)-1]
		this.main = append(this.main, cur)
	}
	return result
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.main) == 0
}
