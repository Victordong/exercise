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
