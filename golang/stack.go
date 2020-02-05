package main

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
