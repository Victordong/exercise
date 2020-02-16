package main

import "fmt"

func main() {
	dp := []int{1, 2, 3}
	left, right := 0, len(dp)
	for left < right {
		mid := left + (right-left)/2
		if dp[mid] < 3 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	fmt.Println(left)
}
