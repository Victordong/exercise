package main

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

func isPalindrome(x int) bool {
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
