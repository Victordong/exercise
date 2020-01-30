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
